package main

import "fmt"

// defer理解
/*
特性
1. 关键字defer用于注册延迟调用
2. 这些调用知道return前才被执行。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了

用途：
1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放
*/

// 1.正常的defer代码 4 3 2 1 0
// func main() {
// 	var users [5]struct{}
// 	for i := range users {
// 		defer fmt.Println(i)
// 	}
// }

// 2.把上面的代码换成defer闭包 4 4 4 4 4 函数正常执行，由于闭包用到的变量i在执行的时候已经变成4，所以输出全部都是4。
// func main() {
// 	var users [5]struct{}

// 	for i := range users {
// 		defer func() { fmt.Println(i) }()
// 	}
// }

// 3.不用闭包换成函数
// func Print(i int) {
// 	fmt.Println(i)
// }
// func main() {
// 	var users [5]struct{}

// 	for i := range users {
// 		defer Print(i)
// 	}
// }

// 4.defer调用引用结构体函数
//type Users struct {
//	name string
//}
//
//func (t *Users) GetName() { // * 是传址的意思 引用Users
//	fmt.Println(t.name)
//}
//
//func main() {
//	list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
//
//	for _, t := range list {
//		defer t.GetName()
//	}
//}

// 5.在4的基础上换一种调用方式
type Users struct {
	name string
}

func (t *Users) GetName() {
	fmt.Println(t.name)
}

func GetName(t Users) {
	t.GetName()
}

func main() {
	list := []Users{{"乔峰"}, {"慕容复"}, {"清风杨"}}

	for _, t := range list {
		defer GetName(t)
	}
}