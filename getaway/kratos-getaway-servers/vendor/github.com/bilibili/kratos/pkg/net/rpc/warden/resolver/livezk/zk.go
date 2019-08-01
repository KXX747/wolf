package livezk

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/livezk"
	"github.com/bilibili/kratos/pkg/net/ip"
	"github.com/bilibili/kratos/pkg/log"
	"net"
	"path"
	"time"
)

const Path  ="/live/service"

// register appid
func Register(config *livezk.Zookeeper, addr string, appId string) (context.CancelFunc, error) {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	z, err := livezk.NewZookeeper(config)
	if err != nil {
		return nil, err
	}
	internalIP := ip.InternalIP()
	ins := &naming.Instance{
		AppID: appId,
		Addrs: []string{fmt.Sprintf("grpc://%s:%s", internalIP, port)},
	}
	return z.Register(context.Background(), ins)
}

//discovery appid
func Discovery(config *livezk.Zookeeper, appId string)(resultChan chan map[string]struct{}){

	zkcli,err:=NewZkClient(config.Addrs,time.Duration(config.Timeout))
	if err!=nil {
		log.Warn("Discovery NewZkClient 连接节点失败 err=%s",err)
		panic(err)
	}

	//组建path
	rootPath:=path.Join(config.Root,Path,appId)

	resultChan = make(chan map[string]struct{}, 1)

	result,zkevent,err:=zkcli.WatchChildren(rootPath)
	if err!=nil {
		log.Warn("Discovery WatchChildren 监听节点失败 err=%s",err)
		panic(err)
	}
	resultChan <- result

	for{
		select {
		case event:=<-zkevent:

			fmt.Println("log event=%s",event)
		}


	}

	return

}