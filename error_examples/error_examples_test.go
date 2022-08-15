package error_examples

import (
	"fmt"
	"testing"
)

func TestErrorTest(t *testing.T) {
	err := ErrorTest()
	println(err.Error())
	fmt.Printf("err:%s\n", err)
	fmt.Printf("err:%v\n", err)
	fmt.Printf("err:%#v\n", err)
	fmt.Printf("err:%s\n", err.Error())
}
