package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 支持多个返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(fn func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := fn(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// 可变参数
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
}
