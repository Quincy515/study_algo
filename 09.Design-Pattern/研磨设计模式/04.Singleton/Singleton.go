package Singleton

import "sync"

// Singleton: 单例模式类
type Singleton struct {
	data int
}

var singleton *Singleton

// 只允许运行一次
var once sync.Once // 内核信号，只能运行一次

// GetInstance: 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() { singleton = &Singleton{100} }) // 相等
	//singleton = &Singleton{100} // 不等
	return singleton
}
