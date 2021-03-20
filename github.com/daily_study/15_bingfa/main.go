package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go中的并发学习：并发是编程里面一个非常重要的概念，go语言在语言层面天生支持并发，这也是go语言流行的一个很重要的原因

// 1.并发与并行
/*
并发：同一时间段内执行对个任务
并行：同一时刻执行对个任务
go语言中通过goroutine实现。goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。goroutine是go语言的运行
时（runtime）调度完成，而线程是由操作系统调度完成。go语言还提供channel在多个goroutine间进行通信，goroutine和channel是go语言秉承CSP并发模式
的重要实现
*/

// 2.goroutine：在Java/c++中我们在实现并发编程时，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个任务，同时需要自己去调度线程执行
// 任务并维护上下文的切换。goroutine实现了一种机制，程序员只需要定义很多个任务，让系统去帮忙我们把这些任务分配到CPU上实现并发
/*
gotoutine 的概念类似于线程，但goroutine是由go的运行时调度和管理的。go程序会智能地将goroutine中的任务合理地分配给每个CPU。go语言之所以被称为
现代化的编程语言，就是因为他在语言层面已经内置了调度和上下文切换的机制。

在go语言中不需要去自己写进程、线程、协程，只需要写好goroutine。当需要让某个任务并发执行的时候，只需要把这个任务包装成一个函数，开启一个goroutine
去执行这个函数就可以了
*/

// 2.1.使用goroutine 只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine
// 一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数

//func Hello() {
//	fmt.Println("Hello Goroutine")
//}
//
//func main() {
//	go Hello()
//	fmt.Println("main goroutine done")
//	time.Sleep(time.Second)
//}

// main函数返回的时候这个goroutine就结束了。main函数是夜王，其他的goroutine是异鬼。main函数启动快，但goroutine在创建goroutine的时候需要
// 花费一些时间。所以goroutine的输出比main函数的输出慢

// 2.2.启动多个goroutine：在go中实现并发是简单的，我们可以启动多个goroutine
//var wg sync.WaitGroup
//
//func Hello(i int) {
//	defer wg.Done()
//	fmt.Println("Hello Goroutine!", )
//}
//
//func main() {
//	for i := 1; i < 10; i++ {
//		wg.Add(1)
//		go Hello(i)
//	}
//	wg.Wait()	// 等待所有登记的goroutine都结束
//}
// 多次执行上面的代码，发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而是goroutine的调度是随机的

// 3.goroutine与线程
// 3.1.可增长的栈：os线程一般都有固定的栈内存（通常为2MB），一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈
// 不是固定的，他可以按需增大和缩小，goroutine的栈大小可以达到1GB，虽然极少会用到那么大。所以在go语言中一次创建十万左右的goroutine也是可以的

// 3.2.goroutine调度：GPM是GO语言运行时层面实现的，是GO语言自己实现的一套调度系统。区别于操作系统OS线程。
/*
G:就是这个goroutine，里面除了存放本goroutine信息外，还有与所在P的绑定等信息
P：管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针、堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度
（比如把占用CPU较长的goroutine暂停、运行后续的goroutine等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了回去其他P的队列抢任务
M：machine是go运行时对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个goroutine最终是要放在M上执行的

P与M一般也是一一对应的关系。他们的关系是：P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载
在新建的M上。当旧的G阻塞完成或者认为其已经死掉时，回收旧的M。

P的个数是通过runtime.GOMAXPROCES设定（256）go1.5版本之后默认为物理线程数，在并发量大的时候会增加一些P和M，但不会太对，切换太频繁的话，得不偿失

单从线程调度讲，go语言相比其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由go运行时自己的调度器调度的，这个调度器使用的是一个称为m:n
调度的技术（复用/调度m个goroutine到n个OS线程）。其一大特点是goroutine的调度是在用户态下完成的，不涉及内核态与用户态之间的频繁切换，包括内存的
分配与释放，都是在用户态维护着一块大的内存池，不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。另一方面充分利用了多核的
硬件资源，近似的把若干个goroutine均分在物理线程上，再加上goroutine的超轻量
*/

// 3.2.GOMAXPROCES
/*
go运行时的调度器使用GOMAXPROCES参数来确定需要使用多少个OS线程来同时执行go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把go
代码同时调度到8个OS线程上（GOMAXPROCES是m:n中的n）
go语言中可以通过runtime.GOMAXPROCES()函数设置当前程序并发时占用的CPU逻辑核心数
go1.5版本之前，默认使用的是单核心执行。go1.5版本之后，默认使用全部的CPU逻辑核心数
*/

// 3.3.可以将任务分配到不同的CPU逻辑核心上实现并行的效果，
//func a()  {
//	for i := 1; i < 10; i++ {
//		fmt.Println("A", i)
//	}
//}
//
//func b()  {
//	for i := 1; i < 10; i++ {
//		fmt.Println("B", i)
//	}
//}
//
//func main() {
//	runtime.GOMAXPROCS(1)
//	go b()
//	go a()
//	time.Sleep(time.Second)
//}

