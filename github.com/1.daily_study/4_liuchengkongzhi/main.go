package main

// 1. if else

// func main() {
// 	score := 40

// 	if score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}
// }

// 1.1.测试if 内声明变量

// func main() {
// 	if score := 65; score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}

// 	// if中声明的变量只在流程中存活，流程一旦结束，内存会被回收，变量会被销毁
// 	// fmt.Println(score)
// }

// 2.for循环
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// 2.1.for循环的几种形式
// func main() {
// 	i := 0
// 	for ; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	i := 0
// 	for i < 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 3.switch case 每个switch只能有一个default分支

// func main() {
// 	finger := 3
// 	switch finger {
// 	case 1:
// 		fmt.Println("大拇指")
// 	case 2:
// 		fmt.Println("食指")
// 	case 3:
// 		fmt.Println("中指")
// 	case 4:
// 		fmt.Println("无名指")
// 	case 5:
// 		fmt.Println("小拇指")
// 	default:
// 		fmt.Println("无效的输入！")
// 	}
// }

// 3.1.一个分支可以有多个值
// func main() {
// 	switch n := 0; n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Printf("%d是奇数\n", n)
// 	case 2, 4, 6, 8, 0:
// 		fmt.Printf("%d是偶数", n)
// 	default:
// 		fmt.Println(n)
// 	}
// }

// 3.2.case 中为表达式
// func main() {
// 	age := 30
// 	switch {
// 	case age <= 25:
// 		fmt.Println("好好学习吧")
// 	case age >= 60:
// 		fmt.Println("好好享受吧")
// 	case age > 25 && age < 60:
// 		fmt.Println("好好工作吧")
// 	default:
// 		fmt.Println("活着真好呀！")
// 	}
// }

// 3.3. fallthrough 语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
//func main() {
//	s := "a"
//	switch {
//	case s == "a":
//		fmt.Println("a")
//		fallthrough
//	case s == "b":
//		fmt.Println("b")
//	case s == "c":
//		fmt.Println("c")
//	default:
//		fmt.Println(s)
//	}
//}

// 4.goto 跳转到指定标签 goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。GO语言中使用goto语句能简化一些代码的实现过程。

// 4.1.未使用goto时的双层for循环退出
// func main() {
// 	var breakFlag bool
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Println(j)
// 			if j == 2 {
// 				// 设置退出标签
// 				breakFlag = true
// 				break
// 			}
// 		}
// 		// 外层for循环判断
// 		if breakFlag {
// 			break
// 		}
// 	}
// }

// 4.2.使用goto简化代码
//func main() {
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 10; j++ {
//			if j == 2 {
//				// 设置退出标签
//				goto breakTag
//			}
//			fmt.Printf("%v-%v \n", i, j)
//		}
//	}
//
//	return
//
//breakTag:
//	fmt.Println("结束for循环")
//}

// 5.break跳出循环，可以结束for,switch,select的代码块
// func main() {
// BREAKMAIN: // 添加标签的形式；标识代码块
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				break BREAKMAIN
// 			}
// 			fmt.Printf("%v-%v \n", i, j)
// 		}
// 	}
// 	fmt.Println("...")
// }

// 6.continue 继续下次执行 仅在for循环中使用, 在continue语句侯添加标签时，表示开始标签对应的循环
//func main() {
//forloop1:
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 10; j++ {
//			if i == 2 && j == 2 {
//				continue forloop1
//			}
//			fmt.Printf("%v-%v\n", i, j)
//		}
//	}
//}

// 练习题1：打印乘法口诀表
//func main() {
//	for i := 1; i < 10; i++ {
//		for j := 1; j <= i; j++ {
//			fmt.Printf("%v * %v = %v   ", j, i, i*j)
//		}
//		fmt.Println()
//	}
//}
