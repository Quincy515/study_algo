package Adapter

import (
	"fmt"
	"testing"
)

var expect = "adaptee method SpecficRequest"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee() // 适配器
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		fmt.Printf("expect: %s, actual: %s", expect, res)
	}
}
