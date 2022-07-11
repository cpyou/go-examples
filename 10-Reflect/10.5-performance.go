package Reflect

import "reflect"

type Data struct {
	X int
}

var d = new(Data)

func set(x int) {
	d.X = x
	//BenchmarkSet-12    	1000000000	         0.3059 ns/op
}

func rset(x int) {
	v := reflect.ValueOf(d).Elem()
	f := v.FieldByName("X")
	f.Set(reflect.ValueOf(x))
	//BenchmarkRSet-12    	13727745	        83.08 ns/op
}

// 以下通过将反射数据“缓存”起来，提高性能。但是差距依然很大。
var v1 = reflect.ValueOf(d).Elem()
var f = v1.FieldByName("X")

func rset2(x int) {
	f.Set(reflect.ValueOf(x))
	//BenchmarkRSet2-12    	87114969	        13.86 ns/op
}

// 以下对比方法直接调用，和反射调用的性能。

func (x *Data) Inc() {
	x.X++
}

var d1 = new(Data)
var v2 = reflect.ValueOf(d1)
var m = v2.MethodByName("Inc")

func call() {
	d1.Inc()
	//BenchmarkCall-12    	968624632	         1.237 ns/op
}

func rcall() {
	m.Call(nil)
	//BenchmarkRCall-12    	 7598384	       158.8 ns/op
}