// 3.4.两个任务只有一个逻辑核心，此时是做完一个任务再做两一个任务。将逻辑核心数设置为2，此时两个任务并行执行，
//func a() {
//	for i := 1; i < 10; i++ {
//		fmt.Println("A", i)
//	}
//}
//
//func b() {
//	for i := 1; i < 10; i++ {
//		fmt.Println("B", i)
//	}
//}
//
//func main() {
//	runtime.GOMAXPROCS(2)
//	go a()
//	go b()
//	time.Sleep(time.Second)
//}

// go语言中的操作系统线程和goroutine的关系
// 1.一个操作系统线程对应用户态多个goroutine
// 2.go程序可以同时使用多个操作系统线程
// 3.goroutine和OS线程是多对多的关系，即m:n

// 4.channel
/*
单纯地将函数并发执行时没有意义的，函数与函数间需要交换数据才能体现并发函数的意义
虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竟态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种
做法势必会造成性能问题

go语言的并发模型csp（communicating Sequential Processes）,提倡通过通信共享内存而不是共享内存而实现通信
如果说goroutine是go程序并发的执行体，channel就是他们的连接，channel是一个goroutine发送特定值到另一个goroutine的通信机制

go语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管
也就是声明channel的时候需要为其制定元素类型
*/
//channel类型：channel是一种类型，一种引用类型。声明通道类型的格式如下
//var 变量 chan 元素类型

// var ch1 chan int // 声明一个传递整型的通道
// var ch2 chan bool // 声明一个传递布尔类型的通道
// var ch3 chan []int // 声明一个传递int切片的通道

// 4.1.创建channel 通道是引用类型，通道类型的空值是nil
//func main() {
//	var ch chan int
//	fmt.Println(ch)
//}

// 声明的通道后需要使用make函数初始化之后才能使用
// 创建channel的格式如下 make(chan 元素类型，[缓冲大小]) channel的缓冲大小是可选的
/*
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
*/

// 4.2.channel操作：发送（send）、接收（receive）、关闭（close）三种操作
// 发送和接收都是用 <- 符号
//func main() {
//    // 定义一个通道
//	ch := make(chan int)
//	// 发送
//	ch <- 10
//	// 接收 从一个通道中接收值
//	x := <- ch
//	fmt.Println(x)
//	// 关闭 通过调用内置的close函数来关闭通道
//	close(ch)
//}
// 关闭通道要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的
// 在结束操作之后关闭文件是必须要做的事情，但关闭通道不是必须的

/*
关闭通道之后有以下特点
1. 对一个关闭的通道再发送值就会导致panic
2. 对一个关闭的通道进行接收会一致获取值直到通道为空
3. 对一个关闭的并且没有值得通道执行接收操作会得到对应类型的零值
4. 关闭一个已经关闭的通道会导致panic
*/

// 4.3.无缓冲的通道：无缓冲的通道又被称为阻塞的通道。
//func main() {
//	ch := make(chan int)
//	ch <- 10
//	fmt.Println("发送成功")	// fatal error: all goroutines are asleep - deadlock!
//}
/*
为什么上述的代码会报错。
因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。无缓冲的通道必须有接收才能发送
上面的代码阻塞在 ch <- 10这一行代码形成死锁
*/

//func recv(c chan int) {
//	ret := <- c
//	fmt.Println("接收成功", ret)
//}
//
//func main() {
//	ch := make(chan int)
//	go recv(ch) // 启用goroutine通道接收值
//	ch <- 10
//	fmt.Println("发送成功")
//}

// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先
// 执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

// 使用无缓冲通道进行信号将导致发送和接收同步化，因此，无缓冲通道也称为同步通道

// 4.4.有缓冲的通道：为解决上面的问题还有一种就是使用有缓冲区的通道，我们可以使用make函数初始化通道的时候为其指定通道的容量
//func main() {
//	ch := make(chan int, 1)
//	ch <- 10
//	fmt.Println("发送成功")
//}

// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。我们可以通过内置的len函数获取通道内元素的数量，使用cap
// 函数获取通道的容量。

// 4.5.for range从通道循环取值：当向通道中发送完数据时，我们可以通过close函数来关闭通道。
// 当通道被关闭时，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值
// channel 练习
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	// 开启goroutine将0~10的数发送到ch1中
//	go func() {
//		for i := 1; i < 10; i++ {
//			ch1 <- i
//		}
//		close(ch1)
//	}()
//	// 开启goroutine从ch1中接收的值，并将该值得平方发送到ch2中
//	go func() {
//		for {
//			i, ok := <- ch1
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//		}
//		close(ch2)
//	}()
//	// 在主goroutine中从ch2中接收值打印
//	for i := range ch2{
//		fmt.Println(i)
//	}
//}
// 从上面的例子中我们看到有两种方式在接收值得时候判断该通道是否被关闭，不过我们通常使用的是for range的方式。使用for range遍历通道，当通道被关闭
// 的时候就会退出for range

