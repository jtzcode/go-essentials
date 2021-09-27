package fib

import "testing"

func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 2
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()

}
