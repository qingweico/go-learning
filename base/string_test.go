package base

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// sting 是值类型 默认的初始化值为空字符 而不是nil
func TestString(t *testing.T) {
	var s string
	t.Log(s == "")
	s = "Goland"
	t.Log(len(s))
	// string是不可变的byte slice
	// s[0] = 'g'
	// 可以存储任何二进制数据
	// s = "\xE4\xB8\xA5"
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

// Unicode 和 UTF8
// Unicode 是一种字符集(code point)
// UTF8 是 Unicode的储存实现(转换为字节序列的规则)

func TestStringToRune(t *testing.T) {
	s := "云无心以出岫"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	// int -> string
	s := strconv.Itoa(10)
	t.Log(reflect.TypeOf(s))
	t.Logf("%T", s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(reflect.TypeOf(i))
	}
}
