package base

import (
	"fmt"
	"runtime"
	"testing"
)

func TestLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	var j = 0
	for j < 5 {
		fmt.Print(j)
		j++
	}
	fmt.Println()
}

func TestIfCondition(t *testing.T) {
	if r := f(); r == 0 {
		t.Log("")
	} else {
		t.Log("")
	}
}
func f() int {
	return 0
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	// 每个case后面不需要指定break语句
	// 若case中有多个结果集可以使用逗号分隔
	case "darwin,xxx":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
}
func TestMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2, 4:
			fmt.Println("Even")
		case 1, 3:
			fmt.Println("Odd")
		}
	}
}
