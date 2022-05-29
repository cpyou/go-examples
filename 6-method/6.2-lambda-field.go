package __method

import "sync"

type data struct {
	sync.Mutex
	buf [1024]byte
}

// 匿名字段可以像访问匿名字段成员那样调用其方法，由编译器负责查找
func lambdaField() {
	d := data{}
	d.Lock() // 编译器会处理为 sync.(*Mutex).Lock()调用
	defer d.Unlock()
}

// 方法也会有同名遮蔽问题。利用这个特性，可以实现类似覆盖（overload）操作。
type user struct{}

type manager struct {
	user
}

func (u user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + "; manager"
}

// 尽管能直接访问匿名字段的成员和方法，但它们依然不属于继承关系。
func lambdaField2() {
	var m manager
	println(m.toString())
	println(m.user.toString())
}
