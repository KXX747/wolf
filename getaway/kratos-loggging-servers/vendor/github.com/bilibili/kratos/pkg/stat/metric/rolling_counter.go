package metric

import (
	"fmt"
	"time"
)

var _ Metric = &rollingCounter{}
var _ Aggregation = &rollingCounter{}

// RollingCounter represents a ring window based on time duration.
// e.g. [[1], [3], [5]]
type RollingCounter interface {
	Metric
	Aggregation
	// Reduce applies the reduction function to all buckets within the window.
	Reduce(func(Iterator) float64) float64
}

// RollingCounterOpts contains the arguments for creating RollingCounter.
type RollingCounterOpts struct {
	Size           int
	BucketDuration time.Duration
}

//关联rollingPolicy
type rollingCounter struct {
	policy *RollingPolicy
}

// NewRollingCounter creates a new RollingCounter bases on RollingCounterOpts.
func NewRollingCounter(opts RollingCounterOpts) RollingCounter {
	//创建window，指定window的buckets的大小
	window := NewWindow(WindowOpts{Size: opts.Size})
	//创建新的point
	policy := NewRollingPolicy(window, RollingPolicyOpts{BucketDuration: opts.BucketDuration})
	//
	return &rollingCounter{
		policy: policy,
	}
}

func (r *rollingCounter) Add(val int64) {
	if val < 0 {
		panic(fmt.Errorf("stat/metric: cannot decrease in value. val: %d", val))
	}
	r.policy.Add(float64(val))
}

func (r *rollingCounter) Reduce(f func(Iterator) float64) float64 {
	return r.policy.Reduce(f)
}

func (r *rollingCounter) Avg() float64 {
	return r.policy.Reduce(Avg)
}

func (r *rollingCounter) Min() float64 {
	return r.policy.Reduce(Min)
}

func (r *rollingCounter) Max() float64 {
	return r.policy.Reduce(Max)
}

func (r *rollingCounter) Sum() float64 {
	return r.policy.Reduce(Sum)
}

func (r *rollingCounter) Value() int64 {
	return int64(r.Sum())
}
