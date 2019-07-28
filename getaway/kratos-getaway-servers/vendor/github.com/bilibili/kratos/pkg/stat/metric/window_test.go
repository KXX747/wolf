package metric

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"time"
	"github.com/bilibili/kratos/pkg/log"
)

func TestWindowResetWindow(t *testing.T) {

	mWindow:= time.Second * 5
	mWinBucket:=time.Duration(50)

	bucketDuration:=mWindow/mWinBucket
	fmt.Println("bucketDuration =",bucketDuration)

	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	log.Info("window,size=%d ",window.size)
	//遍历window
	for index,w:=range  window.window  {
		log.Info("index=%d Points=%f count=%d ",index,w.Points,w.Count)
	}
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1.0)
	}
	log.Info("")

	//遍历window
	for index,w:=range  window.window  {
		log.Info("index=%d Points=%f count=%d ",index,w.Points,w.Count)
	}

	window.Add(0,12)


	//遍历window
	for index,w:=range  window.window  {
		log.Info("index=%d Points=%f count=%d ",index,w.Points,w.Count)
	}
	fmt.Println(window)
	//释放window所有的bucket，数据清空
	window.ResetWindow()
	fmt.Println(window)
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, len(window.Bucket(i).Points), 0)
	}
}

func TestWindowResetBucket(t *testing.T) {
	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1.0)
	}
	fmt.Println(window)
	//清楚window下标为1的bucket，由于buckets的size为3，所以存在下标为0和2的points和count有值，下标为1的被清空了
	window.ResetBucket(1)
	fmt.Println(window)
	assert.Equal(t, len(window.Bucket(1).Points), 0)
	assert.Equal(t, window.Bucket(0).Points[0], float64(1.0))
	assert.Equal(t, window.Bucket(2).Points[0], float64(1.0))
}

func TestWindowResetBuckets(t *testing.T) {
	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1.0)
	}
	//根据指定数据的值，删除window的bucket数据
	window.ResetBuckets([]int{0, 1, 2})
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, len(window.Bucket(i).Points), 0)
	}
}

func TestWindowAppend(t *testing.T) {
	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	//window中的bucket添加points为1.0,count为1
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1.0)
	}
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, window.Bucket(i).Points[0], float64(1.0))
	}
}

func TestWindowAdd(t *testing.T) {
	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	//添加值
	window.Append(0, 1.0)
	//计算值，当bucket中不存在值时，添加到第一位，存在就取出值+val
	window.Add(0, 1.0)
	assert.Equal(t, window.Bucket(0).Points[0], float64(2.0))
}

func TestWindowSize(t *testing.T) {
	opts := WindowOpts{Size: 3}
	window := NewWindow(opts)
	assert.Equal(t, window.Size(), 3)
}
