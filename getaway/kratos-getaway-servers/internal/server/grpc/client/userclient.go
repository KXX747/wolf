package client

import(
	"context"
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/api/account_service"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/log"
	"google.golang.org/grpc"
)

/**
连接grpc的服务
 */
type UserServer struct {
	cfg *warden.ClientConfig
	userRPCClient account_service.UsersClient
	userDetailCommonClient account_service.UserDetailCommonClient

}

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (account_service.UsersClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "127.0.0.1:38889")
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return account_service.NewUsersClient(conn), nil
}

// NewClient new member grpc client
func NewUserCommonClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (account_service.UserDetailCommonClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "127.0.0.1:38889")
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return account_service.NewUserDetailCommonClient(conn), nil
}





//user rpc client
func NewUserServer(cfg *warden.ClientConfig) (mUserServer *UserServer){
	userRPCClient,err := NewClient(cfg)
	if err!=nil {
		log.Error("userRPCClient warden.ClientConfig err=%s",err)
	}

	userCommonRPCClient,err := NewUserCommonClient(cfg)
	if err!=nil {
		log.Error("userCommonRPCClient warden.ClientConfig err=%s",err)
	}


	mUserServer = &UserServer{
		cfg:cfg,
	}
	mUserServer.userRPCClient = userRPCClient
	mUserServer.userDetailCommonClient = userCommonRPCClient
	return

}


//用户注册
func(service *UserServer)AddUserDao(ctx context.Context,mAddUserReq *account_service.AddUserReq)(reply *account_service.UserReply,err error){

	if reply,err =service.userRPCClient.AddUser(ctx,mAddUserReq);err!=nil {
		log.Info("userRPCClient AddUserDao err=%s mAddUserReq=%s",err,mAddUserReq)

		return nil,err
	}
	return
}

//修改用户
func(service *UserServer)UpdateUserDao(ctx context.Context,mUpdateUserReq *account_service.UpdateUserReq) (reply *account_service.UserReply,err error){

	if reply,err =service.userRPCClient.UpdateUser(ctx,mUpdateUserReq);err!=nil  {
		log.Info("userRPCClient UpdateUserDao err=%s mUpdateUserReq=%s",err,mUpdateUserReq)

		return nil,err
	}

	return
}

//删除
func(service *UserServer)DeleteUserDao(ctx context.Context,mDeleteUserReq *account_service.DeleteUserReq) (reply *account_service.UserReply,err error){
	if reply,err =service.userRPCClient.DeleteUser(ctx,mDeleteUserReq);err!=nil  {
		log.Info("userRPCClient DeleteUserDao err=%s mDeleteUserReq=%s",err,mDeleteUserReq)

		return nil,err
	}

	return
}

//查询用户
func(service *UserServer)FindUserDao(ctx context.Context,mFindUserReq *account_service.FindUserReq)(reply *account_service.UserReply,err error){
	if reply,err =service.userRPCClient.FindUser(ctx,mFindUserReq);err!=nil {
		log.Info("userRPCClient FindUserDao err=%s IdNo=%s",err,mFindUserReq.IdNo)
		return nil,err
	}
	return
}

//查询多个用户信息
func(service *UserServer)FindUserListDao(ctx context.Context,mFindUserReq *account_service.FindUserReq)(reply *account_service.UserListReply,err error){

	if reply,err =service.userRPCClient.FindUserList(ctx,mFindUserReq);err!=nil {
		log.Info("userRPCClient FindUserListDao err=%s IdNo=%s",err,mFindUserReq.IdNo)
		return nil,err
	}

	return
}

////查询用户是否存在
//func(service *UserServer)FindUserIsExistDao(ctx context.Context,name string ,mobile string )(reply *account_service.UserReply,err error){
//
//	return
//}


//验证用户登录查询用户
func(service *UserServer)VerifiedIdNoUser(ctx context.Context, in *account_service.UserCommon) (reply *account_service.UserCommon,err error){

	return
}

//更新用户信息
func(service *UserServer) UpdateUserCommon(ctx context.Context, in *account_service.UserCommon) (reply *account_service.UserCommon,err error) {

	if reply,err =service.userDetailCommonClient.UpdateUserCommon(ctx,in);err!=nil {
		log.Info("userDetailCommonClient UpdateUserCommon err=%s idNo=%s",err,in)
		return nil,err
	}
	return
}

//更新用户信息
func(service *UserServer) FindUserCommon(ctx context.Context, in *account_service.UserCommonReq) (reply *account_service.UserCommon,err error) {

	if reply,err =service.userDetailCommonClient.FindUserCommon(ctx,in);err!=nil {
		log.Info("userDetailCommonClient FindUserCommon err=%s in=%s",err,in)
		return nil,err
	}

	return
}









