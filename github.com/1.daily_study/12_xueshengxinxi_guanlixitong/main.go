package main

import (
	"fmt"
	"os"
)

// 学生管理系统的入口

func welcome() {
	//	优先输出的信息
	fmt.Println("欢迎来到学生信息管理系统：")
	fmt.Println("1.添加学生")
	fmt.Println("2.修改学生信息")
	fmt.Println("3.展示学生信息")
	fmt.Println("4.退出系统")
}

func main() {
	welcome()
	// 1.获取用户的信息输入
	var input int
	for {
		fmt.Printf("请选择你的输入：")
		fmt.Scanln(&input)
		switch input {
		case 1:
			// 添加学生
			AddStudent()
		case 2:
			// 更改学生信息
			UpdateStudent()
		case 3:
			// 展示学生信息
			ShowStudent()
		case 4:
			// 退出系统
			os.Exit(0)
		default:
			fmt.Println("无效的输入")
			os.Exit(0)
		}
	}
}
