package function

import (
	"fmt"
	"testing"
)

// 默认参数
// 如果参数过多，建议将其重构为一个复合结构类型，也算是变相实现可选参数和命名实参功能
func TestDefaultOptionFunc(t *testing.T) {
	opt := newOption()
	opt.port = 8085
	fmt.Println(opt)
	DefaultOptionFunc(opt)
}

func TestVariableParams(t *testing.T) {
	variableParams("a", 3, 5, 7)

	// 切片作为变参时，需展开操作
	a := [3]int{10, 20, 30}
	variableParams("slice", a[:]...)
}
