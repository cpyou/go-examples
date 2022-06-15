package __concurrency

import "testing"

func TestDeferExec(t *testing.T) {
	DeferExec()
}

func TestWait(t *testing.T) {
	Wait()
}

func TestWaitGroup(t *testing.T) {
	WaitGroup()
}

func TestIncorrectWait(t *testing.T) {
	IncorrectWait()
}
func TestMultiWait(t *testing.T) {
	MultiWait()
}
