package __concurrency

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func SyncChannel() {
	done := make(chan struct{}) // 结束事件
	c := make(chan string)      // 数据传输通道

	go func() {
		s := <-c // 接受消息
		println(s)
		close(done) // 关闭通道，作为结束通知
	}()

	c <- "hi!" // 发送消息
	<-done     // 阻塞，直到有数据或管道关闭
}

// AsyncChannel 同步模式必须有配对操作的goroutine出现，否则会一直阻塞。而异步模式在缓冲区未满或数据未读完前，不会阻塞。
// 多数时候，异步通道有助于提升性能，减少排队阻塞。
func AsyncChannel() {
	c := make(chan int, 3) // 创建带3个缓冲槽的异步通道

	c <- 1 // 缓冲区未满，不会阻塞
	c <- 2

	println(<-c) // 缓冲区尚有数据，不会阻塞
	println(<-c)
}

func Equal() {
	var a, b = make(chan int, 3), make(chan int)

	var c chan bool

	println(a == b)
	println(c == nil)
	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))

}

// JudgeAsyncExam 内置函数cap和len返回缓冲区大小和当前已缓冲数量；而对于同步通道则都返回0，据此可判断通道是同步还是异步。
func JudgeAsyncExam() {
	a, b := make(chan int, 3), make(chan int)
	println("a:", len(a), cap(a))
	a <- 1
	a <- 2
	println("a:", len(a), cap(a))
	println("b:", len(b), cap(b))
}

func JudgeChanSync(c chan int) bool {
	if len(c) == 0 && cap(c) == 0 {
		return true
	}
	return false
}

// ReceiveAndSend 除使用简单的发送和接受符外，还可以使用ok-item或range模式处理数据
func ReceiveAndSend() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) // 确保发送结束通知

		for true {
			x, ok := <-c
			if !ok { // 据此判断通道是否关闭
				return
			}
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

// ReceiveAndSendRange 对于循环接收数据，range模式更简洁一些。
func ReceiveAndSendRange() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) // 确保发送结束通知

		for x := range c { // 循环获取消息，直到通道被关闭
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3

	close(c) // 及时用close函数关闭通道引发结束通知，否则可能会导致死锁

	<-done
}

// MultipleNotify 通知是群体性的。也未必就是结束通知，可以是任何需要表达的事件。
// 注：一次性事件用close效率更好，没有多余开销。连续或多样性事件，可传递不同数据标志实现。还可使用sync.Cond实现单播或广播事件。
func MultipleNotify() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			println(id, ": ready.") // 运动员准备就绪
			<-ready                 // 等待发令
			println(id, ": running")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready? Go!")

	close(ready) // 嘭！

	wg.Wait()
}

func multipleNotify2() {
	c := make(chan int, 3)

	c <- 10
	c <- 20
	close(c)

	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		println(i, ":", ok, x)
	}

}
