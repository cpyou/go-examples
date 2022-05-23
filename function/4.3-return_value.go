package function

import (
	"errors"
)

// 4.3 返回值
func div(x, y int)(z int, err error){
	if y == 0 {
		err = errors.New("division by zero")
	}
	z = x / y
	return
	//return z, err
}

func test()(int, s string, e error){
	//return 0, "", nil  // '0' (type untyped int) cannot be represented by the type string
	return "", "", nil
}

type User struct {
}

// 如果返回值类型能明确表明其含义，就尽量不要对其命名
func newUser()(*User, error)  {
	return &User{}, nil
}

