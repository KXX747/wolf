package liveconsul

import (
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/hashicorp/consul/api"
	"time"
	"encoding/json"
	"context"
	"path"
	"net"
	"strconv"
	"fmt"
	"strings"
	"net/url"
)


/**
consul服务治理和分布式锁，服务配置
 */

const(
	//microservice/live/service/kratos_server/10.100.62.235:6634
	//服务注册
	basePath = "/live/service"
	configPath = "/live/config"
	scheme   = "grpc"
)

//consul server
type ConsulServer struct {
	Root string
	Addr []string
	MonitorAddr string
	Interval string
	Timeout string //超时时间
	ServerName string//服务名称
	ServerNameId string //服务唯一id
	Tags []string//tag标记 dev等
}


// 服务注册
type liveConsul struct {
	consulClient *api.Client
	consulConfig *ConsulServer
	lastIndex uint64//
}

//配置
type KVS struct {
	consulClient *api.Client
	Key string //存储的key
	Value []byte //数据
	Flags uint64
}

//服务注册的连接
func NewConsul(mConsulServer *ConsulServer) (liveconsul *liveConsul){
	liveconsul = &liveConsul{
		consulConfig:mConsulServer,
	}
	return
}


type consulIns struct {
	Group       string `json:"group"`
	LibVersion  string `json:"lib_version"`
	StartupTime string `json:"startup_time"`
}

//保存到consul的数据
func newZkInsData(ins *naming.Instance) ([]byte, error) {
	zi := &consulIns{
		// TODO group support
		Group:       "default",
		LibVersion:  ins.Version,
		StartupTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(zi)
}

//发现服务,获取服务的list
func (l *liveConsul)DiscoverServicer(ctx context.Context)(err error,serviceNameList []*api.AgentService){

		consulConf:=api.DefaultConfig()
		consulConf.Address =l.consulConfig.Addr[0]
		kv:=l.NewKV()
		//获取client
		l.consulClient,err=api.NewClient(consulConf)
		if err!=nil {
			return
		}
		kv.consulClient = l.consulClient

		//获得service
		var  services map[string][]string
		services, _, err=kv.consulClient.Catalog().Services(&api.QueryOptions{})
		if err!=nil {
			return
		}
		//清理获取服务
		for name:=range services {
			//获取健康存在的服务
			var mServiceEntry []*api.ServiceEntry
			var metainfo *api.QueryMeta
			mServiceEntry, metainfo, err = kv.consulClient.Health().Service(name, "", true,
				&api.QueryOptions{
					//	WaitIndex:l.lastIndex,// 同步点，这个调用将一直阻塞，直到有新的更新
				})

			l.lastIndex = metainfo.LastIndex

			//获取服务名称
			for _, entry := range mServiceEntry {
				if l.consulConfig.ServerName != entry.Service.Service {
					continue
				}

				for _, health := range entry.Checks {
					if health.ServiceName != l.consulConfig.ServerName {
						continue
					}
					//获取数据
					address := entry.Service.Address
					port := entry.Service.Port
					serviceNameId := health.ServiceID
					fmt.Println("  health nodeid:", health.Node, " service_name:", health.ServiceName, " service_id:", serviceNameId, " status:", health.Status, " ip:", address, " port:", port)
					//保存
					serviceNameList = append(serviceNameList, entry.Service)
				}
			}
		}
	return
}


//regist data to consul
func (l *liveConsul)Register(ctx context.Context, ins *naming.Instance) (cancel context.CancelFunc, err error){
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
	//获取rpc地址，拼接serveriname和serverinameid
	var host,port string
	host, port, err =net.SplitHostPort(rpc)
	nodepath :=path.Join(l.consulConfig.Root, basePath,ins.AppID)
	l.consulConfig.ServerNameId =nodepath
	p,_:=strconv.Atoi(port)

	//注册的值
	var value []byte
	value,err=newZkInsData(ins)

	fmt.Println("value = ","   " ,value,l.consulConfig.ServerName,"   " ,l.consulConfig.ServerNameId)

	//注册到consul的数据
	servirce:=&api.AgentServiceRegistration{
		ID:l.consulConfig.ServerNameId,
		Name:l.consulConfig.ServerName,
		Port:p,
		Tags:l.consulConfig.Tags,
		Address:host,
		Check:&api.AgentServiceCheck{
			HTTP:l.consulConfig.MonitorAddr,
			Interval:l.consulConfig.Interval,
			Timeout:l.consulConfig.Timeout,
		},
	}
	//添加系统变量
	//os.Setenv(api.HTTPAddrEnvName,l.consulConfig.Addr[0])
	//api.DefaultConfig()会获取api.HTTPAddrEnvName设置的系统变量
	mConfig:=api.DefaultConfig()
	mConfig.Address = l.consulConfig.Addr[0]
	l.consulClient,err=api.NewClient(mConfig)
	if err!=nil {
		return
	}

	if err=l.consulClient.Agent().ServiceRegister(servirce);err!=nil {
		return
	}

	//
	return func() {
		l.unregister()
	},nil

	return
}

//unregister data in consul
func (l *liveConsul) unregister() (err error) {
	//
	err=l.consulClient.Agent().ServiceDeregister(l.consulConfig.ServerName)

	return 
}


//close consul
func (l *liveConsul)Close() (err error){
	l.consulClient=nil
	return
}

//liveConsul获取KV对象
func (l *liveConsul)NewKV()(kv *KVS){
	kv=&KVS{}
	kv.consulClient =l.consulClient
	return
}


//add config
func (kv *KVS)StoreKeyValue()(err error){

	pair:=&api.KVPair{
		Key:kv.nodeCheck(kv.Key),
		Value:kv.Value,
		Flags:kv.Flags,
	}
	_,err=kv.consulClient.KV().Put(pair,nil)

	return
}


//get config
func (kv *KVS)GetKeyValue()(value []byte,err error){
	var k *api.KVPair
	k,_,err=kv.consulClient.KV().Get(kv.nodeCheck(kv.Key),nil)
	if k!=nil {
		value=k.Value
	}
	return
}


//key不能以/开头
func(kv *KVS)nodeCheck(key string)( string)  {

	//将/开头的删除
	index:=strings.Index(key,"/")
	if index>=0 {
		key = key[1:len(key)-1]
	}
	return key
}