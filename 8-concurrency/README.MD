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

