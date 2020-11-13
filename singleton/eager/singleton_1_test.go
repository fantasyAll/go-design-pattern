package eager

import (
	"sync"
	"testing"
)

// 并行次数
const parCount = 233

func TestGetSingleton_1(t *testing.T) {
	s1 := GetSingleton_1()
	s1.name = "s1"
	s2 := GetSingleton_1()
	s2.name = "s2"
	t.Log(&s1.name, s1)
	t.Log(&s2.name, s2)
	t.Logf("s1：%p %T", s1, s1)
	t.Logf("s2：%p %T", s2, s2)
	if s1 != s2 {
		t.Fatal("不一致！")
	}
}

// 大量并行测试
func TestParallelSingleton_1(t *testing.T) {
	// 控制获取实例开始
	start := make(chan struct{})
	// 控制获取实例结束
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton_1{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetSingleton_1()
			wg.Done()
		}(i)
	}
	// 关闭管道，开始执行获取实例
	close(start)
	// 等待执行完成
	wg.Wait()
	// 判断实例是否一致
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Logf("i：%d %p %T", i, instances[i], instances[i])
			t.Logf("i-1：%d %p %T", i-1, instances[i-1], instances[i-1])
			t.Fatal("不一致！")
		}
	}
}
