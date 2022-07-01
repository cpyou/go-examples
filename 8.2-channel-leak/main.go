package main

import (
	"runtime"
	"time"
)

func test() {
	c := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			<-c
		}()
	}
}

func main() {
	test()

	for {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

// go build -o test
// GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1" ./test
