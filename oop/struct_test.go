package oop

import "testing"

type User struct {
	Id   int64
	Name string
}

func TestCreateStruct(t *testing.T) {
	u1 := User{1, "Milk"}
	u2 := User{Id: 2, Name: "Maria"}
	// 返回地址
	u3 := new(User)
	u3.Id = 3
	u3.Name = "Stoke"
	t.Log(u1)
	t.Log(u2)
	t.Log(u3)
}
