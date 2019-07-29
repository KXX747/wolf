package dao

import (
	"context"
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/model"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/net/trace"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// Dao dao interface
type Dao interface {
   Close()
   Ping(ctx context.Context) (err error)

   //登录
   	Login (ctx context.Context,params *model.LoginInSystem)(err interface{},resp []byte)
	LoginOut (ctx context.Context,params *model.LoginOutSystem)(err interface{},resp []byte)

   ////用户信息
	//AddUserDao(ctx context.Context,name string ,mobile string)(reply *account_service.UserReply,err error)
	//UpdateUserDao(ctx context.Context,id_no  ,mobile ,address string) (reply *account_service.UserReply,err error)
	//DeleteUserDao(ctx context.Context,id_no string,content string) (reply *account_service.UserReply,err error)
	//FindUserDao(ctx context.Context,id_no string)(reply *account_service.UserReply,err error)
	//FindUserListDao(ctx context.Context,id_no []string)(reply *account_service.UserListReply,err error)
	//FindUserIsExistDao(ctx context.Context,name string ,mobile string )(reply *account_service.UserReply,err error)
   //
	////用户实名信息
	//FindUserCommonDao(ctx context.Context, in *account_service.UserCommonReq) (reply *account_service.UserCommon,err error)
	//VerifiedIdNoUser(ctx context.Context, in *account_service.UserCommon) (reply *account_service.UserCommon,err error)
	//InsertUserCommonDao(ctx context.Context, in *account_service.UserCommon) (reply *account_service.UserCommon,err error)
   //
   //
   //
	////视频和视频评价
	//File(ctx context.Context, req *pb.UploadFileReq) (mUploadFileResp *pb.UploadFileResp, err error)
	//New(ctx context.Context, in *pb.NewTokenReq) (mNewTokenResp *pb.NewTokenResp, err error)
	//Addevaluation(ctx context.Context, req *pb.EvaluationVodieReq)(mVodieResp *pb.EvaluationVodieResp, err error)
	//Fileallevalby(ctx context.Context, req *pb.EvaluationGetReq) (mVodieResp *pb.EvaluationListByVodieResp, err error)
	//Listfile(ctx context.Context, req *pb.FileListReq) (mFileListResp *pb.FileListResp, err error)
}

// dao dao.
type dao struct {
	db          *sql.DB
	redis       *redis.Pool
	redisExpire int32
	mc          *memcache.Memcache
	mcExpire    int32
}

//基础信息
type BaseInfo struct {
	Version 	string
	Env			string
	Token 		string
	KeyId		string
	Screet		string
}

type RPC struct {
	User *warden.ClientConfig
}


var Conf =&Config{}

type ExpireConfig struct {
	RedisExpire 	xtime.Duration
}

//
type Config struct {
	Grpc		    *warden.ServerConfig
	Http 			*blademaster.ServerConfig
	Mysql 			*sql.Config
	Expire			*ExpireConfig
	Redis   		*redis.Config
	Base 			*BaseInfo
	Log 			*log.Config
	Tracer 			*trace.Config
	UserHttpClient  *blademaster.ClientConfig
	RPCClient2	    *RPC
}


// New new a dao and return.
func New() (Dao) {
	return &dao{
		// mysql
		db: sql.NewMySQL(Conf.Mysql),
		// redis
		redis:       redis.NewPool(Conf.Redis),
		redisExpire: int32(time.Duration(Conf.Expire.RedisExpire) / time.Second),
		// memcache
		//mc:       memcache.New(mc.Demo),
		//mcExpire: int32(time.Duration(mc.DemoExpire) / time.Second),
	}
}

func BuildConfig()*Config{

	return Conf
}

// Close close the resource.
func (d *dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	if err = d.pingMC(ctx); err != nil {
		return
	}
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *dao) pingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func (d *dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}


/**
实现paladin的Set接口
 */
func (c *Config) Set(text string) error {
	var conf Config
	if err := toml.Unmarshal([]byte(text), &conf); err != nil {
		return err
	}
	*c = conf
	return nil
}