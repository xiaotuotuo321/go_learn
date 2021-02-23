package main

// go中的指针相关
// 区别于C/C++中的指针,Go语言中的指针不能进行偏移和运算,是安全指针
// go中指针的三个概念：指针地址、指针类型和指针取值

/*
1.指针的概念：任何程序数据载入内存后，在内存中都有他们的地址，这就是指针。为了保存一个数据在内存中的地址，我们需要指针变量。
	& 取地址；*根据地址取值
*/

// 1.指针地址和指针类型：每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
// 取变量指针的语法： ptr := &v   // v的类型为T
// v:代表被取地址的变量，类型为T
// ptr:用于接收地址的变量，ptr的类型就为*T,称做T的指针类型。*代表指针。
//func main() {
//	a := 10
//	b := &a
//
//	fmt.Printf("a:%d ptr:%p\n", a, &a)
//	fmt.Printf("b:%p type:%T\n", b, b)
//	fmt.Println(&b)
//}

// 2.指针取值：在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值。
/*
变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	1. 对变量取地址操作，可以获得这个变量的指针变量
	2. 指针变量的值是指针地址
	3. 对指针变量进行取值操作，可以获得指针变量指向原变量的值
*/
//func main() {
//	a := 10
//	b := &a // 取变量a的地址，将指针保存到b中
//	fmt.Printf("type of b:%T\n", b)
//    c := *b // 指针取值（根据指针去内存取值）
//    fmt.Printf("type of c:%T\n", c)
//    fmt.Printf("value of c:%v\n", c)
//}

// 2.1.指针传值示例
//func modfy1(x int) {
//	x = 100
//}
//
//func modfy2(x *int) {
//	*x = 100
//}
//
//func main() {
//	a := 10
//	modfy1(a)
//	fmt.Println(a)
//
//	modfy2(&a)
//	println(a)
//}

// 3.new 和 make
// 下面的代码会出现panic,因为在go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没有办法存储。
// 而对于值类型的声明不需要分配内存空间，因为他们在声明的时候已经默认分配好了内存空间。要分配内存就会有new和make.
//func main() {
//	var a *int
//	*a = 100
//	println(*a)
//
//	var b map[string] int
//	b["沙河娜扎"] = 100
//	fmt.Println(b)
//}

// 3.1.new是一个内置函数，它的函数签名如下： func new(Type) * Type
// Type 表示类型，new函数只接受一个参数，这个参数是一个类型；*Type表示类型指针，new函数返回一个指向该类型内存地址的指针
// new函数不太常用，使用new函数得到的是一个类型的指针,并且该指针对应的值为该类型的零值。
//func main() {
//	a := new(int)
//	b := new(bool)
//	fmt.Printf("%T\n", a)
//	fmt.Printf("%T\n", b)
//	println(*a)
//	println(*b)
//}

// 3.1.1.在本节最开始的时候，var a *int 只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以赋值
//func main() {
//	var a *int
//	a = new(int)
//	*a = 10
//	fmt.Println(*a)
//}

// 3.2.make也是用于内存分配的，区别于new,它只能用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型,
// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
// func make(t Type, size ...IntegerType) Type
// make 函数是无可替代的，在使用slice、map以及channel的时候，都需要使用make进行初始化,然后才可以对他们进行操作
//func main() {
//	var b map[string] int
//
//	b = make(map[string] int, 10)
//	b["沙河娜扎"] = 100
//	fmt.Println(b)
//}

// new和make的区别
/*
1.二者都是用来做内存分配的
2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
3.而new用于类型的内存分配，并且内存对应的值为类型零值。返回的是指向类型的指针。
*/