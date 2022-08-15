package file

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestFileLock_Lock(t *testing.T) {
	testFilePath, _ := os.Getwd()
	lockedFile := testFilePath + "/test.txt"
	println(lockedFile)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			flock := New(lockedFile)
			err := flock.Lock()
			defer flock.Unlock()
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			flock.f.Write([]byte("asddd\n"))
			fmt.Printf("output : %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}
