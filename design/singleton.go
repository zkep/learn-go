package design

import "sync"

// 饿汉式
type Singleton struct{}

var singleton *Singleton

// init 初始化
func init() {
	singleton = &Singleton{}
}

// 获取实例
func GetInstance() *Singleton {
	return singleton
}

// 懒汉式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
