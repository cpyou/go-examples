package __concurrency

import (
	"fmt"
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
