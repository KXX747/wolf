package tracings

import (
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/server/common"
	"github.com/bilibili/kratos/pkg/log"
	"net"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/dao"
	"sync"
	"fmt"
	"github.com/golang/protobuf/proto"
	protogen "github.com/bilibili/kratos/pkg/net/trace/proto"
)
//init
func NewUnix(svc *dao.Config)(mTraceService *TraceService){
	mTraceService = &TraceService{
		AppConfig:svc,
		pool:sync.Pool{
			New: func() interface{} {
				return make([]byte,trace_workers)
			},
		},
		workers:trace_workers,

	}
	//初始化协议
	mTraceService.initUnix(svc)

	return
}

func(mTraceService *TraceService) initUnix(svc  *dao.Config) {

	agent := mTraceService.AppConfig.Tracer
	var unixListener *net.UnixListener
	//var err error
	//svc.AppConfig.Log.Agent.Addr
	//err:=common.ListenUNIX(agent.Addr)
	//if err!=nil {
	//	log.Info("agent start log server err=%s addr=%s",err,agent.Addr)
	//	return
	//}
	common.RemoveFilePath(agent.Addr)
	//ticker:=time.NewTicker(megreWait)
	log.Info("agent addr=%s  peoto=%s timeout=%s", agent.Addr, agent.Network, agent.Timeout)

	unixAddr, _ := net.ResolveUnixAddr(agent.Network, agent.Addr)
	unixListener, _ = net.ListenUnix(agent.Network, unixAddr)
	defer unixListener.Close()

	for {
		unixConn, err := unixListener.AcceptUnix()
		//unixConn.SetDeadline(time.Now().Add())
		if err != nil {
			log.Info("unixListener.AcceptUnix err=%s",err)
			continue
		}
		//go unixPipe(unixConn,mTraceService)
		go unixPipeByte(unixConn)
	}
}

func unixPipeByte(conn *net.UnixConn) {
	//ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" )
		conn.Close()
	}()

	for {
		buf:=make([]byte,1024)

		bufLen,_,err:=conn.ReadFrom(buf)
		if err != nil {
			log.Info("traces server ReadFrom err",err.Error())
			return
		}
		if bufLen<=0 {
			return
		}
		fmt.Println()
		protoSpan := new(protogen.Span)
		content:=buf[:]
		proto.Unmarshal(content,protoSpan)
		fmt.Println("trace message= ",protoSpan)
		buf =buf[0:0]
		fmt.Println()
	}
}

/**
处理数据数据
*/
func unixPipe(conn *net.UnixConn,mTraceService *TraceService) {
	defer func() {
		conn.Close()
	}()
	for {
		p:=mTraceService.buffer()
		defer mTraceService.freeBuffer(p)
		n, _, err :=conn.ReadFrom(p)
		if err!=nil {
			log.Info("traces server ReadFrom err",err.Error())
			return
		}
		if n>0 {
			content :=p[:]
			protoSpan := new(protogen.Span)
			proto.Unmarshal(content,protoSpan)
			fmt.Println("trace message= ",protoSpan)
		}
		//回复客户端
		conn.Write([]byte("\n"))
		//time.Sleep(time.Second)
	}
}


