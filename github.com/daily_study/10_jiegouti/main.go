package main

import "fmt"

// 结构体：struct
// go语言中没有"类"的概念,也不支持"类"的继承等面向对象的概念。go语言中通过结构体的内嵌在配合接口比面向对象具有更改的扩展性和灵活性。

// 1.类型别名和自定义类型
// 1.1.自定义类型
// type MyInt int type关键字的定义，MyInt就是一种新的类型，它具有int的特性

// 1.2.类型别名
// type TypeAlias = Type
// type byte = uint8
// type rune = int32

// 1.3.类型定义和类型别名的区别
// 类型别名与类型定义表面上只有一个等号的差异; MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
//type NewInt int
//type myInt = int
//func main() {
//	var a NewInt
//	var b myInt
//
//	fmt.Printf("type of a:%T \n", a)
//	fmt.Printf("type of b:%T \n", b)
//}

// 2.结构体：go语言中的基础数据类型可以表示一些事务的基本属性，但是当我们想表达一个事物的全部或者部分属性时，这时候再用单一的基本数据类型明显就无法
// 满足需求了，go语言中提供了一种自定义数据类型，可以封装多个基本的数据类型，这种数据类型叫做结构体
// go 语言中通过struct来实现面向对象

// 2.1.结构体的定义：使用type和struct关键字来定义结构体
/*
type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
}
类型名: 标识自定义结构体的名称，在同一个包内不能重复。
字段名：表示结构体字段名。结构体中的字段名必须唯一
字段类型：表示结构体字段的具体类型
*/
//可以定义一个Person的结构体
//type Person struct {
//	name string  // 同类型的可以写在一行
//	city string
//	age int
//}
// 语言内置的数据类型是用来描述一个值的，而结构体是用来描述一组值的。本质上是一种聚合型的数据类型

// 3.结构体实例化：只有当结构体实例化时才会真正的分配内存。也就是实例化后才能使用结构体的字段。
// 结构体本身也是一种类型，可以使用像声明内置类型一样使用var 关键字来声明结构体类型 var 结构体实例结构体类型
//type Person struct {
//	name, city string
//	age int
//}
//
//func main() {
//	var p Person
//	p.name = "娜扎"
//	p.city = "新疆"
//	p.age= 18
//
//	fmt.Printf("p=%v\n", p)
//	fmt.Printf("p=%#v\n", p)
//}

// 3.1.匿名结构体：在定义一些临时数据结构等场景下还可以使用匿名结构体
//func main() {
//	var user struct{Name string; Age int}
//
//	user.Name = "小王子"
//	user.Age = 3
//
//	fmt.Printf("user=%#v \n", user)
//}

// 3.2.创建指针类型结构体： 可以使用new关键字对结构体进行实例化，得到结构体的地址
//func main() {
//	var p = new(Person)
//	fmt.Printf("type of p: %T\n", p)
//	fmt.Printf("p= %#v\n", p)
//}

// 3.2.1.在Go语言中支持对结构体指针直接使用.来访问结构体成员
//func main() {
//	var p = new(Person)
//	p.name = "小王子"
//	p.city = "地球"
//	p.age = 3
//	fmt.Printf("p=%#v\n", p)
//}

// 3.2.2.取结构体的实例化地址：使用&对结构体进行取地址操作相当于对结构体进行一次new的操作
//func main() {
//	p := &Person{}
//	fmt.Printf("%T\n", p)
//	fmt.Printf("p=%#v\n", p)
//	p.name = "温"
//	p.city = "北京"
//	p.age = 18
//	fmt.Println(p)
//	fmt.Printf("p=%#v\n", p)
//}

// 4.结构体的初始化
//type person struct {
//	name string
//	city string
//	age int8
//}

//func main() {
	// 4.1.没有初始化的结构体，其成员变量都是对应类型的零值
	//var p1 person
	//fmt.Printf("p1=%#v", p1)

	// 4.2.使用键值对初始化：使用键值对初始化结构体时，键对应结构体的字段，值对应该字段的初始值
	//p2 := person{
	//	name: "小王子",
	//	city: "北京",
	//	age: 18,
	//}
	//fmt.Printf("p2=%#v \n", p2)

    // 4.3.也可以对结构体指针进行键值对初始化，
	//p3 := &person{
	//	name: "小王子",
	//	city: "北京",
	//	age: 18,
	//}
	//fmt.Printf("p3=%#v", p3)

    // 4.4.当某些字段没有初始值的时候，该字段可以不写。此时没有指定初始值的字段的值就是该字段类型的零值
	//p4 := person{
	//	city: "北京",
	//}
	//fmt.Printf("p4=%#v", p4)

	// 4.5.使用值的列表初始化，在初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
	/* 注意：
		1.必须初始化结构体的所有字段
		2.初始值的填充顺序必须与字段在结构体中的声明顺序一致
		3.该方式不能和键值初始化方式混用
	*/
	//p5 := &person{
	//	"沙河娜扎",
	//	"",
	//	18,
	//}
	//fmt.Printf("p5=%#v", p5)
//}

// 5.结构体内存分布：结构体占用一块连续的内存
//type test struct{
//	a int8
//	b int8
//	c int8
//	d int8
//}
//
//func main() {
//	n := test{
//		1, 2, 3, 4,
//	}
//	fmt.Printf("n.a %p \n", &n.a)
//	fmt.Printf("n.b %p \n", &n.b)
//	fmt.Printf("n.c %p \n", &n.c)
//	fmt.Printf("n.d %p \n", &n.d)
//}

// 5.1.空结构体是不占用空间的
//var v struct{}
//
//func main() {
//	println(unsafe.Sizeof(v))
//}

// 6.结构体的面试题
type student struct {
	name string
	age int
}

func main() {
	m := make(map[string]*student)

	stus := []student{
		{"小王子", 18},
		{"娜扎", 23},
		{"大王八", 9000},
	}

	for _, stu := range stus {
		fmt.Println(stu, &stu)

		m[stu.name] = &stu
		fmt.Println(m)
	}
	fmt.Println(m, "--------")


	for k, v := range m {
		fmt.Println(k, "=>", v.name, v.age)
	}
}