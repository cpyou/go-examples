package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTimesCounter(t *testing.T) {
	counter := NewTimesCounter(60 * time.Second)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(from int) {
			defer wg.Done()
			counter.Incr("a", 1)
			fmt.Println("from:", from)
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.get("a"))
	counter.remove("a")
	fmt.Println(counter.get("a"))
	counter.decr("a", 1)
	counter.decr("a", 1)
	fmt.Println(counter.get("a"))

}
