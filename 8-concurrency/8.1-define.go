package __concurrency

import (
	"sync"
	"time"
)

// 与defer一样，goroutine也会“延迟执行”而立即计算并复制执行参数。
var c int

func counter() int {
	c++
	return c
}

func DeferExec() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second) // 让goroutine在DeferExec函数逻辑之后执行
		println("go:", x, y)
	}(a, counter())

	a += 100
	println("main:", a, counter())
	time.Sleep(time.Second * 3) // 等待goroutine结束
	// main: 200 2
	// go: 100 1
}

func Wait() {
	exit := make(chan struct{}) // 创建通道。因为尽是通知，数据并没有实际意义

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit) // 关闭通道，发送信号
	}()

	println("main ...")
	<-exit // 如通道关闭，立即解除阻塞
	println("main exit.")
	// main ...
	// goroutine done.
	// main exit.
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 累加计数

		go func(id int) {
			defer wg.Done() // 递减计数
			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}

	println("main ...")
	wg.Wait() // 阻塞，直到计数归零
	println("main exit.")
	//main ...
	//goroutine 3 done.
	//goroutine 4 done.
	//goroutine 1 done.
	//goroutine 6 done.
	//goroutine 9 done.
	//goroutine 2 done.
	//goroutine 8 done.
	//goroutine 7 done.
	//goroutine 0 done.
	//goroutine 5 done.
	//main exit.
}

// IncorrectWait 尽管WaitGroup.Add实现了原子操作，但建议在goroutine外累加计数器，以免Add尚未执行，Wait已经退出。
func IncorrectWait() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1) // 来不及设置
		println("hi!")
	}()
	wg.Wait()
	println("exit.")
}

// MultiWait 可在多处使用Wait阻塞，他们都能接收到通知。
func MultiWait() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait() // 等待归零，解除阻塞
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done() // 递减计数
	}()

	wg.Wait() // 等待归零，解除阻塞
	println("main exit.")
	//done.
	//wait exit.
	//main exit.
}
