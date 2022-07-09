package Reflect

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Value() {
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
}

type User struct {
	Name string
	code int
}

func Value2() {
	p := new(User)
	v := reflect.ValueOf(p).Elem()

	name := v.FieldByName("Name")
	code := v.FieldByName("code")

	fmt.Printf("name: canaddr = %v, canset = %v\n", name.CanAddr(), name.CanSet())
	fmt.Printf("code: canaddr = %v, canset = %v\n", code.CanAddr(), code.CanSet())

	if name.CanSet() {
		name.SetString("Tom")
	}

	if code.CanAddr() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}

	fmt.Printf("%+v\n", *p)
}

// Value3 可通过 Interface 方法进行类型推断和转换。
// 也可以直接使用 Value.int、 Bool 等方法进行类型转换，但失败时会引发 panic，且不支持 ok-idiom。
func Value3() {
	type user struct {
		Name string
		Age  int
	}

	u := user{
		"q.yuhen",
		60,
	}

	v := reflect.ValueOf(&u)

	if !v.CanInterface() {
		println("CanInterface: fail.")
		return
	}

	p, ok := v.Interface().(*user)
	if !ok {
		println("Interface: fail.")
		return
	}

	p.Age++
	fmt.Printf("%+v\n", u)

}

func Value4() {
	// 复合类型对象设置示例
	c := make(chan int, 4)
	v := reflect.ValueOf(c)

	if v.TrySend(reflect.ValueOf(100)) {
		fmt.Println(v.TryRecv())
	}
	fmt.Println()

	// 接口有两种 nil 状态，折一直是个潜在麻烦。解决方法是用 IsNil 判断值是否为 nil。
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	fmt.Println(a == nil)
	fmt.Println(b == nil, reflect.ValueOf(b).IsNil())
	fmt.Println()

	// 也可用 unsafe 转换后直接判断 iface.data 是否为零值。
	iface := (*[2]uintptr)(unsafe.Pointer(&b))

	fmt.Println(iface, iface[1] == 0)
	fmt.Println()

	// Value 里的某些方法并未实现 ok-idiom 或返回 error，所以得自行判断返回的是否为 Zero Value。
	v1 := reflect.ValueOf(struct {
		name string
	}{})

	println(v1.FieldByName("name").IsValid())
	println(v1.FieldByName("xxx").IsValid())
	fmt.Println()
}
