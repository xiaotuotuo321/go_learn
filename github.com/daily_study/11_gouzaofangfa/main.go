package main

import "fmt"

// 方法和接收者：go中的方法是一种作用于特定类型变量的函数。这种特定类型叫做接收者
// 语法：func(接收者变量，接收者类型) 方法名(参数列表) (返回参数){ 函数体 }
// 接收者变量：官方建议使用接收者类型名称首字母的小写
// 接收者类型：接收者类型和参数类似；指针类型和非指针类型
// 防范名、参数列表、返回参数；

// 1.举例
type Person struct {
	name string
	age int8
}

// 1.1.构造函数
func NewPerson(name string, age int8) *Person{
	return &Person{
		name: name,
		age: age,
	}
}

// 1.2.Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("小王子", 18)
	p1.Dream()
}