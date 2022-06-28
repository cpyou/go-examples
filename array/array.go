package array

import "fmt"

func Contain() {
	a := []int{8100, 8200}
	m := make(map[int]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	_, ok := m[8200]
	fmt.Println(ok)
}
