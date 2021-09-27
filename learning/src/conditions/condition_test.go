package conditions

import "testing"

func TestCondition(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("Out of 0-3")
		}
	}
}
