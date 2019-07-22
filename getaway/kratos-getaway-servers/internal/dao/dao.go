package dao

import (
	"context"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/BurntSushi/toml"
	"github.com/bilibili/kratos/pkg/net/trace"
)

// Dao dao interface
type Dao interface {
   Close()
   Ping(ctx context.Context) (err error)
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