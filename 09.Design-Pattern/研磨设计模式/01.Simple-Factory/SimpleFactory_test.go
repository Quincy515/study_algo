package SimpleFactory

import "testing"

func TestChinese(t *testing.T) {
	api := NewAPI("zh_cn")
	s := api.Say("小明")
	if s != "你好小明" {
		t.Fatal("Chinese test Fail")
	}
}

func TestEnglish(t *testing.T) {
	api := NewAPI("en")
	s := api.Say("Tom")
	if s != "hello Tom" {
		t.Fatal("Chinese test Fail")
	}
}
