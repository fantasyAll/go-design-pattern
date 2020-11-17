package eager

// go 单例模式饿汉写法一
type singleton_1 struct {
	name string
}

// 直接构建实例，在包加载的时候就会构建
var instance_1 *singleton_1 = &singleton_1{}

func GetSingleton_1() *singleton_1 {
	return instance_1
}
