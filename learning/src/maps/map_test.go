package maps

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	t.Logf("Len of m1=%d", len(m1))
	m2 := map[int]int{}
	m2[5] = 25
	t.Logf("Len of m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("Len of m3=%d", len(m3))

	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestCheckExists(t *testing.T) {
	m1 := map[int]int{}
	m1[2] = 0
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Value of Key 3 is %d", v)
	} else {
		t.Log("Key 3 does not exist")
	}
}

func TestMapFactory(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}
