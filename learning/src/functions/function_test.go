package functions

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func multiReturnFunction() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpentWrapper(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("Time elapsed: ", time.Since(start).Seconds())
		return ret
	}
}

func myFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func SumInt(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestMultiReturn(t *testing.T) {
	a, _ := multiReturnFunction()
	t.Log(a)

	// Myfunc with function calculating time spent
	timeSpentWrapper(myFunc)(10)

	t.Log(SumInt(1, 2, 3))
	t.Log(SumInt(1, 2, 3, 4))
}

func ClearResource() {
	fmt.Println("Resource cleared")
}

func TestDeferFunc(t *testing.T) {
	defer ClearResource()
	t.Log("Start")
	panic("Error!")
	//t.Log("End")
}
