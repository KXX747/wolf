package http

import "github.com/bilibili/kratos/pkg/net/http/blademaster"

const(

	//http://127.0.0.1:38888/account.service.Users/AddUser?name=liuhui122&mobile=15616121488
	//http://127.0.0.1:38888/account.service.Users/UpdateUser?address=USA&mobile=15616121488&id_no=c99e67d2-40e5-42f3-bb5c-c301d98ec37b&name=我是来修改的liuhui122
	//http://127.0.0.1:38888/account.service.Users/DeleteUser?IdNo=abc7c8b2-b33f-4418-871b-5fad3a590602&Content="删除"
	//http://127.0.0.1:38888/account.service.Users/FindUser?id_no=c7e760fb-19a0-43d2-9265-be863704c77a
	//http://127.0.0.1:38888/account.service.Users/FindUserList?id_no=006be871-1621-4d5d-be5c-c52d0844f2d7,006be871-1621-4d5d-be5c-c52d0844f2d7,bb6f406d-217d-4c6c-a08d-6465ad8b98f4,b44612f9-720f-41ef-bb62-0e4be9c893fd,630d8cf7-5c3d-4448-b277-f16489c04f50
	NewUserURL="%s:%s/account.service.Users/AddUser"
	UpdateUserUrl="%s:%s/account.service.Users/UpdateUser"
	DeleteUserUrl="%s:%s/account.service.Users/DeleteUser"
	FindUserUrl="%s:%s/account.service.Users/FindUser"
	FindUserListUrl="%s:%s/account.service.Users/FindUserList"


	//http://127.0.0.1:38888/account.service.UserDetailCommon/UpdateUserCommon?IdNo=1fdc36f8-1551-4d72-9bc0-89bb71f2964b&Mobile=13600000022&Address=CHINA-SHANGHAI&Name=guangzhou1
	//http://127.0.0.1:38888/account.service.UserDetailCommon/FindUserCommon?id_no=c7e760fb-19a0-43d2-9265-be863704c77a&querytype=1
	UpdateUserCommonUrl="%s:%s/account.service.UserDetailCommon/UpdateUserCommon"
	FindUserCommonUrl="%s:%s/account.service.UserDetailCommon/FindUserCommon"
)


/**
添加用户
 */
func NewUser(c *blademaster.Context)  {

}

/**
修改用户
*/
func UpdateUser(c *blademaster.Context)  {

}

/**
删除用户
*/
func DeleteUser(c *blademaster.Context)  {

}

/**
获取用户信息
 */
func FindUserByIdNo(c *blademaster.Context)  {
	
}

/**
获取多个用户信息
*/
func FindUserListByIdNo(c *blademaster.Context)  {


}


/**
修改用户common
*/
func UpdateCommonUser(c *blademaster.Context)  {

}

/**
获取用户信息common
*/
func FindCommonUserByIdNo(c *blademaster.Context)  {

}




/**
一是 Request Body 就是整个文件内容，通过请求头（即 Header ）中的 Content-Type 字段来指定文件类型。
二是用 multipart 表单方式来上传
*/
//上传信息照片
func updatecard(c *blademaster.Context) {
	

}