package dao

import (
	"context"
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/api/account_service"
)




//用户注册
func(d *dao)AddUserDao(ctx context.Context,name string ,mobile string)(reply *account_service.UserReply,err error){

	return
}

//修改用户
func(d *dao)UpdateUserDao(ctx context.Context,id_no  ,mobile ,address string) (reply *account_service.UserReply,err error){

	return
}

//删除
func(d *dao)DeleteUserDao(ctx context.Context,id_no string,content string) (reply *account_service.UserReply,err error){

	return
}

//查询用户
func(d *dao)FindUserDao(ctx context.Context,id_no string)(reply *account_service.UserReply,err error){

	return
}

//查询多个用户信息
func(d *dao)FindUserListDao(ctx context.Context,id_no []string)(reply *account_service.UserListReply,err error){

	return
}

//查询用户是否存在
func(d *dao)FindUserIsExistDao(ctx context.Context,name string ,mobile string )(reply *account_service.UserReply,err error){

	return
}



//common--查询用户
func(d *dao)FindUserCommonDao(ctx context.Context, in *account_service.UserCommonReq) (reply *account_service.UserCommon,err error){

	return
}

//common--查询用户
func(d *dao)VerifiedIdNoUser(ctx context.Context, in *account_service.UserCommon) (reply *account_service.UserCommon,err error){

	return
}


