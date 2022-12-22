package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	fmt.Println(os.Args)
	fmt.Println("hello, world")
	os.Exit(0)
}
