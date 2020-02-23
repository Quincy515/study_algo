package Builder

import (
	"fmt"
	"testing"
)

func TestBuildString(t *testing.T) {
	builder := &StringBuilder{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		fmt.Printf("Builder fail expect 123 acture %s", res)
	}
}

func TestBuildInt64(t *testing.T) {
	builder := &IntBuilder{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		fmt.Printf("Builder fail expect 123 acture %d", res)
	}
}
