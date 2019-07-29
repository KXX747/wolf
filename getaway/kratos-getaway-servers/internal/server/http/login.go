package http

import (
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/model"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"

	"net/url"
)

const (
	LOGIN_URL="http://127.0.0.1:38888/app/v1/common/login"
	//LOGIN_URL="http://%s/app/v1/common/login"
	LOGIN_OUT_URL="http://127.0.0.1:38888/app/v1/common/loginout"
	//LOGIN_OUT_URL="http://%s/app/v1/common/loginout"
	USER_UPLOAD_RUL="http://127.0.0.1:38888/user-account-server/upload/image"
	//LOGIN_OUT_URL="http://%s/app/v1/common/loginout"
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
	params := url.Values{}
	params.Set("name", loginInParams.Name)
	params.Set("mobile", loginInParams.Mobile)
	if req,err:=userHttpClient.NewRequest("GET",LOGIN_URL,"",params);err!=nil {
		log.Warn("getaway login fial err=%s",err)
		c.JSON(nil, ecode.RequestErr)
		return
	}else {
		req.Header.Set("x-seesion-token",c.Request.Header.Get("x-seesion-token"))
		resp:=&model.LoginResponse{}
		if err:=userHttpClient.JSON(c,req,resp);err!=nil {
			log.Warn("getaway login JSON fial err=%s",err)
			c.JSON(nil, ecode.RequestErr)
			return
		}else {
			c.JSON(resp,nil)
		}

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

	params := url.Values{}
	params.Set("name", mLoginOutSystem.Name)
	params.Set("mobile", mLoginOutSystem.Mobile)
	params.Set("token", mLoginOutSystem.Token)
	if req,err:=userHttpClient.NewRequest("GET",LOGIN_OUT_URL,"",params);err!=nil {
		log.Warn("getaway login fial  ")
		c.JSON(nil, ecode.RequestErr)
		return
	}else {
		req.Header.Set("x-seesion-token",c.Request.Header.Get("x-seesion-token"))
		resp:=&model.LoginResponse{}
		if err:=userHttpClient.JSON(c,req,resp);err!=nil {
			log.Warn("getaway login JSON fial err=%s",err)
			c.JSON(nil, ecode.RequestErr)
			return
		}else {
			c.JSON(resp,nil)
		}

	}
}


/**
一是 Request Body 就是整个文件内容，通过请求头（即 Header ）中的 Content-Type 字段来指定文件类型。
二是用 multipart 表单方式来上传
*/
//上传信息照片
func updatecard(c *blademaster.Context) {

	p := new(model.ParamUpload)
	if err := c.Bind(p); err != nil {
		c.JSON(nil, ecode.ReqParamErr)
		return
	}

	c.Request.ParseMultipartForm(32 << 10)


	params := url.Values{}
	params.Set("id_no", p.IdNo)
	params.Set("user_real_name",p.UserRealName)
	params.Set("card_id", p.CradId)
	params.Set("age", p.Age)
	params.Set("sex", p.Sex)
	if req,err:=userHttpClient.NewRequest("POST",USER_UPLOAD_RUL,"",params);err!=nil {
		log.Warn("getaway login fial  ")
		c.JSON(nil, ecode.RequestErr)
		return
	}else {
		//req.MultipartForm.File =
		req.Header.Set("Content-Type", binding.MIMEMultipartPOSTForm)
		req.Header.Set("x-seesion-token",c.Request.Header.Get("x-seesion-token"))
		resp:=&model.LoginResponse{}
		if err:=userHttpClient.JSON(c,req,resp);err!=nil {
			log.Warn("getaway login JSON fial err=%s",err)
			c.JSON(nil, ecode.RequestErr)
			return
		}else {
			c.JSON(resp,nil)
		}

	}

}