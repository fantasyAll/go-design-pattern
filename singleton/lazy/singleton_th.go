package lazy

import "sync"

// 单例模式 懒汉模式线程安全版写法之一
type singleton_th struct {
	name string
}

var instance_th *singleton_th
var once sync.Once

func GetSingleton_th() *singleton_th{
	// sync.Once 可以保证只执行一次
	once.Do(func(){
		instance_th = &singleton_th{}
	})
	return instance_th
}