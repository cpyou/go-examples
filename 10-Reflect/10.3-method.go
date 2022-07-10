package Reflect

import (
	"fmt"
	"reflect"
)

type X1 struct{}

func (X1) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err: %d", x+y)
}

func Method() {
	var a X1
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")
	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}

	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}

}

// 对于变参来说，用 CallSlice 要更方便一些。

func (X1) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

func Method2() {
	var a X1

	v := reflect.ValueOf(&a)
	m := v.MethodByName("Format")

	out := m.Call([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf("x"), // 所有参数都需处理
		reflect.ValueOf(100),
	})

	fmt.Println(out)

	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf([]interface{}{"x", 100}), // 仅一个 []interface{} 即可
	})
	fmt.Println(out)
}
