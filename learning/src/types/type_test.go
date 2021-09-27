package types

import "testing"

type MyInt int64

func TestImplicitConvert(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt = MyInt(b)
	t.Log(a, b, c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 Error: do no support pointer calculation
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("^" + s + "^")
	t.Log(len(s))
}
