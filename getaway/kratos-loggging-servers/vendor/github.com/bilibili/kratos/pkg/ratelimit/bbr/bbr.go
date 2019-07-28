package bbr

import (
	"context"
	"math"
	"sync/atomic"
	"time"

	"github.com/bilibili/kratos/pkg/container/group"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	limit "github.com/bilibili/kratos/pkg/ratelimit"
	"github.com/bilibili/kratos/pkg/stat/metric"

	cpustat "github.com/bilibili/kratos/pkg/stat/sys/cpu"

	"fmt"
)

var (
	cpu         int64 //cpu
	decay       = 0.75 //衰变
	defaultConf = &Config{
		Window:       time.Second * 5, //时间
		WinBucket:    50, //桶
		CPUThreshold: 800, //cpu阀值
	}
)

type cpuGetter func() int64

func init() {
	go cpuproc()
}

//250毫秒采集数据，动态计算
func cpuproc() {
	ticker := time.NewTicker(time.Millisecond * 250)
	defer func() {
		ticker.Stop()
		if err := recover(); err != nil {
			log.Error("rate.limit.cpuproc() err(%+v)", err)
			go cpuproc()
		}
	}()

	// EMA algorithm: https://blog.csdn.net/m0_38106113/article/details/81542863
	for range ticker.C {
		stat := &cpustat.Stat{}
		cpustat.ReadStat(stat)
		prevCpu := atomic.LoadInt64(&cpu)
		//fmt.Println("prevCpu = ",prevCpu," float64(stat.Usage) = ",float64(stat.Usage))
		curCpu := int64(float64(prevCpu)*decay + float64(stat.Usage)*(1.0-decay ))
		//fmt.Println("curCpu = ",curCpu ,"  *decay ",float64(prevCpu)*decay," *(1.0-decay) = ",float64(stat.Usage)*(1.0-decay))
		atomic.StoreInt64(&cpu, curCpu)
		//fmt.Println("atomic.StoreInt64(&cpu, curCpu) = ",atomic.LoadInt64(&cpu))
		//fmt.Println("")
	}
}

// Stats contains the metrics's snapshot of bbr.
/**
|cpu|最近 1s 的 CPU 使用率均值，使用滑动平均计算，采样周期是 250ms|
|inflight|当前处理中正在处理的请求数量|
|pass|请求处理成功的量|
|rt|请求成功的响应耗时|
 */
type Stat struct {
	Cpu         int64
	InFlight    int64
	MaxInFlight int64
	MinRt       int64
	MaxPass     int64
}

// BBR implements bbr-like limiter.
// It is inspired by sentinel.
// https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81
type BBR struct {
	cpu             cpuGetter
	passStat        metric.RollingCounter //统计请求成功
	rtStat          metric.RollingGauge //统计请求所花费的时间
	inFlight        int64
	winBucketPerSec int64
	conf            *Config
	prevDrop        time.Time
}

// Config contains configs of bbr limiter.
type Config struct {
	Enabled      bool
	Window       time.Duration
	WinBucket    int
	Rule         string
	Debug        bool
	CPUThreshold int64
}

//计算请求处理成功的量
func (l *BBR) maxPASS() int64 {
	val := int64(l.passStat.Reduce(func(iterator metric.Iterator) float64 {
		var result = 1.0
		for iterator.Next() {
			bucket := iterator.Bucket()
			count := 0.0
			for _, p := range bucket.Points {
				count += p
			}
			result = math.Max(result, count)
		}
		return result
	}))
	if val == 0 {
		return 1
	}
	return val
}

/**
计算每次请求使用的时间=总请求时间/总请求个数
 */
func (l *BBR) minRT() int64 {
	val := l.rtStat.Reduce(func(iterator metric.Iterator) float64 {
		var result = math.MaxFloat64

		for iterator.Next() {
			bucket := iterator.Bucket()
			if len(bucket.Points) == 0 {
				continue
			}
			total := 0.0
			for _, p := range bucket.Points {
				total += p
			}
			avg := total / float64(bucket.Count)
			result = math.Min(result, avg)
		}
		return result
	})
	return int64(math.Ceil(val))
}

