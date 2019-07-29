package log

/**
"unix", "unixgram", "unixpacket":
 */
import (
	"fmt"
	xtime "github.com/bilibili/kratos/pkg/time"
	"sync"
	"context"
	"github.com/bilibili/kratos/pkg/log/internal/core"
	"net"
	"time"
	"strconv"
	//"log"
	"github.com/bilibili/kratos/pkg/net/trace"
	"github.com/bilibili/kratos/pkg/net/metadata"
)

const (
	megreWait =  1 * time.Second
	MinLogContent = 6  * 1024
	MaxLogContent = 6 * 1024 * 1024
	chanCount = 1024
)

//配置信息
type AgentLogConfig struct {
	// Network is grpc listen network,default value is tcp
	Network string //`dsn:"network"`
	// Addr is grpc listen addr,default value is 0.0.0.0:9000
	Addr string //`dsn:"address"`
	//
	Timeout xtime.Duration //`dsn:"query.timeout"`

	Chan uint;
	Buffer  int64
}

//
type HandlerAgentLog struct {
	ChanLogs  chan []core.Field //go+net
	Waiter    sync.WaitGroup //保护数据
	Pool      sync.Pool // pool保存field
	Config	  *AgentLogConfig
	FilterData map[string]struct{}
	enc       core.Encoder
} 


//new agent
func NewLogAgent(mAgentLogConfig *AgentLogConfig)(handlerAgentLog *HandlerAgentLog){

	handlerAgentLog= &HandlerAgentLog{
		Config:mAgentLogConfig,
		enc:core.NewJSONEncoder(core.EncoderConfig{
			EncodeTime:core.EpochTimeEncoder,
			EncodeDuration:core.SecondsDurationEncoder,
		},core.NewBuffer(0)),
	}

	//new pool
	handlerAgentLog.newPool()
	//chan int
	handlerAgentLog.Config.Chan = chanCount
	//create buffer chan
	handlerAgentLog.ChanLogs=make(chan[]core.Field,handlerAgentLog.Config.Chan)

	if mAgentLogConfig.Buffer ==0 {
		handlerAgentLog.Config.Buffer = 100
	}

	handlerAgentLog.Waiter.Add(1)

	//
	KV(_appID,c.Family).AddTo(handlerAgentLog.enc)

	//send
	go handlerAgentLog.writeToServer()

	return
}

//创建pool
func (agent *HandlerAgentLog)newPool ()  {
	//
	agent.Pool.New = func() interface{} {
		return  make([]core.Field, 0, 16)
	}
}
//获取
func (agent *HandlerAgentLog)getPool()(fields []core.Field) {
	//
	fields=	agent.Pool.Get().([]core.Field)
	return
}
//
func (agent *HandlerAgentLog)putPool (putData []core.Field)  {
	//释放
	putData =putData[0:0]
	agent.Pool.Put(putData)
}


//append log to write to server
func(agent *HandlerAgentLog)Log(ctx context.Context,loglevel Level,args...D)  {


	p:=agent.getPool()
	//fmt.Println("Log ",args,"   p ",p)

	for index:=range args {
		value:=args[index]
		p=append(p,value)
	}

	if t,ok:=trace.FromContext(ctx);ok {
		if s, ok := t.(fmt.Stringer); ok {
			p = append(p, KV(_tid, s.String()))
		} else {
			p = append(p, KV(_tid, fmt.Sprintf("%s", t)))
		}
	}
	if caller := metadata.String(ctx, metadata.Caller); caller != "" {
		p = append(p, KV(_caller, caller))
	}
	if color := metadata.String(ctx, metadata.Color); color != "" {
		p = append(p, KV(_color, color))
	}
	if cluster := metadata.String(ctx, metadata.Cluster); cluster != "" {
		p = append(p, KV(_cluster, cluster))
	}
	if metadata.Bool(ctx, metadata.Mirror) {
		p = append(p, KV(_mirror, true))
	}
	//
	p = append(p, KV(_appID, c.Family))

	select {
	case agent.ChanLogs <-p:
	default:

	}
}

//write to server
func(agent *HandlerAgentLog)writeToServer()  {

	var (
		conn  *net.UnixConn
		err   error
		//count int
		quit  bool
	)
	buf := core.NewBuffer(MinLogContent)
	tick := time.NewTicker(megreWait)
	defer agent.Waiter.Done()

	for  {
		select {
		case p :=<- agent.ChanLogs:
			//log.Println(" agent.ChanLogs p=%s ",buf.Len())
			if p==nil {
				quit=true
				goto NEW
			}

			now := time.Now()
			//buf.Write([]byte("1000"))
			buf.Write([]byte(strconv.FormatInt(now.UnixNano()/1e6, 10)))
			agent.enc.Encode(buf,p...)
			agent.putPool(p)

			if conn != nil&&buf.Len()>=MinLogContent {
				//上传日志
				go func(conn  *net.UnixConn,buffer *core.Buffer) {
					//log.Printf("客户端收集的数据。。。",string(buf.Bytes()))
					if _,err=conn.Write(buf.Bytes());err!=nil {
						//log.Printf("conn.Write err=%s buf=%f",err,buf.Len())
						conn.Close()
						//log.Printf("上传日志失败。。。")
					}else {
						buf.Reset()
					//	log.Printf("上传日志成功。。。")
					}
				}(conn,buf)
			}

		case <-tick.C:

		}
		//fmt.Println("准备连接服务。。。。。")
		conf :=agent.Config
		unixAddr, _ := net.ResolveUnixAddr(conf.Network, conf.Addr)
		if conn, err = net.DialUnix(conf.Network, nil, unixAddr);err!=nil{
			//log.Printf("net.DialUnix err=%s addr=%s network=%s",err,conf.Addr,conf.Network)
			continue
		}
		//fmt.Println("。。。。。。。。。")
	}

	NEW:
		if conn != nil &&buf.Len()>0{
			//fmt.Println("客户端收集的数据。。。",string(buf.Bytes()))
			if _,err=conn.Write(buf.Bytes());err!=nil {
				//log.Printf("conn.Write err=%s buf=%f",err,buf.Len())
				conn.Close()
			//	log.Printf("上传日志失败。。。")
			}else {
				buf.Reset()
			//	log.Printf("上传日志成功。。。")
			}

		}
		if quit {
			if conn != nil && err == nil {
				conn.Close()
			}
			return
		}
}




//close chan wait
func(agent *HandlerAgentLog)Close()(error){
	agent.ChanLogs<-nil
	agent.Waiter.Wait()
	return  nil
}


// SetFormat set render format on log output
// see StdoutHandler.SetFormat for detail
func(agent *HandlerAgentLog)SetFormat(string){

}




