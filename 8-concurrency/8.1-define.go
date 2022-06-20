package __concurrency

import (
	"fmt"
	"math"
	"runtime"
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

//  运行时可能会创建很多线程，但任何时候仅有限的几个线程参与并发任务执行。该数量默认与处理器核数相等，可用runtime.GOMAXPROCS函数（或环境变量）修改。
// 如参数小于1，GOMAXPROCS仅返回当前设置值，不做任何调整。

// 测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	println(x)
}

func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

// 并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func GoMaxProc() {
	println("CPU Num", runtime.NumCPU())
	n := runtime.GOMAXPROCS(0)
	println("n:", n)

	start := time.Now()
	println("start", start.String())
	test(n)
	end := time.Now()
	println("end", end.String())

	fmt.Printf("用时：%0.2fS\n", end.Sub(start).Seconds())
}

func GoMaxProc2() {
	println("CPU Num", runtime.NumCPU())
	n := runtime.GOMAXPROCS(4)
	println("n:", n)

	start := time.Now()
	println("start", start.String())
	test2(n)
	end := time.Now()
	println("end", end.String())

	fmt.Printf("用时：%0.2fS\n", end.Sub(start).Seconds())
}

// LocalStorage 与线程不同，goroutine任务无法设置优先级，无法获取编号，没有局部存储(TLS)，
// 甚至连返回值都会被抛弃。但除优先级外，其他功能都很容易实现。
// 注：如使用map作为局部存储容器，建议做同步处理，因为运行时会对其做并发读写检查。
func LocalStorage() {
	var wg sync.WaitGroup
	var gs [5]struct { // 用于实现类似TLS功能
		id     int // 编号
		result int // 返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) { // 使用参数避免闭包延迟求值
			defer wg.Done()
			//time.Sleep(time.Second)

			gs[id].id = id
			gs[id].result = (id + 1) * 100

		}(i)
	}
	fmt.Printf("wg.Wait ... %+v\n", gs)

	wg.Wait()
	fmt.Printf("wg.Wait done %+v\n", gs)
}
