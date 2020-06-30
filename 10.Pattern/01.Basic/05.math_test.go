package main

import (
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	// int32 最大最小值
	t.Log(math.MaxInt32) // 实际值 1<<31-1
	t.Log(math.MinInt32) // 实际值 -1<<31

	// int64 最大最小值
	t.Log(math.MaxInt64) // 实际值 1<<63-1
	t.Log(math.MinInt64) // 实际值 -1<<63
}
