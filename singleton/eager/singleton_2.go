package eager

// 单例模式饿汉写法二
type singleton_2 struct {
	name string
}

var instance_2 *singleton_2

// 在包初始化的时候执行
func init(){
	instance_2 = &singleton_2{}
}

func GetSingleton_2() *singleton_2{
	return instance_2
}



