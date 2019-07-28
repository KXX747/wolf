package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
	"errors"
)

func TestFromContext(t *testing.T) {
	report := &mockReport{}
	t1 := newTracer("service1", report, &Config{DisableSample: true})
	fmt.Println("t1=",t1)
	sp1 := t1.New("test123")
	fmt.Println("sp1=",sp1)
	ctx := context.Background()
	fmt.Println("ctx =",ctx)
	ctx = NewContext(ctx, sp1)
	fmt.Println("ctx=",ctx)
	sp2, ok := FromContext(ctx)
	if !ok {
		t.Fatal("nothing from context")
	}

	fmt.Println(sp2)
	assert.Equal(t, sp1, sp2)
}

func TestGrpcCarrier_Get(t *testing.T) {

	p:=genID()

	fmt.Println(p)
}

func TestNewContext(t *testing.T) {
	err:= errors.New("sql address not nil...")
	c :=context.TODO()
	var tj Trace
	var  ok bool
	if tj, ok = FromContext(c); ok {
		fmt.Println("t t= ",tj, " err = ",err)
		tj = tj.Fork("sql.client", "ping")
		tj.SetTag(String(TagAddress, "db.conf.Addr"), String("trace.TagComment", ""))
		defer tj.Finish(&err)
	}

	fmt.Println("t = ",tj, " err = ",err)




}