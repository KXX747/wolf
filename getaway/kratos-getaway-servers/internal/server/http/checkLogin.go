package http

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/middleware/auth"
)

func CheckLogin(c *blademaster.Context) blademaster.HandlerFunc {

	auth.IsLogin(c)
	return func(c *blademaster.Context) {

	}
}
