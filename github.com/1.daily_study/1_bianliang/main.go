package main

import "fmt"

// 1.短变量声明
// var age = 18

// func main() {
// 	n := 10
// 	m := 3

// 	fmt.Println(n, m)
// }

// 2.匿名变量，类比python
//func foo() (int, string) {
//	return 18, "小花"
//}
//
//func main() {
//	age, _ := foo()
//	_, name := foo()
//
//	fmt.Println("name=", name)
//	fmt.Println("age=", age)
//}

// 3.常量的声明
const (
	pi = 3.1415
	e  = 2.7182
)

const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)
}

// 3.1.常量中的iota
//const (
//	n1 = iota
//	n2 = 100
//	n3 = iota
//	n4
//)
//const n5 = iota
//
//const (
//	_  = iota
//	KB = 1 << (10 * iota)
//	MB = 1 << (10 * iota)
//	GB = 1 << (10 * iota)
//	TB = 1 << (10 * iota)
//	PB = 1 << (10 * iota)
//)
//
//const (
//	a, b = iota + 1, iota + 2
//	c, d
//	e, f
//)
//
//func main() {
//	fmt.Println("n1=", n1)
//	fmt.Println("n2=", n2)
//	fmt.Println("n3=", n3)
//	fmt.Println("n4=", n4)
//	fmt.Println("n5=", n5)
//	fmt.Println("KB=", KB)
//	fmt.Println("MB=", MB)
//	fmt.Println("GB=", GB)
//	fmt.Println("TB=", TB)
//	fmt.Println("PB=", PB)
//	fmt.Println("a=", a)
//	fmt.Println("d=", d)
//	fmt.Println("e=", e)
//}
