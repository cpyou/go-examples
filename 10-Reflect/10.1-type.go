package Reflect

import (
	"fmt"
	"net/http"
	"reflect"
)

type X int

func Base() {
	var a X = 100
	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind())
}

type Y int

func Base2() {
	var a, b X = 100, 200
	var c Y = 300

	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)

	fmt.Println("ta", ta.Name(), ta.Kind())
	fmt.Println("tb", tb.Name(), tb.Kind())
	fmt.Println("tc", tc.Name(), tc.Kind())

	fmt.Println(ta == tb, ta == tc)
	fmt.Println(ta.Kind() == tc.Kind())
}

// Base3 传入对象应区分基类型和指针类型，因为他们并不属于同一类型。
// 方法 Elem 返回指针、数组、切片、字典（值）或通道的基类型。
func Base3() {
	x := 100

	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)

	fmt.Println(tx, tp, tx == tp)
	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx == tp.Elem())

	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf(map[string]int32{}).Elem())
}

// Base4 只有在获取结构体指针的基类型后，才能遍历它的字段。
type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func Base4() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem() // 获取指针的基类型
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(i, f.Name, f.Type, f.Offset)
		if f.Anonymous { // 输出匿名字段结构体
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}
}

// 对于匿名字段，可用多级索引（按定义顺序）直接访问。

func Base5() {
	var m manager

	t := reflect.TypeOf(m)

	name, _ := t.FieldByName("name") // 按名称查找
	fmt.Println(name.Name, name.Type)

	name2 := t.FieldByIndex([]int{0, 0}) // 按多级索引查找
	fmt.Println(name2.Name, name2.Type)

	age := t.FieldByIndex([]int{0, 1}) // 按多级索引查找
	fmt.Println(age.Name, age.Type)
}

// 同样地，输出方法时，一样区分基类型和指针类型。

type A int

type B struct {
	A
	name string
}

func (A) av()  {}
func (*A) ap() {}

func (A) Av()  {}
func (*A) Ap() {}

func (B) bv()  {}
func (*B) bp() {}

func (B) Bv()  {}
func (*B) Bp() {}

func Base6() {
	// 小写开头的方法不能获取
	b := B{}
	t1 := reflect.TypeOf(B{})

	fmt.Println(t1, ":", t1.NumMethod())

	t := reflect.TypeOf(&b)

	s := []reflect.Type{t, t.Elem()}

	for _, t := range s {
		fmt.Println(t, ":", t.NumMethod())

		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(" ", t.Method(i))
		}
	}
}

// 反射能探知当前包或外包的非导出结构成员。

func Base7() {
	var s http.Server
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
}

// 可用反射提取 struct tag，还能自动分解。其常用于 ORM 映射，或数据格式验证。

type user2 struct {
	name string `field:"name" type:"varchar(50)"`
	age  int    `field:"age" type:"int"`
}

func Base8() {
	var u user2
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

// 辅助判断方法 Implements、ConvertibleTo、AssignableTo 都是运行期间进行动态调用和赋值所必须的。

type X2 int

func (X2) String() string {
	return ""
}

func Base9() {
	var a X2
	t := reflect.TypeOf(a)

	// Implements 不能直接使用类型作为参数，导致这种用法非常别扭
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))

	it := reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it))

	fmt.Println(t.AssignableTo(st), t.AssignableTo(it))
}
