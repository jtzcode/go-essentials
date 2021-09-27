package operators

import "testing"

const (
	Read = 1 << iota
	Write
	Exec
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	c := [...]int{1, 3, 2, 4}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)

}

func TestBitClear(t *testing.T) {
	a := 7        //0111
	a = a &^ Read // this will set all bits of Read to 0
	t.Log(a&Read == Read, a&Write == Write, a&Exec == Exec)
}
