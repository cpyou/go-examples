package cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewCache(t *testing.T) {

	counter := NewCache(5, 5)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(from int) {
			defer wg.Done()
			counter.Set("a", from)
			fmt.Println("from:", from)
			fmt.Println(counter.Get("a"))
		}(i)
	}
	wg.Wait()
}
