package main

// 本文介绍切片：切片是一个拥有相同类型元素的可变长度的序列。他是基于数组类型做的一层封装。灵活且支持自动扩容。

// 1.切片的定义

// 1.1.声明切片的语法：var name []T；name表示变量名，T表示切片中的元素类型
// func main() {
// 	// 声明切片类型
// 	var a []string              //声明一个字符串切片
// 	var b = []int{}             //声明一个整型切片并初始化
// 	var c = []bool{false, true} //声明一个布尔切片并初始化
// 	var d = []bool{false, true} //声明一个布尔切片并初始化

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(a == nil)
// 	fmt.Println(b == nil)
// 	fmt.Println(c == nil)
// 	// fmt.Println(c == d) 切片是引用类型，不支持直接比较，只能和nil比较
// }

// 1.2.切片的长度和容量
// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量

// 1.3.切片表达式：从字符串和数组指向数组或切片的指针构造子字符串或切片。两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high界限值外还指定容量的完整的形式。
// 1.3.1.简单切片表达式
// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}
// 	s := a[1:5]
// 	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s))
// }

// 1.3.2.越界问题
//func main() {
//	a := [...]int{1, 2, 3, 4, 5}
//	s := a[1:3]
//	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s))
//	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
//	fmt.Printf("s2:%v len(s2):%v cap(s2):%v \n", s2, len(s2), cap(s2))
//}

// 1.3.3.完成的切片表达式 a[low : high : max] 条件：0 <= low <= high <= max <= cap(a)
// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	t := a[1:3:5]
// 	fmt.Printf("t:%v len(t):%v cap(t):%v \n", t, len(t), cap(t))
// }

// 1.3.4.使用make()函数构造切片; 当需要动态的创建一个切片时，要使用make()函数: make([]T, size, cap); T：切片的元素类型，size：切片中元素的数量，cap：切片的容量
// func main() {
// 	a := make([]int, 2, 10)
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))
// }

// 1.3.5.检查切片是否为空，只能使用len(s) == 0来判断，而不是应该使用s == nil 来判断

// 练习题1：输出下面的代码
//func main() {
//	var a = make([]string, 5, 10)
//	fmt.Println(a)
//	fmt.Println(len(a))
//	for i := 0; i < 10; i++ {
//		a = append(a, fmt.Sprintf("%v", i))
//		fmt.Println(a)
//	}
//	fmt.Println(a) //[     0 1 2 3 4 5 6 7 8 9]
//}

// 练习题2：请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）。
//func main() {
//	var a = []int{3, 7, 8, 9, 1}
//	sort.Ints(a)
//
//	fmt.Println(a)
//}

//func main() {
//	a := make([]int, 0, 9)
//	for i := 1; i < 20; i++{
//		a = append(a, i)
//	}
//	fmt.Println(a, len(a), cap(a))
//}

/*
切片总结：
1.切片是相同元素可变长的序列，支持自动扩容
2.切片是引用类型，不能直接比较，只能和nil比较
3.切片的定义（声明，声明并初始化）
4.切片表达式，从字符串或数组中"切出来"，通过make构造
5.切片的本质是：对底层数组的封装，包含三个信息：底层数组的指针，切片的长度和切片的容量
*/
