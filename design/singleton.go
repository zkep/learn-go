package design

// 懒汉式

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
