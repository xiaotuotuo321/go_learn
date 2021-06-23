package main

import "fmt"

// 构造学生结构体和方法

type Student struct {
	id int
	name, class string
}

func NewStudent(id int, name, class string) *Student {
	return &Student{
		id: id,
		name: name,
		class: class,
	}
}

func InputStudentInfo() (int, *Student) {
	// 用户输入学生信息
	var (
		id int
		name, class string
	)
	fmt.Println("请输入学生ID：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入学生班级：")
	fmt.Scanln(&class)

	stu := NewStudent(id, name, class)

	return id, stu
}

var ClassInfo = make([]*Student, 0)

func AddStudent() {
	// 添加学生
	_, stu := InputStudentInfo()
	ClassInfo = append(ClassInfo, stu)
	return
}

func UpdateStudent() {
	// 更新学生的信息
	id, stu := InputStudentInfo()
	for _, v := range ClassInfo{
		if v.id == id {
			v.id = stu.id
			v.name = stu.name
			v.class = stu.class
		}
	}
	return
}

func ShowStudent() {
	// 展示学生信息
	for _, v := range ClassInfo{
		id := v.id
		name := v.name
		class := v.class

		fmt.Printf("学生：%d， 姓名：%s，班级：%s \n", id, name, class)
	}
}