// 4.6.单向通道：有的时候我们会将通道做为参数在对个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能
// 发送或者接收。
//func counter(out chan <- int){
//	for i := 0; i < 100; i++{
//		out <- i
//	}
//	close(out)
//}
//
//func squarer(out chan <- int, in <- chan int){
//	for i := range in {
//		out <- i * i
//	}
//	close(out)
//}
//
//func printer(in <- chan int){
//	for i := range in {
//		fmt.Println(i)
//	}
//}
//
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go counter(ch1)
//	go squarer(ch2, ch1)
//	printer(ch2)
//}

/*
其中：
	1.chan <- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作
	2.<- chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作
在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但是反过来是不可以的
*/

/*
channel 异常情况总结
channel		nil		非空			空的			满了			没满
接收			阻塞		接收值		阻塞			接收值		接收值
发送			阻塞		接收值		发送值		阻塞			发送值
关闭			panic	关闭成功		关闭成功		关闭成功		关闭成功
					读完数据后	返回零值		读完数据后	读完数据后
					返回零值					返回零值		返回零值
关闭已经关闭的channel也会引发panic
*/

// 5.worker pool（goroutine 池）我们在工作中通常会使用可以指定启动的goroutine数量- work pool模式，通过控制goroutine的数量，防止goroutine泄露和暴涨
//func worker(id int, jobs <- chan int, results chan <- int) {
//	for j := range jobs {
//		fmt.Printf("worker: %d start job: %d\n", id, j)
//		time.Sleep(time.Second)
//		fmt.Printf("worker: %d end job: %d\n", id, j)
//		results <- j * 2
//	}
//}
//
//func main() {
//	jobs := make(chan int, 100)
//	results := make(chan int, 100)
//
//	// 开启3个goroutine
//	for w := 1; w <= 3; w++{
//		go worker(w, jobs, results)
//	}
//	// 5个任务
//	for j := 1; j <= 5; j++{
//		jobs <- j
//	}
//	close(jobs)
//	// 输出结果
//	for a:= 1; a <= 5; a++{
//		<- results
//	}
//}

// 6.select多路复用：在某些场景下我们需要同时从多个通道接收数据，通道在收数据时，如果没有发生阻塞
//for {
//	// 尝试从ch1接收数值
//	data, ok := <- ch1
//	// 尝试从ch2接收值
//	data, ok := <- ch2
//}
// 这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多，为了应付这种场景，go内置了select关键字，可以同时相应多个通道的操作
// select的使用类似于switch语句，他有一系列case分支和一个默认的额分支。每个case会对应一个通道的通信（接收或发送）过程。select会一致等待，直到
// 某个case的通信操作完成时，就会执行case分支对应的语句
/*
select {
	case <- ch1:
		...
	case data := <- ch2:
		...
	case ch3 <- data:
		...
	default:
		默认操作
}
*/
//func main() {
//	ch := make(chan int, 1)
//	for i := 0; i < 10; i++ {
//		select {
//		case x := <- ch:
//			fmt.Println(x)
//			case ch <- i:
//		}
//	}
//}
/*
使用select语句能提高代码的刻度性
- 可以处理一个或多个channel的发送/接收操作
- 如果多个case同时满足，select会随机选择一个
- 对于没有case的select{}会一直等待，可用于阻塞main的函数
*/

// 7.并发安全和锁：有时候在go代码中可能会存在多个goroutine同时操作一个资源（临界区）,这种情况会发生竟态问题（数据竟态）。'
//var x int64
//var wg sync.WaitGroup
//
//func add() {
//	for i := 1; i < 5000; i++ {
//		x += 1
//	}
//	wg.Done()
//}
//func main() {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

// 上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符

// 7.1.互斥锁：互斥锁是一种常见的共享资源的访问方式，他能保证在同时只有一个goroutine可以访问共享资源。go语言中使用sync包的Mutex类型来实现
// 互斥锁
//var x int64
//var wg sync.WaitGroup
//var lock sync.Mutex
//
//func add() {
//	for i := 1; i < 5000; i++ {
//		lock.Lock()	// 加锁
//		x += 1
//		lock.Unlock() // 解锁
//	}
//	wg.Done()
//}
//func main() {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

// 使用互斥锁能保证同一个时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区
// 多个goroutine同时等待一个锁时，唤醒的策略是随机的。

