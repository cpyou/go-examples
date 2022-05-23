package function

func lambda() {
	// 直接执行
	func(s string) {
		println(s)
	}("hello word")

	// 赋给变量
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 2))
}

// 作为参数
func lambda2(f func()) {
	f()
}

// 作为返回值
func lambda3() func(int, int) int {
	return func(i int, i2 int) int {
		return i + i2
	}
}

// 作为结构体字段
func lambdaStruct() {
	type calc struct {
		mul func(x, y int) int
	}
	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}
	println(x.mul(2, 3))
}

// 经通道传递
func lambdaChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}
	println((<-c)(1, 2))
}

// 不曾使用的匿名函数会被编辑器当做错误
//func lambdaNotUsed() {
//	func(s string) {
//		println(s)
//	}
//}

// 除闭包因素外，匿名函数也是一种常见重构手段。可将大函数分解成多个相对独立的匿名函数块，然后用相对简洁的调用完成逻辑流程，以实现框架和细节的分离。
// 相比语句块，匿名函数的作用域被隔离（不使用闭包），不会引发外部污染，更加灵活。没有定义顺序限制，必要时可以抽离，便于实现干净、清晰地代码层次。
