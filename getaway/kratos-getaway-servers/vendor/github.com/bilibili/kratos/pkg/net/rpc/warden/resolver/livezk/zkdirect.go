package livezk
//发现

import (
	"context"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/livezk"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const (
	// Name is the name of direct resolver
	Name = "grpc"
)


var _ naming.Resolver = &DirectZk{}


type DirectZk struct {

	path string //监控的地址
	zkConn *zk.Conn
	zkEvent <-chan zk.Event
	config *livezk.Zookeeper
}

//构造器
func NewDirectZk()(directZk *DirectZk) {

	return &DirectZk{}
}



// New new live zookeeper registry
func NewZookeeper(config *livezk.Zookeeper)( directZk *DirectZk, err  error){

	directZk =&DirectZk{
		config:config,
	}

	directZk.zkConn,directZk.zkEvent,err=zk.Connect(config.Addrs,time.Duration(config.Timeout))
	if err!=nil {
		return nil,err
	}
	return
}



//

func (mDirectZk *DirectZk) Fetch(context.Context) (mInstancesInfo *naming.InstancesInfo, isFound bool) {


	return
}


func (mDirectZk *DirectZk) Watch() <-chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return ch
}

func(mDirectZk *DirectZk) Close() (err error) {

	return
}

