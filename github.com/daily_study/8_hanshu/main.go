package main

// 函数相关 go语言中支持函数，匿名函数和闭包，并且函数在go语言中属于“一等公民”

// 1.函数的定义
/*
func 函数名(参数) (返回值) {
	函数体
}
- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名称不能重复
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用，分隔
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用，分隔
- 函数体：实现指定功能的代码块
*/

// func intSum(x int, y int) int {
// 	return x + y
// }

// 1.1.函数的参数和返回值都是可选的，
// func sayHello() {
// 	fmt.Println("hello 沙河")
// }

// func main() {
// 	a := intSum(10, 5)
// 	sayHello()
// 	fmt.Println(a)
// }

// 2.参数
// 2.1.参数中如果相邻变量的类型是相同的,可以省略类型
// func intSum(x, y int) int {
// 	return x + y
// }

// 2.2.可变参数：指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识   ps：可变参数通常要作为函数的最后一个参数。
// 2.3.固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
// 重点：函数的可变参数时通过切片来实现的。
// func intSum2(x ...int) int {
// 	fmt.Println(x) // x是一个切片

// 	sum := 0
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum
// }

// func main() {
// 	ret1 := intSum2()
// 	ret2 := intSum2(10)
// 	ret3 := intSum2(10, 20)
// 	ret4 := intSum2(10, 20, 30)
// 	fmt.Println(ret1, ret2, ret3, ret4)
// }

// 3.变量作用域
// 3.1.全局变量：全局变量定义在函数外部的变量，他在程序整个运行周期内都有效。在函数中可以访问到全局变量

// 定义全局变量
// var num int64 = 10

// func testGlobalVar() {
// 	fmt.Printf("num=%d \n", num)
// }

// func main() {
// 	testGlobalVar()
// }

// 3.2.局部变量：局部变量分为两种，函数内定义的变量无法在该函数外使用；如果局部变量和全局变量重名，优先访问局部变量
// 在语句块中定义的变量，通常会在if条件判断、for循环、switch语句上使用这种定义变量的方式
// func testLocalVar() {
// 	// 定义一个函数局部变量x,仅在该函数内生效
// 	var x int64 = 100
// 	fmt.Printf("x=%d \n", x)
// }

// func main() {
// 	testLocalVar()
// 	// fmt.Println(x)  // 不能引用在其他函数中声明的局部变量
// }

// 3.2.1.优先访问局部变量的例子
// var num int64 = 10

// func testNum() {
// 	num := 100
// 	fmt.Printf("x=%d \n", num)
// }

// func main() {
// 	testNum()
// }

// 4.函数类型和变量
// 4.1.定义函数类型：可以使用type关键字来定义一个函数类型，type calculation func(int, int) int; 上面的语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并返回一个int类型的返回值
//func add(x, y int) int {
//	return x + y
//}

// func sub(x, y int) int {
// 	return x - y
// }
// add和sub都能赋值给calculation类型的变量

// 4.2.函数类型变量：声明函数类型的变量并且为该变量赋值
//type calculation func(int, int) int
//
//func main() {
//	var c calculation
//	c = add
//	fmt.Printf("type of c: %T \n", c)
//	fmt.Println(c(1, 2))
//
//	f := add
//	fmt.Printf("type of f: %T\n", f)
//	fmt.Println(f(10, 20))
//}

// 5.高阶函数：高阶函数分为函数作为参数和函数作为返回值两部分
// 5.1.函数作为参数
// func add(x, y int) int {
// 	return x + y
// }

// func calc(x, y int, op func(int, int) int) int {
// 	return op(x, y)
// }

// func main() {
// 	ret2 := calc(10, 20, add)
// 	fmt.Println(ret2)
// }

// 5.2.函数作为返回值
// func do(s string) (func (int, int) int, error {
// 	switch s {
// 	case "+":
// 		return add, nil
// 	case "-":
// 		return sub, nil
// 	default:
// 		err := errors.new("无法识别操作符")
// 		return nil, err
// 	}
// })

// 6.匿名函数和闭包
// 6.1.匿名函数：函数可以作为返回值，但是在go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
// 匿名函数就是没有名的函数， func (参数) (返回值) {函数体}
// 匿名函数因为没有函数名，所以没有办法像普通函数那样被调用，所以匿名函数需要保存到某个变量或者作为立即执行函数：
// func main() {
// 	// 将匿名函数保存到变量中
// 	add := func(x, y int) {
// 		fmt.Println(x + y)
// 	}

// 	add(10, 20) // 通过变量调用匿名函数

