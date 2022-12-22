package base

import (
	"fmt"
	"testing"
)

// 使用 == 比较两个数组
// 只有当被比较的数组维数和含有的元素个数都相同时,才可以比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	c := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	// error
	// t.Log(a == c)
	fmt.Println(c)

}
func TestBitClear(t *testing.T) {
	a := 7
	// &^ 按位置零操作符: 右操若为1 则左侧变为0 若是0则不变(二进制)

	// 7
	t.Log(a &^ 0)
	// 6
	t.Log(a &^ 1)
}
