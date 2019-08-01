package livezk
/**
zoookeeper 服务治理和分布式锁，服务配置
 main location
 */

 /**
  [liveZK]
    root = "/microservice/resource-service/"
    addrs = ["172.16.33.54:2181"]
    timeout = "30s"
  */

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/samuel/go-zookeeper/zk"
	xtime "github.com/bilibili/kratos/pkg/time"
	"net/url"
	"path"
	"strings"
	"time"
)

const(
	//microservice/live/service/kratos_server/10.100.62.235:6634
	//服务注册
	basePath = "/live/service"
	scheme   = "grpc"
	//分布式锁
	//microservice/live/lock/kratos_server/key01
	lock = "/live/lock"
)


//zookeeper server 
type Zookeeper struct {
	Root string
	Addrs []string
	Timeout xtime.Duration
}

// livezk live service zookeeper registry
type livezk struct {
	zkConfig *Zookeeper
	zkConn   *zk.Conn
	zkEvent  <-chan zk.Event
}


// New new live zookeeper registry
func NewZookeeper(config *Zookeeper)(*livezk,error){

	lz :=&livezk{
		zkConfig:config,
	}
	var err error
	lz.zkConn,lz.zkEvent,err=zk.Connect(lz.zkConfig.Addrs,time.Duration(lz.zkConfig.Timeout))
	if err!=nil {
		//
		go lz.eventproc()
	}
	return lz,err
}


//type zkIns struct {
//	Group       string `json:"group"`
//	//LibVersion  string `json:"lib_version"`
//	StartupTime string `json:"startup_time"`
//	*naming.Instance
//
//}

func newZkInsData(ins *naming.Instance) ([]byte, error) {
	//zi := &zkIns{
	//	// TODO group support
	//	Group:       "default",
	//	//LibVersion:  ins.Version,
	//	StartupTime: time.Now().Format("2006-01-02 15:04:05"),
	//}
	ins.LastTs = time.Now().Unix()
	return json.Marshal(ins)
}


//注册到zk中
func (l *livezk)Register(ctx context.Context, ins *naming.Instance) (cancel context.CancelFunc, err error){
	//
	nodePath := path.Join(l.zkConfig.Root, basePath, ins.AppID)
	if err = l.createAll(nodePath); err != nil {
		return
	}

	//遍历获取rpc地址
	var rpc string
	for _,addr:=range ins.Addrs {
		var url *url.URL
		url,err=url.Parse(addr)
		if url != nil && url.Scheme == scheme {
			rpc = url.Host
			break
		}
	}
	if rpc == "" {
		err = errors.New("no GRPC addr")
		log.Info(" rpc ==null   err=%s",err)

		return
	}
	dataPath := path.Join(nodePath, rpc)
	var data []byte
	data, err =newZkInsData(ins)
	if err != nil {
		return nil, err
	}
	//zk注册零时节点
	_, err=l.zkConn.Create(dataPath,data,zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return nil, err
	}

	//将zk close 通过参数提供出去
	return func() {
		l.unregister(dataPath)
	}, nil
}

//创建服务节点
func (l *livezk) createAll(nodePath string) (err error) {
	seps := strings.Split(nodePath, "/")
	lastPath := "/"
	ok := false
	for _, part := range seps {
		if part == "" {
			continue
		}
		lastPath = path.Join(lastPath, part)
		if ok, _, err = l.zkConn.Exists(lastPath); err != nil {
			return err
		} else if ok {
			continue
		}
		if _, err = l.zkConn.Create(lastPath, nil, 0, zk.WorldACL(zk.PermAll)); err != nil {
			return
		}
	}
	return
}


//删除
func (l *livezk) unregister(dataPath string) error {
	return l.zkConn.Delete(dataPath, -1)
}


//
func (l *livezk)Close() (err error){
	l.zkConn.Close()
	return
}

func (l *livezk) eventproc() {
	//for event := range l.zkEvent {
	//	// TODO handle zookeeper event
	//	log.Info("zk event: err: %s, path: %s, server: %s, state: %s, type: %s",
	//		event.Err, event.Path, event.Server, event.State, event.Type)
	//}

	for {

		select {
		case event :=<-l.zkEvent:
			log.Info("zk event: err: %s, path: %s, server: %s, state: %s, type: %s",
				event.Err, event.Path, event.Server, event.State, event.Type)
		default:

		}
	}
}

//func a() {
//	chan1 := make(chan int)
//	chan2 := make(chan int)
//	//select基本用法
//	select {
//	case <-chan1:
//		// 如果chan1成功读到数据，则进行该case处理语句
//	case chan2 <- 1:
//		// 如果成功向chan2写入数据，则进行该case处理语句
//	default:
//		// 如果上面都没有成功，则进入default处理流程
//	}
//}