package oop

import (
	"fmt"
	"testing"
)

// 空接口和断言

// 空接口可以表示任何类型
// 通过断言将空接口转换为指定类型
func TypeOf(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Other Type")
	}
}
func isInteger(p interface{}) bool {
	if _, ok := p.(int); ok {
		return true
	} else {
		return false
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	fmt.Println(isInteger(2.5))
	TypeOf(10)
	TypeOf("10")
}
