package loggings

import (
	"fmt"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/dao"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/common"
	"github.com/bilibili/kratos/pkg/log"
	"net"
	"sync"
	"time"
)
/**
日志收集
 */

const (
	megreWait = 1 * time.Second
	log_bufsize= 32 * 1024
)

type LogService struct {
	AppConfig *dao.Config
	pool sync.Pool
} 

//init
func NewUnix(svc *dao.Config)(mLogService *LogService){
	mLogService = &LogService{
		AppConfig:svc,
		pool:sync.Pool{
			New: func() interface{} {
				return make([]byte,log_bufsize)
			},
		},
	}
	//初始化协议
	 mLogService.initUnix(svc)

	return
}

func(mLogService *LogService) initUnix(svc *dao.Config) {

	agent := mLogService.AppConfig.Log.Agent
	var unixListener *net.UnixListener
	var err error
	//svc.AppConfig.Log.Agent.Addr
	common.RemoveFilePath(agent.Addr)
	//err=common.ListenUNIX(agent.Addr)
	//if err!=nil {
	//	log.Info("agent start log server err=%s addr=%s",err,agent.Addr)
	//	return
	//}

	//ticker:=time.NewTicker(megreWait)
	log.Info("agent addr=%s  peoto=%s timeout=%s", agent.Addr, agent.Network, agent.Timeout)
	unixAddr, _ := net.ResolveUnixAddr(agent.Network, agent.Addr)
	unixListener, err = net.ListenUnix(agent.Network, unixAddr)
	if err!=nil {
		log.Info("loggings server start fial=%s", err)
		return
	}
	defer unixListener.Close()

	for {
		unixConn, err := unixListener.AcceptUnix()
		if err != nil {
			log.Info("logging unixListener.AcceptUnix err=%s",err)
			continue
		}
		go unixPipe(unixConn,mLogService)
	}
}

/**
处理数据数据
*/
func unixPipe(conn *net.UnixConn,mLogService *LogService) {
	//agent:= mLogService.AppConfig.Log.Agent
	defer func() {
		conn.Close()
	}()
	for {
		p:=mLogService.buffer()
		defer mLogService.freeBuffer(p)
		//conn.SetReadDeadline(time.Now().Add())
		n, _, err :=conn.ReadFrom(p)
		if err!=nil {
			log.Info("loggings servers ReadFrom err",err.Error())
			return
		}
		if n>0 {
			//解析opentracing数据
			content:=p[:]

		}
	}
}

func (mLogService *LogService) buffer() []byte {
	return mLogService.pool.Get().([]byte)
}

func (mLogService *LogService) freeBuffer(p []byte) {
	mLogService.pool.Put(p)
}
