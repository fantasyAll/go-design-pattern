package lazy

import "testing"

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
