package main

import "fmt"

// 介绍array

// 数组声明时就确定，可以修改数组成员，但是数组的大小不可变化

// 语法：var 数据变量名 [元素数组] T

// 1.数组初始化的几种方式
// 1.1.方法一：初始化数组时可以使用初始化列表来设置数组元素的值
// func main() {
// 	var testArray [3]int
// 	var numArray = [3]int{1, 2}
// 	var cityArray = [3]string{"北京", "上海", "深圳"}

// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Println(cityArray)
// }

// 1.2.方法二：确保提供的初始值和数组长度一致，可以让编译器自动推断数组的长度
// func main() {
// 	var testArray [3]int
// 	var numArray = [...]int{1, 2}
// 	var cityArray = [...]string{"北京", "上海", "深圳"}

// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Printf("type of numArray: %T \n", numArray)
// 	fmt.Println(cityArray)
// 	fmt.Printf("type of cityArray: %T \n", cityArray)
// }

// 1.3.方法三：可以指定索引值的方式来初始化数组，
//func main() {
//	a := [...]int{1: 10, 10: 1}
//
//	fmt.Println(a)
//	fmt.Printf("type of a: %T \n", a)
//}

// 2.数组的遍历：
// func main() {
// 	var a = [...]string{"北京", "上海", "深圳"}
// 	// 方法1：for循环遍历
// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(a[i])
// 	}

// 	// 方法2：for range遍历
// 	for index, value := range a {
// 		fmt.Println(index, value)
// 	}
// }

// 3.多维数组：go语言支持多维数组；二维数组的定义和遍历

// func main() {
// 	a := [3][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}

// 	fmt.Println(a)
// 	fmt.Println(a[2][1])

// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s \t", v2)
// 		}
// 		fmt.Println()
// 	}
// }

// 3.1.多维数组只有第一层可以使用...来让编译器推导数组长度
// 支持的写法
// a := [...][2]string{
// 	{"北京", "上海"},
// 	{"广州", "深圳"},
// 	{"成都", "重庆"},
// }

// 不支持的写法
// b := [3][...]string{
// 	{"北京", "上海"},
// 	{"广州", "深圳"},
// 	{"成都", "重庆"},
// }

// 4.数组是值类型 数组是值类型；赋值和传参会复制整个数组。一次改变副本的值，不会改变本身的值
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a) // 在modify中修改的是a的副本x
	fmt.Println(a)
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b)
	fmt.Println(b)
}

// 练习题1：求数组[1, 3, 5, 7, 8]所有元素的和

// func main() {
// 	var a = [...]int{1, 3, 5, 7, 8}
// 	count := 0

// 	for _, val := range a {
// 		count += val
// 	}

// 	fmt.Println(count)
// }

// 练习题2：找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

//func main() {
//	var a = [...]int{1, 3, 5, 7, 8}
//	for i, val := range a {
//		for j := i; j < len(a); j++ {
//			if val+a[j] == 8 {
//				fmt.Printf("(%d, %d) \n", i, j)
//			}
//
//		}
//	}
//}

/*
总结数组的特定：
1.数组的长度是数组类型的一部分，数组的成员变量可以修改，但是数组的长度不能改变。
2.数组是值类型，赋值和传参会复制整个数组。改变副本的值，不会改变本身的值。
*/
