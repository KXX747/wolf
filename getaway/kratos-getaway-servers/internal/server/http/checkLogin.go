package http

import (
	"github.com/KXX747/wolf/getaway/kratos-getaway-servers/internal/model"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/middleware/auth"
	"net/http"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func CheckLogin(c *bm.Context) bm.HandlerFunc {

	auth.IsLogin(c)
	return func(c *bm.Context) {

	}
}


func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!! docker",
	}
	log.Info("Golang 大法好 !!! ip=%s",c.Request.RemoteAddr)
	c.JSON(k, nil)
}
