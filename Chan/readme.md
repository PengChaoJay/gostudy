## go chan 的简单使用

#### 基本使用方式

```go
court := make(chan int) // 不带缓冲的
tasks := make(chan string, taskLoad) //带缓冲的
      := make(<-chan string)  //只读的通道
      := make(chan<- string)  //只写的通道
```

```go

	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
    task, ok := <-tasks
```
