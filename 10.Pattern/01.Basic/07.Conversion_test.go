package main

import (
	"strconv"
	"testing"
)

func TestTypeConversion(t *testing.T) {
	// byte 转数字
	s := "12345"                    // s[0] 类型是 byte
	num := int(s[0] - '0')          // 1
	str := string(s[0])             // "1"
	b := byte(num + '0')            // '1'
	t.Logf("%d %s %c", num, str, b) // 1 1 1

	// 字符串与数字转换
	num, _ = strconv.Atoi(s)
	t.Logf("%T", num) // int
	str = strconv.Itoa(num)
	t.Logf("%T", str) // string
}
