package main

import "fmt"

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
func main() {
    // 定义一个通道
	ch := make(chan int)
	// 发送
	ch <- 10
	// 接收 从一个通道中接收值
	x := <- ch
	fmt.Println(x)
	// 关闭 通过调用内置的close函数来关闭通道
	close(ch)
}

