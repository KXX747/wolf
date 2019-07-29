package http

import (
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/model"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/ecode"
)

const (
	LOGIN_URL="/user-account-server/login"
	LOGIN_OUT_URL="/user-account-server/loginout"
)


/**
用户登录
*/
func LoginSys(c *blademaster.Context)  {

	loginInParams:=new(model.LoginInSystem)
	if err := c.Bind(loginInParams); err != nil {
		c.JSON(nil, ecode.ReqParamErr)
		return
	}

	//返回结果
	if code,resp:=svc.Dao.Login(c,loginInParams);code!=nil{
		c.JSON(nil, code.(*ecode.Code))
	}else {
		c.JSON(resp, nil)
	}

}


/**
退出用户
*/
func LoginOut(c *blademaster.Context)  {

	mLoginOutSystem:=new(model.LoginOutSystem)
	if err := c.Bind(mLoginOutSystem); err != nil {
		c.JSON(nil, ecode.ReqParamErr)
		return
	}

	//返回结果
	if code,resp:=svc.Dao.LoginOut(c,mLoginOutSystem);code!=nil{
		c.JSON(nil, code.(*ecode.Code))
	}else {
		c.JSON(resp, nil)
	}
}
