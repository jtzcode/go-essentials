package oo

import (
	"fmt"
	"testing"
)

type Eggy struct {
	ID   string
	Name string
	Age  int
}

type Coder interface {
	WriteGracefulCode() string
}

type GoCoder struct {
}

func (gc *GoCoder) WriteGracefulCode() string {
	return "This code is graceful!"
}

func (e *Eggy) toString() string {
	return fmt.Sprintf("ID: %s / Name: %s / Age: %d", e.ID, e.Name, e.Age)
}

func TestInterface(t *testing.T) {
	var coder Coder
	//Duck type
	coder = new(GoCoder)
	t.Log(coder.WriteGracefulCode())
}
func TestEggy(t *testing.T) {
	e := Eggy{"1", "dandan", 1}
	t.Log(e.toString())
}
