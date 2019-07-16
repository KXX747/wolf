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
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/dao"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/unix"
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

	g:=dao.Conf.Log.Agent
	log.Info("Log  addr=%s  protol=%s timeout=%s",g.Addr,g.Network,g.Timeout)

	log.Init(dao.Conf.Log) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("kratos-loggging-servers start")
	svc := service.New()
	httpSrv := http.New(svc)
	unix.NewUnix(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
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
