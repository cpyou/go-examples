# 8.1 并发的含义

只需在函数调用前添加go关键字即可创建并发任务。

关键字go并非执行并发操作，而是创建一个并发任务单元。

与defer一样，goroutine也会“延迟执行”而立即计算并复制执行参数。

Wait

GOMAXPROCS

Local Storage

Gosched

暂停，释放线程去执行其他任务。当前任务被放回队列，等待下次调度时恢复执行。

Goexit

Goexit 立即终止当前任务，运行时确保所有已注册延迟调用被执行。该函数不会影响其他并发任务，不会引发panic，自然无法捕捉。

# 8.2 通道

对于 closed 或 nil 通道，发送或接收操作都有相应规则：

- 向已关闭通道发送数据，引发panic。
- 从已关闭通道接收数据，返回已缓冲数据或零值。
- 无论收发，nil通道都会阻塞。

重复关闭，或者关闭nil通道都会引发panic错误
```shell
panic: close of closed chanel
panic: close of nil channel
```

## 单向

通道默认是双向的，并不区分发送和接受端。但某些时候，我们可限制收发操作的方向来获得更严谨的操作逻辑。

通常使用类型转换来获取单向通道，并分别赋予操作双方。

不能在单向通道上做逆向操作。

同样，close不能用于接受端。

无法将单向通道重新转换回去。

## 选择

如要同时使用多个通道，可选用select语句。它会随机选择一个可用通道做收发操作。

## 模式

通常使用工厂方法将goroutine和通道绑定。

鉴于通道本身就是一个并发安全的队列，可用ID generator、Pool等用途。

用通道实现信号量（semaphore）

标准库 time 提供了 timeout 和 tick channel 实现。

捕捉 INT、TERM 信号，顺便实现一个简易的atexit函数。示例：8.2-channel-at-exit

## 性能

将发往通道的数据打包，减少传输次数，可有效提升性能。从实现来说，通道队列依旧使用锁同步机制，单次获取更多数据（批处理），
可改善因频繁加锁造成的性能问题。

## 资源泄露

通道可能引发 goroutine leak，确切地说，是指 goroutine 处于发送或接收阻塞状态，但一直未被唤醒。
垃圾回收器并不收集此类资源，导致他们会在等待队列里长久休眠，形成资源泄露。

示例：8.2-channel-leak

8.3 同步

// 应将 Mutex 锁粒度控制在最小范围内，及早释放。

```go
// 错误用法
func doSomething()  {
    m.Lock()
    url := cache["key"]
    http.Get(url) // 该操作并不需要锁保护
    m.unLock()
}

// 正确做法
func doSomething()  {
    m.Lock()
    url := cache["key"]
    m.unLock() // 如使用 defer，则依旧将 Get 保护在内
    http.Get(url) // 该操作并不需要锁保护
}
```

Mutex 不支持递归锁，即便在同一 goroutine 下也会导致死锁。

相关建议：
- 对性能要求较高时，应避免使用defer Unlock。
- 读写并发时，应用 RWMutex 性能会更好一些。
- 对单个数据读写保护，可尝试原子操作。
- 执行严格测试，尽可能打开数据竞争检查。