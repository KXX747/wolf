package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/http"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/trace"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/dao"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/common"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/loggings"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/tracings"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	if err:=paladin.Watch("app.toml",dao.Conf);err!=nil {
		panic(err)
		return
	}

	go loggings.NewUnix(dao.Conf)
	log.Init(dao.Conf.Log) // debug flag: log.dir={path}
	defer log.Close()

	go tracings.NewUnix(dao.Conf)
	trace.Init(dao.Conf.Tracer)
	defer trace.Close()
	log.Info("kratos-loggging-servers start")
	svc := service.New()
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			common.RemoveFilePath(svc.AppConfig.Log.Agent.Addr)
			common.RemoveFilePath(svc.AppConfig.Tracer.Addr)
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("kratos-loggging-servers exit")
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
