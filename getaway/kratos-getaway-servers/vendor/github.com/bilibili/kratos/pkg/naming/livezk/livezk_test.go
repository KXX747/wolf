package livezk

import (
	"testing"
	"net"
	"fmt"
	"github.com/bilibili/kratos/pkg/net/ip"
	"github.com/bilibili/kratos/pkg/naming"
	"context"
	"time"
	"path"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/bilibili/kratos/pkg/log"

)

const (

	discoveryID = "kratos_server"
)

var(

	zkconfig *Zookeeper //zkconfig
	zconn *livezk//zkconn
	rpcaddr ="0.0.0.0:6634" //注册到zk的节点
)


func init()  {
	//var zkconfig Zookeeper
	//获取ftp配置
	zkconfig =&Zookeeper{
		Addrs:[]string{"192.168.57.134:2181","192.168.57.135:2181","192.168.57.136:2181"},
		Root:"/microservice",
		Timeout:1* time.Second,
	}

	var err error
	zconn,err=NewZookeeper(zkconfig)
	if err != nil {
		fmt.Println("zk 初始化失败。。。",err)
		return
	}else {
		fmt.Println(" zk初始化成功 。。。",zconn)
	}


	//for{
	//	select {
	//	case event :=<-zconn.zkEvent:
	//		//event事件监听
	//		log.Info("zk event: err: %s, path: %s, server: %s, state: %s, type: %s",
	//			event.Err, event.Path, event.Server, event.State.String(), event.Type.String())
	//	default:
	//
	//	}
	//}

}


func TestS(t *testing.T) {

	//FlagEphemeral := 1
	FlagSequence := 2
	fmt.Println(FlagSequence&FlagSequence == FlagSequence)

	//return func() {
	//	fmt.Println(FlagSequence&FlagSequence==FlagSequence)
	//}

}

func TestZK(t *testing.T)  {
	//获取本机ip
	internalIP:=ip.InternalIP()
	//获取rpc地址的端口
	_, port, err := net.SplitHostPort(rpcaddr)
	if err != nil {
		log.Info("获取rpc端口失败。。。err=%s",err)
		return
	}

	ins:=&naming.Instance{
		AppID:discoveryID,
		Addrs:[]string{fmt.Sprintf("grpc://%s:%s", internalIP, port)},
	}

	//将服务注册到zk中
	var cancelzk context.CancelFunc = func() {}
	_,err=zconn.Register(context.Background(),ins)
	log.Info("z.Register err=%s",err)
	if err!=nil {
		cancelzk()
		time.Sleep(1 * time.Second)
	}

	time.Sleep(time.Second*1000000)

}

//获取注册数据
func TestGetZKPath(t *testing.T)  {
	var stat *zk.Stat
	var data []byte
	var err error

	appid:="kratos_server/0_"
	nodePath := path.Join(zkconfig.Root, basePath, appid)
	data,stat,err=zconn.zkConn.Get(nodePath)
	if err!=nil {
		log.Info("zconn.zkConn.Get data=%s err=%s",string(data),err)
		return
	}else {
		log.Info("zconn.zkConn.Get data=%s stat=%d",string(data),stat.Aversion)
	}
}

//创建节点
func TestCreateZKPath(t *testing.T)  {
	//appid
	appid:="kratos_server/midder_server"
	data :=[]byte("我是33年的请求数据")
	nodePath := path.Join(zkconfig.Root, basePath, appid)

	//创建root
	//if err :=createAll(nodePath,0); err != nil {
	//	//fmt.Println("TestCreateZKPath createAll err=",err)
	//	return
	//}


	//0 持久节点（PERSISTENT）
	//1 临时节点（EPHEMERAL）
	//2 持久顺序节点（PERSISTENT_SEQUENTIAL） 带顺序编号 2_0000000007
	//3 临时顺序节点(EPHEMERAL_SEQUENTIAL)带顺序编号 33_0000000009
	/**
	 ls /microservice/live/service/kratos_server/33_
	 get /microservice/live/service/kratos_server
	 */
	for i:=0;i<10;i++ {
		nodePath="/microservice/live/lock/kratos_server/200"
		path,err:=zconn.zkConn.Create(nodePath,data,2, zk.WorldACL(zk.PermAll))
		if err!=nil {
			log.Info("zconn.zkConn.Create nodePath=%s data=%s err=%s",nodePath,data,err)
			return
		}else {
			log.Info("zconn.zkConn.Create nodePath=%s data=%s path=%s",nodePath,data,path)
		}
		//time.Sleep(time.Second*1)
		fmt.Println("节点关闭",zconn.Close())
	}
}


