package error_examples

import "fmt"

func ErrorTest() (err error) {
	err = fmt.Errorf("asd")
	return
}
