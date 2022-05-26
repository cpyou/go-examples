package data

import (
	"bytes"
	"strings"
)

// 用加法操作拼接字符串时，每次都会重新分配内存
func stringJoin() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

// 使用内置strings.Join 拼接字符串，提升性能
func stringsJoin() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

// 使用bytes.Buffer提升拼接字符串性能
func stringJoinBuffer() string {
	var b bytes.Buffer
	b.Grow(1000) // 事先准备足够的内存，避免中途扩张

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}

// 少量字符串拼接可使用fmt.Sprintf、text/template等方法

func Join(a []string, seq string) string {
	n := len(seq) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	// 一次分配所需内存
	b := make([]byte, n)

	// 拷贝数据
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], seq)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
