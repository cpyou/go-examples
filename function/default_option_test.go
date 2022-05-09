package function

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	opt := newOption()
	opt.port = 8085
	fmt.Println(opt)
	DefaultOptionFunc(opt)
}
