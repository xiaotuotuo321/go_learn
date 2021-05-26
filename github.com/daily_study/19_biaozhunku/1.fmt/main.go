package main

// fmt 了解这个包

/*
fmt的作用：分为向外输出内容和获取输入内容
*/

// 1.向外输出
// 1.1.print系列函数会将内容输出到系统的标准输出，区别在于print函数直接输出内容，printf函数支持格式化的输出字符串，println函数会输出内容的结尾添加一个换行符

//func main() {
//	fmt.Print("在终端中打印信息。")
//	name := "沙河小王子"
//	fmt.Printf("我是：%s\n", name)
//	fmt.Println("在终端打印单独一行显示")
//}

// 1.2.Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
// 向标准输出写入内容
//func main() {
//	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
//	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil{
//		fmt.Println("打开文件出错，err:", err)
//	}
//	name := "小明"
//	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
//}

// 满足io.Writer接口的类型都支持写入

// 1.3.Sprint 系列函数会把传入的数据生成并返回一个字符串
//func main() {
//	s1 := fmt.Sprint("沙河小王子")
//	name := "沙河小王子"
//	age := 18
//	s2 := fmt.Sprintf("name:%s, age:%d", name, age)
//	s3 := fmt.Sprintln("沙河小王子")
//	fmt.Println(s1, s2, s3)
//}

// 1.4.Errorf: 根据format参数生成格式化字符串并返回一个包含该字符串的错误
//func main() {
//	err := fmt.Errorf("这是一个错误")
//	fmt.Println(err)
//	fmt.Printf("%T\n", err)
//}

// 2.格式化占位符
/*
通用占位符
占位符			说明
%v 				值得默认格式表示
%+v				类似%v,但输出结构体时会添加字段名
%#v				值得Go语法表示
%T				打印值得类型
%% 				百分号
*/

//func main() {
//	fmt.Printf("%v\n", 100)
//	fmt.Printf("%v\n", false)
//	o := struct {
//		name string
//	}{"小王子"}
//
//	fmt.Printf("%v\n", o)
//	fmt.Printf("%#v\n", o)
//	fmt.Printf("%T\n", o)
//	fmt.Printf("100%%\n")
//}

/*
布尔类型占位符
占位符		说明
%t			true或者false
*/

/*
整型
占位符			说明
%b				表示为二进制
%c				表示为对应的unicode码值
%d				表示为10进制
%o				表示为8进制
%x				表示为16进制，使用的是a-f
%X				表示为16进制，使用的是A-F
%U				表示为Unicode格式：U + 1234，等价于'U + %04X'
%q				该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
*/

//func main() {
//	n := 65
//	fmt.Printf("%b\n", n)
//	fmt.Printf("%c\n", n)
//	fmt.Printf("%d\n", n)
//	fmt.Printf("%o\n", n)
//	fmt.Printf("%x\n", n)
//	fmt.Printf("%X\n", n)
//	fmt.Printf("%U\n", n)
//	fmt.Printf("%q\n", n)
//}

/*
浮点数与复数
占位符				说明
%b					无小数部分，二进制指数的科学计数法，如-123456p-78
%e					科学计数法，-12345.456e+78
%E					科学计数法，-12345.456E+78
%f					有小数部分但无指数部分，123.456
%F					等价于%f
%g					根据实际情况采用%e或%f格式
%G					根据实际情况采用%E或%F格式
*/

//func main() {
//	f := 12.34
//	fmt.Printf("%b\n", f)
//	fmt.Printf("%e\n", f)
//	fmt.Printf("%E\n", f)
//	fmt.Printf("%f\n", f)
//	fmt.Printf("%F\n", f)
//	fmt.Printf("%g\n", f)
//	fmt.Printf("%G\n", f)
//}