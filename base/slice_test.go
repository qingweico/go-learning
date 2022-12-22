package base

import "testing"

func TestSliceInit(t *testing.T) {
	// 切片声明的几种方式

	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// len是3 容量是5
	s2 := make([]int, 3, 5)

	t.Log(s2)

	// error 访问越界
	// t.Log(s2[3])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "六月"
	t.Log(Q2)
	t.Log(year)
}
func TestSliceComparing(t *testing.T) {
	// 切片只能和nil比较 切片之间不可以比较
}
