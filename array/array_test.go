package array

import (
	"fmt"
	"testing"
)

func TestArrayContain(t *testing.T) {
	a := []int{8100, 8200}
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}
	_, ok := m[8200]
	fmt.Println(ok)

}
