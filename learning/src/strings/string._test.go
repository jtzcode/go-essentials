package strings

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringToRune(t *testing.T) {
	s := "中华人们共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFunctions(t *testing.T) {
	s := "X,Y,Z"
	chars := strings.Split(s, ",")
	for _, char := range chars {
		t.Log(char)
	}
	t.Log(strings.Join(chars, "->"))
}

func TestStringConvert(t *testing.T) {
	s := strconv.Itoa(100)
	t.Log("string: " + s)
	if res, err := strconv.Atoi("100"); err == nil {
		t.Log(100 + res)
	}
}