// 7.2.读写互斥锁：互斥锁是完全互斥的，但是有很多实际的情况下是读多写少，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下，
// 使用读写锁是更好的选择。读写锁在go中使用sync包中的RWMutex类型。
//var (
//	x int64
//	wg sync.WaitGroup
//	lock sync.Mutex
//	rwlock sync.RWMutex
//)
//
//func writer() {
//	lock.Lock()	// 加互斥锁
//	//rwlock.Lock()	// 加写锁
//	time.Sleep(10 * time.Millisecond)	// 假设读操作耗时10秒
//	//rwlock.Unlock()	// 解写锁
//	lock.Unlock()	// 解互斥锁
//	wg.Done()
//}
//
//func read() {
//	lock.Lock()	// 加互斥锁
//	//rwlock.Lock()	// 加读锁
//	time.Sleep(time.Millisecond)	// 假设操作时间为1毫秒
//	//rwlock.Unlock()	// 解读锁
//	lock.Unlock()	// 解互斥锁
//	wg.Done()
//}
//
//func main() {
//	start := time.Now()
//	for i := 0; i < 10; i++{
//		wg.Add(1)
//		go writer()
//	}
//	for i := 0; i < 100000; i++ {
//		wg.Add(1)
//		go read()
//	}
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//}
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作区别不大，读写锁的优势就发挥不出来了

// 7.3.sync.WaitGroup,在代码中生硬的使用time.Sleep肯定是不合适的，go语言中可以使用 sync.WaitGroup来实现并发任务的同步
/*
方法名						功能
Add(delta int)				计数器+delta
Done()						计数器-1
Wait()						阻塞到计数器变为0
*/
// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。当启动了N个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将
// 计数器减1，通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有的并发任务已经完成

//var wg sync.WaitGroup
//
//func hello() {
//	defer wg.Done()
//	fmt.Println("Hello Goroutine")
//}
//
//func main() {
//	wg.Add(1)
//	go hello()	// 启动另外一个goroutine去执行hello函数
//	fmt.Println("main goroutine done!")
//	wg.Wait()
//}
// sync.WaitGroup是一个结构体，传递的时候要传指针

// 7.4. sync.Once 这是一个进阶知识点
// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，比如只加载一次配置文件、只关闭一次通道等。
// go预言者中sync包中提供了一个针对只执行一次场景的解决方案- sync.Once 它只有一个Do方法
// func (o *Once) Do(f func()) {} 如果要执行的函数f 需要传递参数就需要搭配闭包来使用

// 加载配置文件示例
/*
延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践，因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且
有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的
*/

// 7.4.1.
//var icons map[string]image.Image
//
//func loadIcons() {
//	icons = map[string]image.Image{
//		"left": loadIcon("left.png"),
//		"up": loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down": loadIcon("down.png"),
//	}
//}
//
//// Icon 被多个goroutine调用时并不是并发安全的
//func Icon(name string) image.Image {
//	if icons == nil{
//		loadIcons()
//	}
//	return icons[name]
//}

// 7.4.2.
// 对个goroutine并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序。
// loadIcons函数可能会被重排为以下结果

//func loadIcons() {
//	icons = make(map[string]image.Image)
//	icons["left"] = loadIcon("left.png")
//	icons["up"] = loadIcon("up.png")
//	icons["right"] = loadIcon("right.png")
//	icons["down"] = loadIcon("down.png")
//}

//在这种情况下就会出现判断了icons不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法是添加互斥锁，保证初始化icons的时候不会被
// 的goroutine操作，但是这样做优惠引发性能问题

// 使用sync.Once改造的代码如下
//var icons map[string]image.Image
//
//var loadIconsOnce sync.Once
//
//func loadIcons() {
//	icons = map[string]image.Image{
//		"left": loadIcon("left.png"),
//		"up": loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down": loadIcon("down.png"),
//	}
//}
//// Icon 是并发安全的
//func Icon(name string) image.Image{
//	loadIconsOnce.Do(loadIcons)
//	return icons[name]
//}

// 7.4.3.并发安全的单例模式：借助sync.Once实现的并发安全的单例模式
//type singleton struct {}
//
//var instance *singleton
//var once sync.Once
//
//func GetInstance() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}

// sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发
// 安全的并且初始化操作也不会被执行多次

// 7.4.4.sync.Map: go语言中内置的map不是并发安全的
//var m = make(map[string]int)
//
//func get(key string) int{
//	return m[key]
//}
//
//func set(key string, value int){
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 2; i++{
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key, n)
//			fmt.Printf("k=%v, v=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

// 上面的代码开启少量几个goroutine的时候可能没有什么问题，当并发对了之后执行上面的代码就会出现fatal error: concurrent map writes 错误

// 像这种场景下就需要为map加锁来保证并发的安全性了，go语言的sync包中提供了一个开箱即用的并发安全版map-sync.Map。开箱即用表示不用像内置的map一样使用
// make函数初始化就能直接使用。同时sync.map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++{
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=%v, v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}