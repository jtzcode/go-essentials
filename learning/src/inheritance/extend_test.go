package inheritance

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Eggy struct {
	Pet
}

// func (e *Eggy) Speak() {
// 	fmt.Println("Miao!")
// }

func TestExtend(t *testing.T) {
	eggy := new(Eggy)
	eggy.Speak()
}
