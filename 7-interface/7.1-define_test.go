package __interface

import "testing"

func TestTester(t *testing.T) {
	var d data
	// var tt  tester = d // Cannot use 'd' (type data) as the type tester Type does not implement 'tester' as the 'test' method has a pointer receiver

	var tt tester = &d
	tt.test()
	println(tt.string())
}

func TestEmptyInterface(t *testing.T) {
	emptyInterface()
}

// TestEmbed 测试接口嵌入
func TestEmbed(t *testing.T) {
	var d data2
	var t2 tester2 = &d
	t2.test()
	println(t2.string())
}

func TestSubSetExec(t *testing.T) {
	SubSetExec()
}

func TestLambdaInterface(t *testing.T) {
	LambdaInterface()
}
