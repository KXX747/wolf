package liveconsul

import (
	"testing"
	"fmt"
	"net"
	"time"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/ip"
	"context"
	"path"
)

const(
	discoveryID = "rpc_server"
	rpcaddr ="0.0.0.0:6634" //注册到zk的节点
)

var(
	mliveconsul *liveConsul
	mConsulServer *ConsulServer
)

func init() {

	mConsulServer=&ConsulServer{
		Addr:[]string{"192.168.57.134:8500","192.168.57.135:8500","192.168.57.136:8500"},
		Root:"/microservice",
		//MonitorAddr:"192.168.57.136:8501",
		MonitorAddr:"http://192.168.57.180:80/",
		Interval:"5s",
		Timeout:"1s",
		Tags:[]string{"DEV","krator-server","go","regsit"},
	}
	mliveconsul=NewConsul(mConsulServer)
	fmt.Println("liveconsu=",mliveconsul.consulConfig)



}

//初始化consul连接
func TestNewConsul(t *testing.T) {
	fmt.Println("TestNewConsul....")

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
	//注册服务
	ctx:=context.TODO()
	var cancelzk context.CancelFunc = func() {}
	_,err=mliveconsul.Register(ctx,ins)
	if err!=nil {
		log.Info("z.Register err=%s",err)
		cancelzk()
		time.Sleep(1 * time.Second)
	}else {
		log.Info("z.Register success")
	}

	fmt.Println()



	//select {}

	time.Sleep(30 * time.Second)
}

func TestDiscover(t *testing.T)  {

	ctx:=context.TODO()
	//获取服务
	nodepath :=path.Join(mliveconsul.consulConfig.Root, basePath,discoveryID)
	mliveconsul.consulConfig.ServerName ="1000_woshi"
	mliveconsul.consulConfig.ServerNameId =nodepath
	err,list:=mliveconsul.DiscoverServicer(ctx)
	if err!=nil {
		fmt.Println("DiscoverServicer err=",err)
	}
	fmt.Println("list=",list)

	TestUnregister(t)
	//TestKVKeyValue(t)
	time.Sleep(30 * time.Second)
}

func TestUnregister(t *testing.T)  {

	mliveconsul.unregister()
}


//添加和查询配置
func TestKVKeyValue(t *testing.T) {

	if mliveconsul==nil||mliveconsul.consulClient==nil {
		fmt.Println("请先初始化连接。。。")
		return
	}

	key:=path.Join("/microservice",configPath,"krotas_server","nginx")
	value:=[]byte("我是1111的value")


	//添加
	store:=mliveconsul.NewKV()
	store.Key=key
	store.Value=value
	store.Flags=42
	err:=store.StoreKeyValue()
	if err!=nil {
		fmt.Println("store.Err ",err)
		return
	}

	//查询
	query:=mliveconsul.NewKV()
	query.Key=key
	v,err:=query.GetKeyValue()
	if err!=nil {
		fmt.Println("query = err",err)
	}else {

		fmt.Println("query =",string(v))
	}
}



