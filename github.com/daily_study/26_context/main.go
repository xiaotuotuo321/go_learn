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

// 3.context接口： context.Context是一个接口，该接口定义了四个需要实现的方法。
//type Context interface{
//	Deadline() (deadline time.Time, ok bool)
//	Done() <- chan struct{}
//	Err() error
//	Value(key interface{}) interface{}
//}

/*
Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间
Done方法需要返回一个channel，这个channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个channel
Err 方法会返回当前context结束的原因，他只会在done返回的channel被关闭时才会返回非空的值
	如果当前context被取消就会返回canceled错误
	如果当前context超时就会返回deadlineExceeded错误
value方法会从context中返回键对应的值，对于同一个上下文来说，多次调用value并传入相同的key会返回相同的结果，该方法仅用于传递跨API和进程间请求域的数据；
*/

// 4.with系列函数
// 4.1.withCancel的函数标签
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// withCancel 返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况
// 取消此上下文将释放与其关联的资源，一次代码应该在此上下文中运行的操作完成后立即调用cancel

//func gen(ctx context.Context) <- chan int{
//	dst := make(chan int)
//	n := 1
//	go func(){
//		for {
//			select {
//			case <- ctx.Done():
//				return // return 结束该goroutine，防止泄露
//			case dst <- n:
//				n++
//			}
//		}
//	}()
//	return dst
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()	// 当取完需要的整数后调用cancel
//	for n := range gen(ctx){
//		fmt.Println(n)
//		if n == 100{
//			break
//		}
//	}
//}

// gen 函数在单独的goroutine中生成整数并将它们发送到返回的通道，gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄露

// 4.2.withDeadline的函数签名：func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
/*
返回父上下文的副本，并将deadline调整为不迟于D。如果父上下文的deadline已经早于d,则WithDeadline(parent, d)在语义上等于父上下文。当截止日期过期时，
当调用返回的cancel函数时，或者当父上下文的Done函数关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel
*/

//func main() {
//	d := time.Now().Add(500 * time.Millisecond)
//	ctx, cancel := context.WithDeadline(context.Background(), d)
//
//	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的时间
//	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
//
//	defer cancel()
//	select{
//	case <- time.After(1 * time.Second):
//		fmt.Println("overslept")
//	case <- ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}

// 上面的代码中，定义了一个50毫秒之后过期的deadline，然后调用context.WithDeadline(context.Background(), d)得到一个上下文（ctx）和一个取消
// 函数(cancel) 然后使用一个select让主程序陷入等待状态：等待1秒时间打印overslept退出或者等待ctx过期后退出。因为ctx50毫秒之后就过期。所以
// ctx.Done()会先接收到值，然后ctx.Err会打印取消的原因

// 4.3.WithTimeout func WithTimeOut(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithTimeout 返回WithDeadline(parent, time.Now().Add(timeout))
// 取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。

//var wg sync.WaitGroup
//func worker(ctx context.Context){
//LOOP:
//	for {
//		fmt.Println("db connecting...")
//		time.Sleep(time.Millisecond * 10)	// 假设正常连接数据库耗时10毫秒
//		select{
//		case <- ctx.Done():	// 50毫秒后自动调用
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker done!")
//	wg.Done()
//}
//
//func main() {
//	// 设置一个50毫秒的超时
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}

// 4.4.WithValue 能够将请求作用域的数据与Context对象建立关系。
// func WithValue(parent Context, key, val interface{}) Context
// WithValue 返回父节点的副本，其中与key关联的值为val
/*
仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递参数给函数。 ***************
所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。WithValue的用户应该为键定义自己的类型。为了
避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。或者，导出的上下文关键变量的静态类型应该是指针或接口。
*/

//type TraceCode string
//var wg sync.WaitGroup
//
//
//func worker(ctx context.Context){
//	key := TraceCode("TRACE_CODE")
//
//	traceCode, ok := ctx.Value(key).(string)
//
//	if !ok{
//		fmt.Println("invalid trace code")
//	}
//LOOP:
//	for{
//		fmt.Printf("worker, trace code:%s\n", traceCode)
//		time.Sleep(time.Millisecond * 10)	// 假设正常连接数据库耗时10毫秒
//		select{
//		case <- ctx.Done():
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker done!")
//	wg.Done()
//}
//
//func main() {
//	// 设置一个50毫秒的超时
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)
//	// 在系统的入口中设置trace code 传递给后续启动的goroutine实现日志数据聚合
//	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "1111111111111")
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}

// 5.使用context的注意事项
/*
	1.推荐以参数的方式显示传递context
	2.以context为参数的函数，应该以context为第一个参数
	3.给一个函数传递context的时候，不要传递nil，不知道传递什么的时候，传递context.TODO()
	4.context的value相关方法应该传递请求域的必要参数，比应该用于传递可选参数
	5.context是线程安全的，可以放心的在多个goroutine中是用
*/

// 6.客户端超时取消示例：调用服务端API时如何在客户端实现超时控制  调用服务端的API时如何在客户端实现超时控制？

