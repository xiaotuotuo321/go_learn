package main

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
*/