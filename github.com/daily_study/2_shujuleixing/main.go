package main

import "fmt"

// "math"

// 1.数字 数字字面量语法
// func main() {
// 	// 十进制
// 	var a int = 10
// 	fmt.Printf("%d \n", a)
// 	fmt.Printf("%b \n", a)

// 	// 八进制 以0开头
// 	var b = 077
// 	fmt.Printf("%o \n", b)

// 	// 十六进制 以0x开头
// 	var c int = 0xff
// 	fmt.Printf("%x \n", c)
// 	fmt.Printf("%X \n", c)
// }

// 2.浮点型

// func main() {
// 	fmt.Printf("%f \n", math.Pi)
// 	fmt.Printf("%.2f \n", math.Pi)
// }

// 3.复数: 复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
// var c1 complex64
// var c2 complex128

// func main() {

// 	c1 = 2 + 3i
// 	c2 = 5 - 4i

// 	fmt.Println(c1)
// 	fmt.Println(c2)
// }

// 4.布尔值，默认为false；不能将整型转化为布尔型；布尔型无法参加数值运算，也无法与其他类型进行转换

// 5.字符串  s1 := "hello"
// 5.1.字符串转义

// func main() {
// 	fmt.Println("str := \"c:\\Code\\lesson1\\go.ext\"")
// }

// 5.2.多行字符串和操作

// func main() {
// 	s1 := `第一行
// 	第二行
// 	第三行`
// 	s2 := `第四行
// 	第五行
// 	第六行`
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	fmt.Printf("s1+s2的长度为 %d", len(s1+s2))
// }

// 6.byte和rune类型
/*
go 语言的字符有一下两种：
	1. uint8类型,或者叫做byte类型，代表了ascii的一个字符
	2. rune类型，代表一个utf-8字符
*/
// func main() {
// 	s := "hello沙河"
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%v(%c) ", s[i], s[i])
// 	}
// 	fmt.Println()

// 	for _, r := range s {
// 		fmt.Printf("%v(%c) ", r, r)
// 	}
// }

// 6.1.修改字符串 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
// func changeString() {
// 	s1 := "big"
// 	// 强制类型转换
// 	byteS1 := []byte(s1)
// 	byteS1[0] = 'p'
// 	fmt.Println(string(byteS1))

// 	s2 := "白萝卜"
// 	runeS2 := []rune(s2)
// 	runeS2[0] = '红'
// 	fmt.Println(string(runeS2))
// }

// func main() {
// 	changeString()
// }

// 7.类型转换 go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候起作用；语法 T(表达式)

// func sqrtDemo() {
// 	var a, b = 3, 4
// 	var c int

// 	c = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// func main() {
// 	sqrtDemo()
// }

// 练习题1: 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。

// var (
// 	inta    int
// 	floata  float64
// 	boola   bool
// 	stringa string
// )

// func main() {
// 	inta = 10
// 	floata = 1.1
// 	boola = true
// 	stringa = "peace"

// 	fmt.Printf("%d 的类型为%T \n", inta, inta)
// 	fmt.Printf("%v 的类型为%T \n", floata, floata)
// 	fmt.Printf("%v 的类型为%T \n", boola, boola)
// 	fmt.Printf("%s 的类型为%T \n", stringa, stringa)
// }

// 练习题2：编写代码统计出字符串"hello沙河小王子"中汉字的数量。

// func main() {
// 	str := "hello沙河小王子"

// 	// temp := []rune(str)

// 	// var count int
// 	// for _, v := range temp {
// 	// 	if v > 256 {
// 	// 		count++
// 	// 		fmt.Println(string(v))
// 	// 	}
// 	// }
// 	count := 0
// 	for _, v := range str {
// 		if len(string(v)) == 3 {
// 			count++
// 		}
// 	}

// 	fmt.Printf("汉字的数量为%d \n", count)
// 	str2 := "a"
// 	fmt.Println("单个字母的string类型长度为", len(str2))
// 	str2 = "人"
// 	fmt.Println("单个汉字的string类型长度为", len(str2))
// }

// 练习题3： 乘法口诀表

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
