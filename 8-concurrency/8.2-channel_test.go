package __concurrency

import (
	"testing"
)

func TestSyncChannel(t *testing.T) {
	SyncChannel()
}

func TestAsyncChannel(t *testing.T) {
	AsyncChannel()
}

func TestEqual(t *testing.T) {
	Equal()
}

func TestJudgeAsync(t *testing.T) {
	JudgeAsyncExam()
}

func TestJudgeChanSync(t *testing.T) {
	type args struct {
		c chan int
	}
	c1 := make(chan int)
	c2 := make(chan int, 1)
	c3 := make(chan int, 2)
	c4 := make(chan int, 3)
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{name: `sync`, args: args{c: c1}, want: true},
		{name: `async`, args: args{c: c2}, want: false},
		{name: `async`, args: args{c: c3}, want: false},
		{name: `async`, args: args{c: c4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt)
			if got := JudgeChanSync(tt.args.c); got != tt.want {
				t.Errorf("JudgeChanSync() = %v, want %v", got, tt.want)
			}
		})
	}
}
