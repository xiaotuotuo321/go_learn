package main

// go标准库的Context

// 在go http包的server中，每一个请求都有一个对应的goroutine去处理。请求处理函数通常会启动额外的goroutine用来访问后端服务，如果数据库和RPC服务.
// 用来处理一个请求的goroutine通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。当一个请求被取消或超时时。
// 所有用来处理该请求的goroutine都应该迅速推出，然后系统才能释放这些goroutine占用的资源。

// 1.为什么需要context
// 1.1.基本示例
//var wg sync.WaitGroup
//
//func worker(){
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//	}
// 	// 如何接受外部命令实现退出
// 	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	// 实现结束子程序的goroutine
//	wg.Wait()
//	fmt.Println("over")
//}

// 1.2.全局变量的方式
//var wg sync.WaitGroup
//var exit bool

/*
使用全局变量的问题：
	1. 使用全局变量在跨包调用时不容易统一
	2. 在worker中再启动goroutine，就不容易进行控制
*/

//func worker(){
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		if exit{
//			break
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	time.Sleep(time.Second * 3)
//	exit = true
//	wg.Wait()
//	fmt.Println("over")
//}

// 1.3.通过通道的方式
//var wg sync.WaitGroup
//// 管道方式存在的问题：
//// 1.使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个公用的channel
//func worker(exitChan chan struct{}){
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <- exitChan:	// 等待接收上级通知
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	var exitChan = make(chan struct{})
//	wg.Add(1)
//	go worker(exitChan)
//	time.Sleep(time.Second * 3)		// sleep3秒以免程序过快退出
//	exitChan <- struct{}{}			// 给子goroutine发送退出信号
//	close(exitChan)
//	wg.Wait()
//	fmt.Println("over")
//}

// 1.4.官方的答案
//var wg sync.WaitGroup
//
//func worker(ctx context.Context){
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select{
//		case <- ctx.Done():
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 3)
//	cancel()	// 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}

// 1.4.1.当goroutine又启动了另一个goroutine时，只需要将ctx传入即可
//var wg sync.WaitGroup
//
//func worker(ctx context.Context){
//	go worker1(ctx)
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <- ctx.Done():
//			break LOOP
//			default:
//		}
//	}
//	wg.Done()
//}
//
//func worker1(ctx context.Context){
//LOOP1:
//	for {
//		fmt.Println("worker1")
//		time.Sleep(time.Second)
//		select {
//		case <- ctx.Done():
//			break LOOP1
//		default:
//		}
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 3)
//	cancel()	// 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}

// 2.context
/*
GO1.7加入了一个新的标准库context，定义了context类型，专门用来简化，对于处理单个请求的多个goroutine之间的请求域的数据、取消信号、介质时间等相关
操作，这些操作可能涉及多个API的调用。
对于服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。他们之间的函数调用链必须传递上下文，或者可以使用withCancel、withDeadline、
withTimeout、withValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消
*/

