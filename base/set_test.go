package base

import "testing"

// Map 的 value 可以是一个方法
// 与 Dock type 接口方式一起 可以方便的实现单一方法对象的工厂模式
func TestMapWithFn(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](1), m[2](2), m[3](3))
}

// 使用 Map 实现 Set
func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	n := 3
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	set[3] = true
	t.Log(len(set))
	// 删除元素
	delete(set, 1)
	n = 1
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