/**
opentracing上传服务器数据：
version:1 operation_name:"Redis:GET" trace_id:4562668293271138332 span_id:6268248949528539164
parent_id:93271138332 start_time:<seconds:1563434045 nanos:367193000 > duration:<nanos:936655 >
tags:<key:"region" value:"region01" >
tags:<key:"zone" value:"zone01" >
tags:<key:"hostname" value:"747-2.local" >
tags:<key:"ip" value:"10.100.62.235" >
tags:<key:"span.kind" value:"client" >
tags:<key:"span.kind" value:"client" >
tags:<key:"component" value:"pkg/cache/redis" >
tags:<key:"peer.service" value:"redis" >
tags:<key:"peer.address" value:"192.168.57.136:19000" >
tags:<key:"db.statement" value:"GET account:user:c7e760fb-19a0-43d2-9265-be863704c77a" >


opentracing上传服务器数据：
version:1 operation_name:"queryrow" trace_id:4562668293271138332 span_id:484785961719115804
parent_id:45271138332 start_time:<seconds:1563434045 nanos:368744000 > duration:<nanos:8851374 >
tags:<key:"region" value:"region01" >
tags:<key:"zone" value:"zone01" >
tags:<key:"hostname" value:"747-2.local" >
tags:<key:"ip" value:"10.100.62.235" >
tags:<key:"span.kind" value:"client" >
tags:<key:"legacy.address" value:"192.168.57.110:18066" >
tags:<key:"legacy.comment" value:"select id,name,id_no,mobile,address,create_at,create_ip,create_by from user where id_no=? limit 1" >


opentracing上传服务器数据：
version:1 operation_name:"Redis:SET" trace_id:4562668293271138332 span_id:6216402105785987100 parent_id:93271138332
start_time:<seconds:1563434045 nanos:378213000 > duration:<nanos:818329 > tags:<key:"region" value:"region01" >
tags:<key:"zone" value:"zone01" >
tags:<key:"hostname" value:"747-2.local" >
tags:<key:"ip" value:"10.100.62.235" >
tags:<key:"span.kind" value:"client" >
tags:<key:"span.kind" value:"client" >
tags:<key:"component" value:"pkg/cache/redis" >
tags:<key:"peer.service" value:"redis" >
tags:<key:"peer.address" value:"192.168.57.136:19000" >
tags:<key:"db.statement" value:"SET account:user:c7e760fb-19a0-43d2-9265-be863704c77a" >

opentracing上传服务器数据：
version:1 operation_name:"/account.service.Users/FindUser" trace_id:4562668293271138332 span_id:45626682332
start_time:<seconds:1563434045 nanos:366501000 > duration:<nanos:12942201 >
tags:<key:"region" value:"region01" >
tags:<key:"zone" value:"zone01" >
tags:<key:"hostname" value:"747-2.local" >
tags:<key:"ip" value:"10.100.62.235" >
tags:<key:"span.kind" value:"server" >
tags:<key:"component" value:"net/http" >
tags:<key:"http.method" value:"GET" >
tags:<key:"http.url" value:"/account.service.Users/FindUser?id_no=c7e760fb-19a0-43d2-9265-be863704c77a" >
tags:<key:"span.kind" value:"server" >
tags:<key:"caller" >
 */


 /**
 trace message=  version:1 operation_name:"Redis:GET" trace_id:371447714717446245 span_id:59246413122187365 parent_id:371447714717446245 start_time:<seconds:1563452880 nanos:704528000 > duration:<nanos:1994286 > tags:<key:"region" value:"region01" > tags:<key:"zone" value:"zone01" > tags:<key:"hostname" value:"747-2.local" > tags:<key:"ip" value:"192.168.1.101" > tags:<key:"span.kind" value:"client" > tags:<key:"span.kind" value:"client" > tags:<key:"component" value:"pkg/cache/redis" > tags:<key:"peer.service" value:"redis" > tags:<key:"peer.address" value:"192.168.57.136:19000" > tags:<key:"db.statement" value:"GET account:user:c7e760fb-19a0-43d2-9265-be863704c77a" >


trace message=  version:1 operation_name:"queryrow" trace_id:371447714717446245 span_id:8333405337253916773 parent_id:371447714717446245 start_time:<seconds:1563452880 nanos:706637000 > duration:<nanos:4588060 > tags:<key:"region" value:"region01" > tags:<key:"zone" value:"zone01" > tags:<key:"hostname" value:"747-2.local" > tags:<key:"ip" value:"192.168.1.101" > tags:<key:"span.kind" value:"client" > tags:<key:"legacy.address" value:"192.168.57.110:18066" > tags:<key:"legacy.comment" value:"select id,name,id_no,mobile,address,create_at,create_ip,create_by from user where id_no=? limit 1" >


trace message=  version:1 operation_name:"Redis:SET" trace_id:371447714717446245 span_id:9165681080510787685 parent_id:371447714717446245 start_time:<seconds:1563452880 nanos:711411000 > duration:<nanos:973435 > tags:<key:"region" value:"region01" > tags:<key:"zone" value:"zone01" > tags:<key:"hostname" value:"747-2.local" > tags:<key:"ip" value:"192.168.1.101" > tags:<key:"span.kind" value:"client" > tags:<key:"span.kind" value:"client" > tags:<key:"component" value:"pkg/cache/redis" > tags:<key:"peer.service" value:"redis" > tags:<key:"peer.address" value:"192.168.57.136:19000" > tags:<key:"db.statement" value:"SET account:user:c7e760fb-19a0-43d2-9265-be863704c77a" >


trace message=  version:1 operation_name:"/account.service.Users/FindUser" trace_id:371447714717446245 span_id:371447714717446245 start_time:<seconds:1563452880 nanos:703974000 > duration:<nanos:8685472 > tags:<key:"region" value:"region01" > tags:<key:"zone" value:"zone01" > tags:<key:"hostname" value:"747-2.local" > tags:<key:"ip" value:"192.168.1.101" > tags:<key:"span.kind" value:"server" > tags:<key:"component" value:"net/http" > tags:<key:"http.method" value:"GET" > tags:<key:"http.url" value:"/account.service.Users/FindUser?id_no=c7e760fb-19a0-43d2-9265-be863704c77a" > tags:<key:"span.kind" value:"server" > tags:<key:"caller" >


  */