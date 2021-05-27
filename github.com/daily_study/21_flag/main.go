package main

import (
	"fmt"
	"os"
)

// flag 获取命令行参数，使得go编写命令行工具更方便

// 1.go原生的获取命令行参数的方法 os.Args []string是一个字符串类型的切片 它的第一个元素是执行文件的名称
func main() {
	if len(os.Args) > 0{
		for index, arg := range os.Args{
			fmt.Printf("args[%d] = %v\n", index, arg)
		}
	}
}

// 2.参数的声明  获取  解析
//func main() {
//	// 定义命令行的参数
//	var name string
//	var age int
//	var married bool
//	var delay time.Duration
//
//	flag.StringVar(&name, "name", "张三", "姓名")
//	flag.IntVar(&age, "age", 18, "年龄")
//	flag.BoolVar(&married, "married", false, "婚否")
//	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
//
//	// 解析命令行参数
//	flag.Parse()
//	fmt.Println(name, age, married, delay)
//	// 返回命令行参数后的其他参数
//	fmt.Println(flag.Args())
//	// 返回命令行参数后的其他参数的个数
//	fmt.Println(flag.NArg())
//	// 返回使用的命令行参数个数
//	fmt.Println(flag.NFlag())
//}


