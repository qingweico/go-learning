package base

import "testing"

func TestMapInit(t *testing.T) {
	// 初始化 map 的几种方式
	m1 := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}

	t.Log(m1["k1"])
	m1["k4"] = "v4"

	// 对于 map 无法使用cap
	t.Log(len(m1), m1)

	m2 := map[int]int{}
	t.Log(len(m2), m2)

	m3 := make(map[string]string, 5)
	t.Log(len(m3), m3)
}

func TestMapTravel(t *testing.T) {
	m := map[string]string{"k1": "v1", "k2": "v2"}
	for k, v := range m {
		t.Log(k, v)
	}
}
func TestAccessNotExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	m[3] = 0
	if v, present := m[3]; present {
		t.Log("Existing", v)
	} else {
		t.Log("Not existing.")
	}
}
