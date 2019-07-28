package ratelimit

import (
	"testing"
	"fmt"
)

func TestDefaultAllowOpts(t *testing.T) {

	allowOptions:=DefaultAllowOpts()

	fmt.Println(allowOptions)
}