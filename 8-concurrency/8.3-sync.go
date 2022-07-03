package __concurrency

import (
	"sync"
	"time"
)

// 标准库 sync 提供了互斥和读写锁，另有原子操作等。可基本满足日常开发需要。
// Mutex 、 RWMutex 的使用并不复杂，只有几个地方需要注意。
// 将 Mutex 作为匿名字段时，相关方法必须实现为 pointer-receiver，否则会因复制导致锁失效。
type data struct {
	sync.Mutex
}

// test 此示例，锁失效，将 receiver 类型改为 *data 后正常。也可以嵌入 *Mutex 来避免复制问题，但那需要专门初始化。
//func (d *data) test(s string) {
func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

// Sync ...
func Sync() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()

}

// DeadLock Mutex 不支持递归锁，即便在同一 goroutine 下也会导致死锁。
func DeadLock() {
	var m sync.Mutex

	m.Lock()
	{
		m.Lock()
		m.Unlock()
	}
	m.Unlock()
}

type cache struct {
	sync.Mutex
	data []int
}

func (c *cache) count() int {
	c.Lock()
	n := len(c.data)
	c.Unlock()

	return n
}

func (c *cache) get() int {
	c.Lock()
	defer c.Unlock()

	var d int
	if n := c.count(); n > 0 { // count 重复锁定，导致死锁
		d = c.data[0]
		c.data = c.data[1:]
	}
	return d
}

func CacheMain() {
	c := cache{
		data: []int{1, 2, 3, 4},
	}
	println(c.get())
	println(c.get())
	println(c.get())
}
