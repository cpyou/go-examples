package __function

// errors.New ; fmt.Errorf
import (
	"fmt"
	"log"
)

type DivError struct {
	x, y int
}

func (DivError) Error() string {
	return "division by zero"
}

func div1(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}
	return x / y, nil
}

func error1() {
	z, err := div1(5, 0)

	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
		log.Println(err)
	}
	println(z)
}

// 大量函数和方法返回error，使得代码变得很难看，一堆堆的检查语句充斥在代码行间。解决思路：
// * 使用专门的检查函数处理逻辑错误（比如记录日志），简化代码检查
// * 在不影响逻辑的情况下，使用defer延后处理错误状态(err退化赋值)
// * 在不中断逻辑的情况下，将错误作为内部状态保存，等最终"提交"时，再处理。

// panic, recover;
// 连续调用panic，仅最后一个会被recover
// 除非时不可恢复性、导致系统无法正常工作的错误，否则不建议使用panic
