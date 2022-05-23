package function

import "testing"

func TestLambda(t *testing.T) {
	lambda()
}

func TestLambda2(t *testing.T) {
	lambda2(func() {
		println("hello world")
	})
}

func TestLambda3(t *testing.T) {
	add := lambda3()
	println(add(1, 2))
}
