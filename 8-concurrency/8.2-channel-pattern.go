package __concurrency

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// 通常使用工厂方法将goroutine和通道绑定。
type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)

	go func() {
		defer r.Done()
		for x := range r.data { // 接收消息，直到通道被关闭
			println("recv:", x)
		}
	}()
	return r
}

func Receive() {
	r := newReceiver()

	r.data <- 1
	r.data <- 2

	close(r.data) // 关闭通道，发出结束通知
	r.Wait()      // 等待接收者处理结束
}

// 鉴于通道本身就是一个并发安全的队列，可用ID generator、Pool等用途。
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p: //返回
	default: // 返回失败，新建
		v = make([]byte, 10)
	}

	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b: // 放回
	default: // 放回失败，放弃
	}
}

func Queue() {
	q := newPool(3)
	for i := 0; i < 4; i++ {
		s := fmt.Sprintf("%d", i)
		q.put([]byte(s))
	}
	for i := 0; i < 5; i++ {
		v := q.get()
		fmt.Printf("%d, v: %s\n", i, v)
	}
}

// Semaphore 用通道实现信号量（semaphore）
func Semaphore() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许2个并发同时执行

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // acquire: 获取信号
			defer func() { <-sem }() // release: 释放信号

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}

// TimeTick 标准库 time 提供了 timeout 和 tick channel 实现。
func TimeTick() {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("timeout ...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)

		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil) // 直接用 nil channel 阻塞进程
}
