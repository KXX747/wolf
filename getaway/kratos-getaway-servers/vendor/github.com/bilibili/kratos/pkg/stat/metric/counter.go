package metric

import (
	"fmt"
	"sync/atomic"
)

var _ Metric = &counter{}

/**
Counter

Counter用于累计值，例如记录请求次数、任务完成数、错误发生次数。一直增加，不会减少。重启进程后，会被重置。

例如：http_response_total{method=”GET”,endpoint=”/api/tracks”}
100，10秒后抓取http_response_total{method=”GET”,endpoint=”/api/tracks”} 100。
https://studygolang.com/articles/12874
https://blog.csdn.net/wuxiaobingandbob/article/details/78954406
 */
// Counter stores a numerical value that only ever goes up.
type Counter interface {
	Metric
}

// CounterOpts is an alias of Opts.
type CounterOpts Opts

type counter struct {
	val int64
}

// NewCounter creates a new Counter based on the CounterOpts.
func NewCounter(opts CounterOpts) Counter {
	return &counter{}
}

func (c *counter) Add(val int64) {
	if val < 0 {
		panic(fmt.Errorf("stat/metric: cannot decrease in negative value. val: %d", val))
	}
	atomic.AddInt64(&c.val, val)
}

func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
