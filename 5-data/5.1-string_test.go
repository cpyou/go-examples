package data

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 转换
// 修改字符串，需将其转换为可变类型（[]rune 或 []byte，）待完成后再转换回来。
// 但不管如何转换，都必须重新分配内存，并复制数据。
func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func TestPP(t *testing.T) {
	s := "hello, world"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs: %x\n", &bs)
	pp("[]byte to string, s2: %x\n", &s2)

	rs := []rune(s)
	s3 := string(rs)

	pp("string to []rune: %x\n", &rs)
	pp("[]rune to string, s3: %x\n", &s3)
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoin()
	}
	// BenchmarkStringJoin-4   	    6386	    164621 ns/op
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin()
	}
	// BenchmarkStringsJoin-4   	  103008	     10734 ns/op
}

func BenchmarkStringJoinBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoinBuffer()
	}
	// BenchmarkStringJoinBuffer-4   	  199977	      5665 ns/op
}
