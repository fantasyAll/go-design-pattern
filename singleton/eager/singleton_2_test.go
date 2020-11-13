package eager

import (
	"sync"
	"testing"
)

func TestGetSingleton_2(t *testing.T) {
	s1 := GetSingleton_2()
	s1.name = "hh"
	s2 := GetSingleton_2()
	s2.name = "xx"
	t.Log(&s1.name, s1)
	t.Log(&s2.name, s2)
	t.Logf("s1:%p %T", s1, s1)
	t.Logf("s2:%p %T", s2, s2)
	if s1 != s2 {
		t.Fatal("不一致！")
	}
}

func TestParallelSingleton_2(t *testing.T) {
	// 控制获取实例开始
	start := make(chan struct{})
	// 控制获取实例结束
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton_2{}
	// 获取实例
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetSingleton_2()
			wg.Done()
		}(i)
	}
	// 开始并行获取实例
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Logf("i：%d %p %T", i, instances[i], instances[i])
			t.Logf("i：%d %p %T", i-1, instances[i-1], instances[i-1])
		}
	}
}
