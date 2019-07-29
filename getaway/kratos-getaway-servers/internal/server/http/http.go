package http

import (
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/server/grpc/client"
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/service"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/middleware/auth"
)

var (
	svc *service.Service
	userRPCServer *client.UserServer
	userHttpClient  *bm.Client
	streamRPCClient  *client.StreamServer
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {

	svc = s
	initServer(svc)
	engine = bm.DefaultServer(s.AppConfig.Http)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	////限流中间件
	//limiter := bm.NewRateLimiter(nil)
	//e.Use(limiter.Limit())
	g := e.Group("/servers/user/api/v1")
	{
		g.GET("/start", howToStart)
	}

	//登录退出服务
	loginServer := e.Group("/servers/login/api/v1")
	{

		loginServer.POST("/login",LoginSys)
		loginServer.POST("/loginout",auth.CheckLogin,LoginOut)

	}
	//用户
	userServer := e.Group("/servers/account/api/v1",auth.CheckLogin)
	{
		userServer.GET("/newUser", NewUser)
		userServer.GET("/updateUser", UpdateUser)
		userServer.GET("/deleteUser", DeleteUser)
		userServer.GET("/updatecard", updatecard)
		userServer.GET("/findUserByIdNo", FindUserByIdNo)
		userServer.GET("/findUserListByIdNo", FindUserListByIdNo)

		//common
		userServer.GET("/findCommonUserByIdNo", FindCommonUserByIdNo)
		userServer.GET("/updateCommonUser", UpdateCommonUser)
	}

	//流视频处理
	streamServer := e.Group("/servers/stream/api/v1",auth.CheckLogin)
	{
		streamServer.POST("/upload",UploadFile);
		streamServer.POST("/findidno",FindListFileByIdNo);
		streamServer.POST("/findByid",FindFileAllevalbyVid);
		streamServer.POST("/newalleval",NewFileAllevalbyVid);
	}
}

//grpc client
func initServer(s *service.Service)  {
	userRPCServer=client.NewUserServer(s.AppConfig.RPCClient2.User)
	userHttpClient =bm.NewClient(s.AppConfig.UserHttpClient)
	//
	streamRPCClient=client.NewStreamServer(s.AppConfig.RPCClient2.Stream)


}

