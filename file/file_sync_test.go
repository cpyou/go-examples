package file

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			s, err := Read()
			fmt.Println(s)
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			count, err := strconv.Atoi(s)
			count += 1
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}

			err = ReadWrite()
			if err != nil {
				return
			}
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("output : %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}
