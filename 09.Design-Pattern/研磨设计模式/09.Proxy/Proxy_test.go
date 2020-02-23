package Proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{money: 100}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fatal()
	}
}
