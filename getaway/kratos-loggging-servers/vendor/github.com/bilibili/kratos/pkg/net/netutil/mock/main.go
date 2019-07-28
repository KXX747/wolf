package main

import (
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/net/netutil/breaker"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

func main() {
	g:=ExampleGroup()
	a:=g.Get("a")
	fmt.Println(a)

	t :=time.NewTicker(time.Microsecond)
	for  {
		select {
			case c:=<-t.C:
				c.Second()
				if err:=a.Allow();err!=nil {
					a.MarkFailed()
					fmt.Println("MarkFailed")
					return
				}else {

					a.MarkSuccess()
					//fmt.Println("MarkSuccess")
				}
				err:=deal()
				fmt.Println(err)

			default:
		}
	}
}

func deal()error{
	fmt.Println("NewTicker c = deal")
  //  time.Sleep(time.Second*10)
return errors.New("ssssssssssssssssssss")
}



// ExampleGroup show group usage.
func ExampleGroup() *breaker.Group{
	c := &breaker.Config{
		Window:  xtime.Duration(1 * time.Second),
		Sleep:   xtime.Duration(100 * time.Millisecond),
		Bucket:  10,
		Ratio:   0.5,
		Request: 1000,
	}
	// init default config
	breaker.Init(c)
	// new group
	g := breaker.NewGroup(c)
	// reload group config
	//c.Bucket = 100
	//c.Request = 200
	g.Reload(c)
	//// get breaker by key
	//bkr:=g.Get("key")
	//return bkr

	return g
}

// ExampleBreaker show breaker usage.
func ExampleBreaker() {
	// new group,use default breaker config
	g := breaker.NewGroup(nil)
	brk := g.Get("key")
	// mark request success
	brk.MarkSuccess()
	// mark request failed
	brk.MarkFailed()
	// check if breaker allow or not
	if brk.Allow() == nil {
		fmt.Println("breaker allow")
	} else {
		fmt.Println("breaker not allow")
	}
}

// ExampleGo this example create a default group and show function callback
// according to the state of breaker.
func ExampleGo() {
	run := func() error {
		return nil
	}
	fallback := func() error {
		return fmt.Errorf("unknown error")
	}
	if err := breaker.Go("example_go", run, fallback); err != nil {
		fmt.Println(err)
	}
}
