package main

import (
	"net/url"
	"fmt"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"log"
	"time"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {

	rawdsn:="unixgram:///var/run/dapper-collect/dapper-collect.sock"
	u, _ := url.Parse(rawdsn)
	fmt.Println(u.Scheme,"<><>",u.Path)


	/**
2019/07/17 19:57:23 Initializing logging reporter
foo3 =  Hello foo3
2019/07/17 19:57:23 Reporting span 2f32714bb6490712:58534296ed7ef426:2f32714bb6490712:1
foo4 =  Hello foo4
foo =  foo3Reply foo4Reply
2019/07/17 19:57:24 Reporting span 2f32714bb6490712:c784eacc32ffc02:2f32714bb6490712:1
span 上传完成
2019/07/17 19:57:24 Reporting span 2f32714bb6490712:2f32714bb6490712:0:1
	 */
	//config:=&trace_jaeger.JgConfig{
	//	Addr:"192.168.57.134:6831",
	//	ServiceName:"jaeger-demo",
	//}
	//trace_jaeger.NewConfigMain(config)
	//fmt.Println("span 上传完成")

	//m()

	tracerTime()
}

type T struct {

	Name string
	Time     int64
	Duration int64
}


//time 分析
/**
 fmt.Println(time.Now().Unix()) //获取当前秒
    fmt.Println(time.Now().UnixNano())//获取当前纳秒
    fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒
    fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒
    c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型
    fmt.Println(c.String()) //输出当前英文时间戳格式

int64到time(将纳秒转time)
tt := time.Unix(0,1515049539324129700) //将纳秒转换为 time 类型
fmt.Println(tt.String())

int64到time(将毫秒转time)
tt := time.Unix(0,毫秒*1e6) //将纳秒转换为 time 类型
fmt.Println(tt.String())

int64到time(将秒转time)
tt := time.Unix(1136214245,0) //将秒转换为 time 类型
fmt.Println(tt.String())

作者：吃猫的鱼0
链接：https://www.jianshu.com/p/f2434fa75f70
来源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
 */
func tracerTime()  {

	t1:=&T{Name:"span_root_wanshang",Time:1563604665750488,Duration:1005956}
	t2:=&T{Name:"span_foo4",Time:1563604666252276,Duration:504382}
	t3:=&T{Name:"span_foo3",Time:1563604665750499,Duration:501350}


	t1Time:=time.Unix(t1.Time/1e6,0)
	t2Time:=time.Unix(t2.Time/1e6,0)
	t3Time:=time.Unix(t3.Time/1e6,0)

	fmt.Println("t1Time=",t1Time)
	fmt.Println("t2Time=",t2Time)
	fmt.Println("t3Time=",t3Time)

	startTime :=time.Now()

	time.Sleep(time.Microsecond*500)

	duration:=time.Now().Sub(startTime)

	fmt.Println(duration<=1000*time.Millisecond )
}

func m()  {

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831", // 替换host
		},
	}
	closer, err := cfg.InitGlobalTracer(
		"serviceName",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	var ctx = context.TODO()
	span1, ctx := opentracing.StartSpanFromContext(ctx, "span_1")
	span1.SetTag("key_1","value_1")
	time.Sleep(time.Second / 2)
	span11, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second / 2)
	span11.Finish()
	span1.Finish()
	defer closer.Close()
}







