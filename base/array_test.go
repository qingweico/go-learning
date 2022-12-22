package base

import "testing"

func TestArrayInit(t *testing.T) {
	// 数组初始化的几种方式

	// 声明但不赋值, 则默认零值
	var a1 [3]int
	t.Log(a1)

	a2 := [4]int{1, 2, 3, 4}
	t.Log(a2)

	a3 := [...]int{2, 4, 0, 9, 6}
	t.Log(a3)
}

func TestArrayTravel(t *testing.T) {
	a := [4]int{1, 2, 3, 4}

	// 带有索引
	for idx, item := range a {
		t.Log(idx, item)
	}
	// 使用 _ 占位参数
	for _, item := range a {
		t.Log(item)
	}
}

func TestArraySlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(a[2:5])
	t.Log(a[:4])
	t.Log(a[2:])
	t.Log(a[:])
}
