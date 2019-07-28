package redis

import (
	"testing"
	"time"
	"context"
	"fmt"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/bilibili/kratos/pkg/container/pool"
	"net"
	"strconv"
)
var p *Pool
var config *Config

func init() {
	config = getConfig()
	//p =redis.NewPool(config)
	p=NewPool(config)

}

func getConfig() (c *Config) {
	c = &Config{
		Name:         "test",
		Proto:        "tcp",
		Addr:         "192.168.57.135:19000",
		DialTimeout:  xtime.Duration(time.Second),
		ReadTimeout:  xtime.Duration(time.Second),
		WriteTimeout: xtime.Duration(time.Second),
	}
	c.Config = &pool.Config{
		Active:      20,
		Idle:        2,
		IdleTimeout: xtime.Duration(90 * time.Second),
	}
	return
}



/**
计算redis的时间
 */
func TestTime(t *testing.T)  {
	start:=time.Now()
	fmt.Println("start = ",start)

	f:=time.Duration(int64(float64(DefaultExpiry)*DefaultFactor))
	fmt.Println("f = ",f)

	su:=time.Now().Sub(start)
	fmt.Println("Sub now = ",time.Now(),"  su = ",su)

	until := time.Now().Add(DefaultExpiry - su -f  + 2*time.Millisecond)
	fmt.Println("time.Now() =",time.Now()," until = ",until)


	fmt.Println(" until = ",until)
	fmt.Println(" until = ",until)

	fmt.Println( time.Unix(0, 0))
}

func BenchmarkRedislock(b *testing.B) {

	for i:=0;i<b.N;i++ {
		fmt.Println()
		//lock1("woshi1")
		//lock2("woshi2")
		lock3(fmt.Sprintf("woshi%d",i))
		//add("woshi1")
		fmt.Println("i= ",i," ",add("woshi1"))
	}
}



//redis锁
func TestRedislock(t *testing.T)  {

	//lock1("woshi1")
	//lock2("woshi2")
	//lock3("woshi3")

	host, port, err := net.SplitHostPort("192.168.57.180:8080")

	fmt.Println("host=",host," port=",port," err=",err)

}

func add(name string) (n string){
	ctx:=context.TODO()
	rml,_:=NewMutex(ctx,name,p)
	err:=rml.Lock()
	fmt.Println("rml.Lock()",err)

	ni:=199+99+9;
	n=strconv.Itoa(ni)
	n=n+name

	err=rml.Unlock()
	fmt.Println("rml.Unlock()",err)
	return
}



func lock3(name string)  {
	ctx:=context.TODO()
	rml,_:=NewMutex(ctx,name,p)
	err:=rml.Lock()
	fmt.Println("rml.Lock()",err)
	time.Sleep(10*time.Millisecond)
	//err=rml.Unlock()
	fmt.Println("rml.Unlock()",err)
}

func lock2(name string)  {
	ctx:=context.TODO()
	rml,_:=NewMutex(ctx,name,p)
	err:=rml.Lock()
	fmt.Println("rml.Lock()",err)
	time.Sleep(2*time.Second)
	err=rml.Unlock()
	fmt.Println("rml.Unlock()",err)
}

func lock1(name string)  {
	ctx:=context.TODO()
	rml,_:=NewMutex(ctx,name,p)
	err:=rml.Lock()
	fmt.Println("rml.Lock()",err)
	time.Sleep(1*time.Second)

	err=rml.Unlock()
	fmt.Println("rml.Unlock()",err)
}
