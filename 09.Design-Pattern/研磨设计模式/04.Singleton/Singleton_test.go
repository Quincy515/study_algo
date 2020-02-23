package Singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	i1 := GetInstance()
	i2 := GetInstance()
	if i1 == i2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不等")
	}

}
