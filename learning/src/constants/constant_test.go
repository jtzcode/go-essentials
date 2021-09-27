package constants

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}
