package function

import (
	"fmt"
	"log"
	"time"
)

// 默认参数
type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{ // 默认参数
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/variableParams",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func DefaultOptionFunc(opt *serverOption) {}

// 4.2 可变参数
func variableParams(s string, a ...int){
	fmt.Println(s, a)
}