func TestExistsW(t *testing.T)  {
	appid:="kratos_server"
	nodePath := path.Join(zkconfig.Root, lock, appid)
	isExists, stat,evnts, err:=zconn.zkConn.ExistsW(nodePath)
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	fmt.Println("isExists =",isExists ," Stat = ",stat," error=",err)

	for{

		select {
		case event:=<-evnts:
			fmt.Println("event =",event)
		}
	}
}


/**
基于ZooKeeper分布式锁的流程

    在zookeeper指定节点（locks）下创建临时顺序节点node_n
    获取locks下所有子节点children
    对子节点按节点自增序号从小到大排序
    判断本节点是不是第一个子节点，若是，则获取锁；若不是，则监听比该节点小的那个节点的删除事件
    若监听事件生效，则回到第二步重新进行判断，直到获取到锁

https://blog.csdn.net/zhangcongyi420/article/details/84204153
 */

 //加锁
func TestZkLock(t *testing.T)  {
	name:="/200"
	appid:="kratos_server"
	root := path.Join(zkconfig.Root, lock, appid)
	nodePath:=root+name
	//fmt.Println(nodePath)

	zkMutex:=&ZkMutex{
			Name:name,
			ctx:context.TODO(),
			flags:3,
			acl:zk.WorldACL(zk.PermAll),
			path:nodePath,
			root:root}



	//go func() {
	//	for i:=0;i<2;i++ {
	//		zkMutex.lockZk()
	//		//设置执行时间
	//		time.Sleep(1*time.Second)
	//		zkMutex.unlockZk()
	//	}
	//}()

	//for i:=10000;i<30000;i++ {
	//	zkMutex.lockZk()
	//	//设置执行时间
	//	time.Sleep(10*time.Millisecond)
	//	zkMutex.unlockZk()
	//}
	//
	//go func() {
	//	for i:=0;i<10000;i++ {
	//		zkMutex.lockZk()
	//		//设置执行时间
	//		time.Sleep(5*time.Millisecond)
	//		zkMutex.unlockZk()
	//	}
	//}()

	zkconfig =&Zookeeper{
		Addrs:[]string{"192.168.57.134:2181","192.168.57.135:2181","192.168.57.136:2181"},
		Root:"/microservice",
		Timeout:1* time.Second,
	}

	var err error
	zconn,err=NewZookeeper(zkconfig)
	if err != nil {
		fmt.Println("zk 初始化失败。。。",err)
		return
	}else {
		fmt.Println(" zk初始化成功 。。。",zconn)
	}

	NewZkMutex(zconn,zkMutex)
	for i:=0;i<10;i++ {
		zkMutex.lockZk()
		//设置执行时间
		//time.Sleep(1*time.Second)
		fmt.Println("我是分布式锁。。。。")
		zkMutex.unlockZk()
	}
	time.Sleep(80*time.Second)

}

//添加临时节点
func BenchmarkAddNode(b *testing.B) {

	for i:=0;i<1;i++ {
		nodePath:="/microservice/live/lock/kratos_server/%d"
		nodePath=fmt.Sprintf(nodePath,(201+i))
		_,err:=zconn.zkConn.Create(nodePath,nil,3, zk.WorldACL(zk.PermAll))
		if err!=nil {
			log.Info("zconn.zkConn.Create nodePath=%s err=%s",nodePath,err)
			return
		}
		time.Sleep(5*time.Millisecond)

	}

}