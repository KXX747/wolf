package trace_jaeger

import (
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go"
	"log"
	"context"
	"time"
	"fmt"
	"io"
	"github.com/opentracing/opentracing-go"

)

type JgConfig struct {
	Addr 			string
	ServiceName 	string
}

type  JaegerServer struct {
	config *JgConfig

}

func NewConfig(mJgConfig *JgConfig)(mJaegerServer *JaegerServer){

	mJaegerServer=&JaegerServer{
		config:mJgConfig,
	}

	cfg:=config.Configuration{
		Sampler:&config.SamplerConfig{
			Type:jaeger.SamplerTypeConst,
			Param:1,
		},
		Reporter:&config.ReporterConfig{
			LogSpans:true,
			LocalAgentHostPort:mJgConfig.Addr,
		},
	}

	tracer,err:=cfg.InitGlobalTracer(mJgConfig.ServiceName)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}


	ctx:=context.TODO()

	span1,ctx:=opentracing.StartSpanFromContext(ctx,"span_1")
	time.Sleep(time.Second / 2)
	span2, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second *1)
	span2.Finish()
	span1.Finish()
	defer tracer.Close()
	return
}

/**
GOROOT=/usr/local/Cellar/go/1.11.4/libexec #gosetup
GOPATH=/Users/a747/go #gosetup
/usr/local/Cellar/go/1.11.4/libexec/bin/go build -i -o /private/var/folders/n9/d2_kgzm56fg6gc9723jnk3q80000gn/T/___go_build_main_go /Users/a747/go/src/github.com/bilibili/kratos/pkg/main.go #gosetup
/private/var/folders/n9/d2_kgzm56fg6gc9723jnk3q80000gn/T/___go_build_main_go #gosetup
unixgram <><> /var/run/dapper-collect/dapper-collect.sock
2019/07/20 14:37:45 Initializing logging reporter
foo3 =  Hello foo3
2019/07/20 14:37:46 Reporting span 5c15adca979b6a76:32c2f49951547f3b:5c15adca979b6a76:1
foo4 =  Hello foo4
foo =  foo3Reply foo4Reply
2019/07/20 14:37:46 Reporting span 5c15adca979b6a76:4c9dc9cc616b65a:5c15adca979b6a76:1
span 上传完成
2019/07/20 14:37:46 Reporting span 5c15adca979b6a76:5c15adca979b6a76:0:1
 */
func NewConfigMain(mJgConfig *JgConfig) (mJaegerServer *JaegerServer){
	mJaegerServer=&JaegerServer{
		config:mJgConfig,
	}

	tracer, closer := initJaeger(mJgConfig)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)//StartspanFromContext创建新span时会用到

	span := tracer.StartSpan("span_root_wanshang")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo3("Hello foo3", ctx)
	r2 := foo4("Hello foo4", ctx)
	fmt.Println("foo = ",r1, r2)
	span.Finish()

	return
}
func initJaeger(mJgConfig *JgConfig) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort:mJgConfig.Addr,
		},
	}
	tracer, closer, err := cfg.New(mJgConfig.ServiceName, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}



func foo3(req string, ctx context.Context) (reply string){
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo3")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println("foo3 = ",req)
	//2.模拟处理耗时
	time.Sleep(time.Second/2)
	//3.返回reply
	reply = "foo3Reply"
	return
}
//跟foo3一样逻辑
func foo4(req string, ctx context.Context) (reply string){
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo4")
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println("foo4 = ",req)
	time.Sleep(time.Second/2)
	reply = "foo4Reply"
	return
}




