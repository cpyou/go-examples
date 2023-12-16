package __function

import "testing"

func TestClosure(t *testing.T) {
	// closure函数返回的匿名函数会引用上下文环境变量x。当该函数在其他函数中执行时，它仍然可正确读取x的值，这种现象叫做闭包。
	f := closure(123)
	f()
}

func TestClosure2(t *testing.T) {
	f := closure2(0x100)
	//f := closure2(123)
	f()
}

func TestClosure3(t *testing.T) {

	for _, f := range closure3() {
		f()
	}
}

func TestClosure4(t *testing.T) {
	for _, f := range closure4() {
		f()
	}
}

func TestClosure5(t *testing.T) {
	// 返回值是100， 110
	// 多个匿名函数引用同一环境变量，会让事情变得更加复杂。任何的修改行为都会影响其他函数的取值，在并发模式下可能需要做同步处理。
	a, b := closure5(100)
	b()
	a()
	b()
	a()
	b()
}

// 闭包让我们不用传递参数就可读取或修改环境状态，但也要为此付出额外代价，对于性能要求较高的场合，需慎重使用。
