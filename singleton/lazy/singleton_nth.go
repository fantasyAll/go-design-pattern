package lazy

// 单例模式 懒汉版非线程安全写法
type singleton_nth struct{
	name string
}

var instance *singleton_nth

func GetSingleton_nth()*singleton_nth{
	if instance == nil{
		instance = &singleton_nth{}
	}
	return instance
}



