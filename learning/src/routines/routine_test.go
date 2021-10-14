package routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
}

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var lock sync.Mutex
	var wgroup sync.WaitGroup
	counter := 0
	for i := 0; i < 1000; i++ {
		wgroup.Add(1)
		go func() {
			defer func() {
				lock.Unlock()
			}()
			lock.Lock()
			counter++
			wgroup.Done()
		}()
	}
	wgroup.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter)
}
