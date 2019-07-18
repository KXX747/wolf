package tracings

import (
	"time"
	"sync"
	"github.com/KXX747/wolf/getaway/kratos-loggging-servers/internal/dao"
	)
const (
	megreWait = 1 * time.Second
	trace_bufsize= 32 * 1024
	trace_workers = 10;
)

/**
trace追踪系统转换为jaeger数据
 */


 //服务端接收数据
type TraceService struct {
	AppConfig 		*dao.Config
	pool 			sync.Pool
	wg  			sync.WaitGroup
	workers     	int
}


func (mTraceService *TraceService) buffer() []byte {
	return mTraceService.pool.Get().([]byte)
}

func (mTraceService *TraceService) freeBuffer(p []byte) {
	mTraceService.pool.Put(p)
}

//数据在jaeger显示
type OpenTracingServer struct {

}