//最大处理的请求数量
func (l *BBR) maxFlight() int64 {
	//100*50*20=10000/1000+0.5 =10.5
	pass:=l.maxPASS()
	rt:=l.minRT()
	finght:=int64(math.Floor(float64(pass*rt*l.winBucketPerSec)/1000.0 + 0.5))
	fmt.Println("maxFlight=",finght)
	return finght
}

//超出限制
func (l *BBR) shouldDrop() bool {
	//获取执行的请求
	inFlight := atomic.LoadInt64(&l.inFlight)
	//获取最大的执行请求在指定时间内=
	maxInflight := l.maxFlight()
	//判断cpu的值大于设置cpu的阀值，并且请求数大于最大的请求数时，结束，超出限制
	if l.cpu() < l.conf.CPUThreshold {
		if time.Now().Sub(l.prevDrop) <= 1000*time.Millisecond {
			return inFlight > 1 && inFlight > maxInflight
		}
		return false
	}
	return inFlight > 1 && inFlight > maxInflight
}

// Stat tasks a snapshot of the bbr limiter.
func (l *BBR) Stat() Stat {

	return Stat{
		Cpu:         l.cpu(), //获取当cpu的使用率
		InFlight:    atomic.LoadInt64(&l.inFlight),//获取当前正在请求的接口
		MinRt:       l.minRT(), //
		MaxPass:     l.maxPASS(),
		MaxInFlight: l.maxFlight(),
	}
}


// Allow checks all inbound traffic.
// Once overload is detected, it raises ecode.LimitExceed error.
func (l *BBR) Allow(ctx context.Context, opts ...limit.AllowOption) (func(info limit.DoneInfo), error) {
	allowOpts := limit.DefaultAllowOpts()
	for _, opt := range opts {
		opt.Apply(&allowOpts)
	}
	if l.shouldDrop() {
		l.prevDrop = time.Now()
		return nil, ecode.LimitExceed
	}
	atomic.AddInt64(&l.inFlight, 1)
	stime := time.Now()
	return func(do limit.DoneInfo) {
		//请求所用的时间
		rt := int64(time.Since(stime) / time.Millisecond)
		l.rtStat.Add(rt)
		atomic.AddInt64(&l.inFlight, -1)
		switch do.Op {
		case limit.Success:
			l.passStat.Add(1)
			return
		default:
			return
		}
	}, nil
}

/**
	Enabled      bool
	Window       time.Duration
	WinBucket    int
	Rule         string
	Debug        bool
	CPUThreshold int64
 */
func newLimiter(conf *Config) limit.Limiter {
	if conf == nil {
		conf = defaultConf
	}
	//fmt.Println("conf=",conf)
	size := conf.WinBucket
	bucketDuration := conf.Window / time.Duration(conf.WinBucket)
	//fmt.Println("bucketDuration=",bucketDuration ,"   conf.Window=",conf.Window ,"  time.Duration(conf.WinBucket)=",time.Duration(conf.WinBucket))
	passStat := metric.NewRollingCounter(metric.RollingCounterOpts{Size: size, BucketDuration: bucketDuration})
	rtStat := metric.NewRollingGauge(metric.RollingGaugeOpts{Size: size, BucketDuration: bucketDuration})
	cpu := func() int64 {
		return atomic.LoadInt64(&cpu)
	}
	limiter := &BBR{
		cpu:             cpu,
		conf:            conf,
		passStat:        passStat,
		rtStat:          rtStat,
		winBucketPerSec: int64(time.Second) / (int64(conf.Window) / int64(conf.WinBucket)),
		prevDrop:        time.Unix(0, 0),
	}

	fmt.Println("limiter = ",limiter.winBucketPerSec)
	//fmt.Println()
	return limiter
}

// Group represents a class of BBRLimiter and forms a namespace in which
// units of BBRLimiter.
type Group struct {
	group *group.Group
}

// NewGroup new a limiter group container, if conf nil use default conf.
func NewGroup(conf *Config) *Group {
	if conf == nil {
		conf = defaultConf
	}
	group := group.NewGroup(func() interface{} {
		return newLimiter(conf)
	})

	return &Group{
		group: group,
	}
}

// Get get a limiter by a specified key, if limiter not exists then make a new one.
func (g *Group) Get(key string) limit.Limiter {
	limiter := g.group.Get(key)
	//fmt.Println("(g *Group) = ",limiter)
	return limiter.(limit.Limiter)
}
