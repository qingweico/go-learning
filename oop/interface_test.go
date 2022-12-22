package oop

import (
	"fmt"
	"testing"
)

// Go 接口
// 接口为非侵入性 实现不依赖于接口的定义 所以接口的定义可以包含在接口使用者包内
// ------------------------- Colour -------------------------
type Colour interface {
	Color() string
}

// ------------------------- Shape -------------------------
type Shape struct {
}

func (s *Shape) ToString() string {
	return "Shape"
}

// ------------------------- Circle -------------------------
type Circle struct {
	p *Shape
}

func (c *Circle) ToString() string {
	return c.p.ToString()
}

func (c *Circle) Color() string {
	return "Red Circle"
}

// ------------------------- Square -------------------------
type Square struct {
}

func (s *Square) Color() string {
	return "Color Square"
}

// ------------------------- Ellipse -------------------------
type Ellipse struct {
}

func (e *Ellipse) Color() string {
	return "Color Ellipse"
}

func printShape(c Colour) {
	fmt.Printf("%T, %v\n", c, c.Color())
}
func TestPolymorphism(t *testing.T) {
	square := new(Square)
	ellipse := new(Ellipse)
	printShape(square)
	printShape(ellipse)
}

func TestInterface(t *testing.T) {
	var c Colour
	c = new(Circle)
	fmt.Println(c.Color())
	var id i64 = 10
	fmt.Println(id)
}

// 自定义数据类型
type i64 int64

// 复用而不是继承
