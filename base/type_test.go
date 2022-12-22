package base

import (
	"testing"
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func Test(t *testing.T) {
	t.Log(Readable, Writable, Executable)
}

func TestTypeConvert(t *testing.T) {
	var a int32 = 1
	var b int64 = 2
	// 不会发生隐式数据类型转换
	var c int64 = int64(a) + b
	t.Log(c)
}
func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	// 不支持指针运算操作
	// aPtr++
	t.Logf("%T, %T", a, aPtr)
}
