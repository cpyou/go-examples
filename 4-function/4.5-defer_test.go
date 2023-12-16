package __function

import (
	"sync"
	"testing"
)

func TestDefer1(t *testing.T) {
	defer1()
}

func TestDefer2(t *testing.T) {
	defer2()
}

func TestDefer3(t *testing.T) {
	defer3()
}

func TestDefer4(t *testing.T) {
	println("test:", defer4())
}

func TestDefer5(t *testing.T) {
	defer5()
}

func TestDefer6(t *testing.T) {
	defer6()
}

func TestDefer7(t *testing.T) {
	defer7()
}

// 性能
// 相比直接用CALL汇编指令调用函数，延迟调用则需花费更大的代价。
// 这其中包括注册、调用等操作，还有额外的内存开销。
// 性能要求高且压力大的算法，应避免延迟调用。
var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}

// BenchmarkCall-4   	82310996	        15.28 ns/op
// BenchmarkDefer-4   	60474811	        18.26 ns/op
