package lazy

import (
	"sync"
	"testing"
)

// 并行次数
const parCount = 233

func TestGetSingleton_nth(t *testing.T) {
	s1 := GetSingleton_nth()
	s1.name = "xx"
	s2 := GetSingleton_nth()
	s2.name = "zz"
	t.Log(&s1.name, s1)
	t.Log(&s2.name, s2)
	t.Logf("s1：%p %T", s1, s1)
	t.Logf("s2：%p %T", s2, s2)
	if s1 != s2 {
		t.Fatal("不一致!")
	}
}

// 测试并行
func TestParallelSingleton_nth(t *testing.T) {
	// 控制获取实例开始
	start := make(chan struct{})
	// 控制获取实例结束
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton_nth{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetSingleton_nth()
			// 表示执行完了
			wg.Done()
		}(i)
	}
	// 结束管道开始获取实例
	close(start)
	// 阻塞程序，等待获取实例完成
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Logf("i：%d %p", i, instances[i])
			t.Logf("i-1：%d %p", i-1, instances[i-1])
			t.Fatal("不一致!")
		}
	}
}
