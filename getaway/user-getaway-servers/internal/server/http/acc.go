package http

import (
	"github.com/KXX747/wolf/getaway/user-getaway-servers/internal/model"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"context"
)


// example for http request handler.
func howToStart(c *blademaster.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}


/**
添加用户
 */
func addUser(c *blademaster.Context)  {

	p :=new(model.ParamAddUser)

	if err:=c.Bind(p);err!=nil {
		return
	}

	resp,err:=svc.AddUser(context.TODO(),p.Name,p.Mobile)
	if err!=nil {
		c.JSON(nil,err)
	}else {
		c.JSON(resp, nil)
	}

}

/**
删除用户
 */
func deleteUser(c *blademaster.Context)  {

	p :=new(model.ParamDeleteUser)

	if err:=c.Bind(p);err!=nil {
		return
	}

	resp,err:=svc.DeleteUser(context.TODO(),p.IdNo,p.Content)
	if err!=nil {
		c.JSON(nil,err)
	}else {
		c.JSON(resp, nil)
	}

}


/**
更新用户
 */
func updateUser(c *blademaster.Context)  {

	p :=new(model.ParamUpdateUser)

	if err:=c.Bind(p);err!=nil {
		return
	}

	resp,err:=svc.UpdateUser(context.TODO(),p.IdNo,p.Name,p.Mobile,p.Address)
	if err!=nil {
		c.JSON(nil,err)
	}else {
		c.JSON(resp, nil)
	}

}


