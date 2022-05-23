package function

import (
	"fmt"
	"log"
	"os"
)

// 延迟调用
// 语句defer向当前函数注册稍后执行的函数调用。
//这些调用直到函数执行结束前才被执行，常用于资源释放、解除锁定，以及错误处理等操作。

func defer1() {
	f, err := os.Open("./4.5-defer.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

}

func defer2() {
	x, y := 1, 2

	defer func(a int) {
		println("defer x,y=", a, y) // y为闭包引用
	}(x) // 注册时复制调用参数

	x += 100 // 对x的修改不会影响延迟函数
	y += 100
	println(x, y)

}

func defer3() {
	// 多个延迟注册按FILO(先进后出)次序执行
	defer println("a")
	defer println("b")
}

func defer4() (z int) {
	defer func() {
		println("defer:", z)
		z += 100 // 修改命名返回值
	}()

	return 100 // 实际执行次序：z=100, call defer, ret
}

// 延迟调用在函数结束时才被执行。不合理的使用方式会浪费更多资源，甚至造成逻辑错误。

// 循环处理多个日志文件，不恰当的defer导致文件关闭时间延长。

func defer5() {
	for i := 0; i < 10000; i++ {
		path := fmt.Sprintf("./log/%d.txt", i)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			continue
		}

		// 这个关闭操作在defer5函数结束时，才会执行，而不是当前循环中执行
		// 这无端延长了逻辑结束时间和f的生命周期，平白多消耗了内存等资源。
		defer f.Close()
	}
}

// 上述文件关闭操作应该改成直接调用，或者重构为函数，将循环和处理算法分离开。
func defer6() {
	do := func(n int) {
		path := fmt.Sprintf("./log/%d.txt", n)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}

		// 该延迟调用在此匿名函数结束时执行，而非defer6
		defer f.Close()
	}
	for i := 0; i < 10000; i++ {
		do(i)
	}
}
