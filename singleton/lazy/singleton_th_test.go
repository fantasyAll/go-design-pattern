package lazy

import (
	"sync"
	"testing"
)

func TestGetSingleton_th(t *testing.T) {
	s1 := GetSingleton_th()
	s1.name = "jesse"
	s2 := GetSingleton_th()
	s2.name = "tom"
	t.Log(&s1.name, s1.name)
	t.Log(&s2.name, s2.name)
	t.Logf("s1：%p %T", s1, s1)
	t.Logf("s2：%p %T", s2, s2)
	if s1 != s2 {
		t.Fatal("不一致！")
	}
}

// 测试并行声明
func TestParallelSingleton(t *testing.T) {
	// 控制获取实例化开始
	start := make(chan struct{})
	// 控制获取实例化结束
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton_th{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetSingleton_th()
			wg.Done()
		}(i)
	}
	// 关闭chan开始获取实例化
	close(start)
	// 阻塞
	wg.Done()
	// 判断所有的实例是否一致
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Logf("i：%d %p", i, instances[i])
			t.Logf("i-1：%d %p", i-1, instances[i-1])
			t.Fatal("不一致！")
		}
	}
}
