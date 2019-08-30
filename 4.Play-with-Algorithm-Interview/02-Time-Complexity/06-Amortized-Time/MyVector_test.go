package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestMyVector(t *testing.T) {
	for i := 10; i <= 26; i++ {
		n := int(math.Pow(float64(2), float64(i)))

		startTime := time.Now()
		vec := NewMyVector()
		for num := 0; num < n; num++ {
			vec.push_back(num)
		}
		fmt.Print(n, " operations: \t") // 2 倍关系
		fmt.Println(time.Now().Sub(startTime))
	}
}
