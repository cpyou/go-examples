package __interface

// 接口通常以er作为名称后缀，方法名是声明组成部分，但参数名不同或省略。
type tester interface {
	test()
	string() string
}

type data struct{}

func (d *data) test() {}

func (d data) string() string { return "" }

// 如果接口没有任何方法声明，那么就是一个空接口（interface{}），它的用途类似面向对象里的根类型Object，可被赋值为任何类型。
func emptyInterface() {
	var t1, t2 interface{}
	println(t1 == nil, t2 == nil)
	t1, t2 = 100, 10
	println(t1 == t2)
	t1, t2 = map[string]int{}, map[string]int{}
	println(t1 == t2)
}

// 嵌入其他接口类型，相当于将其声明的方法导入。这就要求不能够有同名的方法，因为不支持重载。还有，不能嵌入自身或循环嵌入，会导致递归错误。
type stringer interface {
	string() string
}

type tester2 interface {
	stringer // 嵌入其他接口
	test()
}

type data2 struct{}

func (*data2) test() {}

func (data2) string() string {
	return "string val"
}

// 超集接口变量可隐式转换为子集，反过来不行。
func pp(a stringer) {
	println(a.string())
}

func SubSetExec() {
	var d data2
	var t tester2 = &d
	pp(t) // 隐式转换为子集接口

	var s stringer = t
	println(s.string())

	// var t2 tester2 = s // Cannot use 's' (type stringer) as the type tester2 Type does not implement 'tester2' as some methods are missing: test()
}

// 支持接口类型，可直接用于变量定义，或作为结构字段类型。

type data3 struct{}

func (data3) string() string {
	return "string val"
}

type node struct {
	data interface { // 匿名接口类型
		string() string
	}
}

func LambdaInterface() {
	var t interface { // 定义匿名接口变量
		string() string
	} = data3{}

	n := node{
		data: t,
	}
	println(n.data.string())
}