// 	// 自执行函数：匿名函数定义完 add() 直接执行
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(10, 20)
// }
// 匿名函数多用于实现回调函数和闭包

// 6.2.闭包：闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
//func adder() func(int) int {
//	var x int
//	return func(i int) int {
//		x += i
//		return x
//	}
//}
// // 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。在f的声明周期内，变量x也一直有效
//func main() {
//	var f = adder()
//	fmt.Println(f(10))
//	fmt.Println(f(20))
//	fmt.Println(f(30))
//
//	f1 := adder()
//	fmt.Println(f1(40))
//	fmt.Println(f1(50))
//}

// 6.2.1.闭包进阶1
//func adder2(x int) func(int) int {
//	return func(i int) int {
//		x += i
//		return x
//	}
//}

//func main() {
//	var f = adder2(10)
//	fmt.Println(f(10))
//	fmt.Println(f(20))
//	fmt.Println(f(30))
//
//	f1 := adder2(20)
//	fmt.Println(f1(40))
//	fmt.Println(f1(50))
//}

// 6.2.2.闭包进阶示例2
// func makeSuffixFunc(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// 6.2.3.闭包进阶示例3
// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}

// 	return add, sub
// }

// func main() {
// 	f1, f2 := calc(10)
// 	fmt.Println(f1(1), f2(2))
// 	fmt.Println(f1(3), f2(4))
// 	fmt.Println(f1(5), f2(6))
// }

// 7.defer语句：go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行。
//func main() {
//	fmt.Println("start")
//	defer fmt.Println(1)
//	defer fmt.Println(2)
//	defer fmt.Println(3)
//	fmt.Println("end")
//}

// 7.1.defer执行时机：在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回赋值操作后，RET指令前
// defer经典案例
//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	fmt.Printf("f1的x值为：%d \n", x)
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	fmt.Printf("f2的x值为：%d \n", x)
//	return x
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	fmt.Printf("f3的x值为：%d \n", x)
//	return x
//}
//
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	fmt.Printf("f4的x值为：%d \n", x)
//	return 5
//}
//
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}

// 7.2.defer的面试题:defer注册要延迟执行的函数时该函数所有的参数都需要确定值
// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main() {
// 	x := 1
// 	y := 2
// 	defer calc("AA", x, calc("A", x, y))
// 	x = 10
// 	defer calc("BB", x, calc("B", x, y))
// 	y = 20
// }

// 8.内置函数介绍 panic 和 recover 用来做错误处理
// panic 可以在任何地方引发，但recover只有在defer调用的函数中有效。通过recover将程序恢复
// recover() 必须搭配defer使用；defer一定要在可能引发panic的语句之前定义
//func funcA() {
//	fmt.Println("func A")
//}
//
//func funcB() {
//
//	defer func() {
//		err := recover()
//		// 如果程序出现可panic错误，可以通过recover恢复过来
//		if err != nil {
//			fmt.Println("recover in B")
//		}
//	}()
//	panic("panic in B")
//}
//
//func funcC() {
//	fmt.Println("func C")
//}
//
//func main() {
//	funcA()
//	funcB()
//	funcC()
//}

// 练习题1
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

//var (
//	coins = 50
//	users = []string{
//		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
//	}
//	distribution = make(map[string]int, len(users))
//)
//
//func dispatchCoin() int {
//	coin := 50
//	for _, name := range users {
//		for _, str := range strings.Split(name, "") {
//			switch {
//			case str == "e" || str == "E":
//				distribution[name]++
//				coin--
//			case str == "i" || str == "I":
//				distribution[name] += 2
//				coin -= 2
//			case str == "o" || str == "O":
//				distribution[name] += 3
//				coin -= 2
//			case str == "u" || str == "U":
//				distribution[name] += 4
//				coin -= 4
//			}
//		}
//		fmt.Printf("%v 用户所得金币为：%d\n", name, distribution[name])
//	}
//	return coin
//}
//
//func main() {
//	left := dispatchCoin()
//	fmt.Println("剩下：", left)
//}


/*
函数总结：
	1.匿名函数：没有函数名的函数，需要保存到某个变量或者立即执行；实现回调函数和闭包
	2.闭包：一个函数和其引用相关的环境结合成的实体。闭包=函数+引用环境
	3.defer是在值赋值和return之前进行的操作，在所有值确定的时候才会执行defer的内容，遵循先进后出的策略
	4.panic 可以在任何地方引发，但recover只有在defer调用的函数中有效。通过recover将程序恢复
	5.recover() 必须搭配defer使用；defer一定要在可能引发panic的语句之前定义
	6.高阶函数：高阶函数分为函数作为参数和函数作为返回值
*/