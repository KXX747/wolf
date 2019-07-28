package metric

import (
	"math/rand"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	counter := NewCounter(CounterOpts{})
	count := rand.Intn(100)
	for i := 0; i < count; i++ {
		counter.Add(1)
	}
	val := counter.Value()
	fmt.Println(val, count)
	assert.Equal(t, val, int64(count))

}
