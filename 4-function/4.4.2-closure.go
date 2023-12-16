package __function

// 闭包（closure）是在词法上下文中引用了自由变量的函数，或者说是函数和其作用的环境的组合体。

func closure(x int) func() {
	return func() {
		println(x)
	}
}

func closure2(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}

func closure3() []func() {
	var s []func()

	// for循环复用局部变量i，每次添加的匿名函数引用的是同一个变量。添加操作仅仅是将匿名函数放入列表，并未执行。
	// 在外部执行这些函数时，它们读取的是环境变量i最后一次循环时的值。结果是2。
	for i := 0; i < 2; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}
	return s
}

func closure4() []func() {
	var s []func()

	// 解决上述问题的方法是每次用不同的环境变量或传参复制，让各自闭包环境各不相同。
	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}

func closure5(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10 // 修改环境变量
		}, func() {
			println(x) // 显示环境变量
		}
}
