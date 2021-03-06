package main

import "fmt"

// 方法和接收者：go中的方法是一种作用于特定类型变量的函数。这种特定类型叫做接收者
// 语法：func(接收者变量，接收者类型) 方法名(参数列表) (返回参数){ 函数体 }
// 接收者变量：官方建议使用接收者类型名称首字母的小写
// 接收者类型：接收者类型和参数类似；指针类型和非指针类型
// 防范名、参数列表、返回参数；

// 1.举例
//type Person struct {
//	name string
//	age int8
//}
//// 1.1.构造函数
//func NewPerson(name string, age int8) *Person{
//	return &Person{
//		name: name,
//		age: age,
//	}
//} // 1.2.Dream Person做梦的方法 func (p Person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
//}
//func main() {
//	p1 := NewPerson("小王子", 18)
//	p1.Dream()
//}

// 1.1.指针类型的接收者：指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后都是有效的。
//func (p *Person) SetAge(newAge int8){
//	p.age = newAge
//}
//
//func main() {
//	p1 := NewPerson("小王子", 18)
//	p1.SetAge(28)
//	fmt.Println(p1.age)
//}

// 1.2.值类型的接收值：当方法作用于值类型接收者时，go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改
// 操作知识针对副本，无法修改接收者变量本身
//func (p Person) SetAge2(newAge int8) {
//	p.age = newAge
//}
//
//func main() {
//	p1 := NewPerson("小王子", 18)
//	p1.SetAge2(28)
//	fmt.Println(p1.age)
//}

// 1.3.总结：什么时候用指针类型接收者
// 		1.需要修改接收者中的值
// 		2.接收者时拷贝代价比较大的大对象
//   	3.保证一致性，如果某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者


// 2.任意类型添加方法：接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。不能给别的包的类型定义方法
//type MyInt int
//
//func (i MyInt) SayHello(){
//	fmt.Println("Hello, I am a Int")
//}
//
//func main() {
//	var v1 MyInt
//	v1.SayHello()
//	v1 = 100
//	fmt.Printf("%#v %T\n", v1, v1)
//}

// 3.结构体的匿名字段：结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就是匿名字段
// 匿名字段的说法并不是说没有字段名，而是会采用类型名来做为字段名，结构体要求字段名称必须唯一，因为一个结构体中同种类型的匿名字段只有一个
//type Person struct{
//	string
//	int8
//}
//
//func main() {
//	p1 := Person{"小王子", 18}
//	fmt.Printf("%#v\n", p1)
//	fmt.Println(p1.string, p1.int8)
//}

// 4.嵌套结构体：一个结构体中可以嵌套的包含另一个结构体或者结构体指针；这个结构体中用到了匿名字段
//type Address struct {
//	Province, City string
//}
//type User struct {
//	Name, Gender string
//	Address
//}
//
//func main() {
//	user1 := User{
//		"小王子",
//		"男",
//		Address{
//			"河北",
//			"衡水",
//		},
//	}
//	fmt.Printf("%#v \n", user1)
//}

// 4.1.嵌套结构体中可能存在相同的字段名，为了避免歧义需要指定具体的内嵌套结构体字段
//type Address struct {
//	Province, City, CreateTime string
//}
//
//type Email struct {
//	Account, CreateTime string
//}
//
//type User struct {
//	Name, Gender string
//	Address
//	Email
//}
//
//func main() {
//	var user User
//	user.Name = "小王子"
//	user.Gender = "男"
//	user.Address = Address{
//		"河北",
//		"衡水",
//		"2021-03-04",
//	}
//	user.Email = Email{
//		"xiaowangzi@126.com",
//		"2023-03-04",
//	}
//	fmt.Printf("%#v \n", user)
//	fmt.Println(user.Address.CreateTime, user.Email.CreateTime)
//}

// 5.结构体的继承：结构体可以像其他语言中面向对象的集成
//type Animal struct {
//	name string
//}
//
//func (a *Animal)move(){
//	fmt.Printf("%s会动\n", a.name)
//}
//
//type Dog struct {
//	Feet int8
//	*Animal  // 通过嵌套匿名结构体实现继承
//}
//
//func (d *Dog) wang() {
//	fmt.Printf("%s会汪汪汪~\n", d.name)
//}
//
//func main() {
//	d1 := &Dog{
//		Feet: 4,
//		Animal: &Animal{
//			name: "大黄",
//		},
//	}
//	d1.wang()
//	d1.move()
//}

// 6.结构体字段的可见性：结构体中字段大写开头表示可公开访问，小写表示私有

// 7.结构体和json序列化
//type Student struct {
//	Id int
//	Gender string
//	Name string
//}
//
//type Class struct {
//	Title string
//	Students []*Student
//}
//
//func main() {
//	c := &Class{
//		Title: "火箭101",
//		Students: make([]*Student, 0, 30),
//	}
//
//	for i:=1; i < 10; i++ {
//		stu := &Student{
//			Name: fmt.Sprintf("stu%02d", i),
//			Gender: "男",
//			Id: i,
//		}
//		c.Students = append(c.Students, stu)
//	}
//
//	// json序列化, 结构体-->json格式的字符串
//	data, err := json.Marshal(c)
//	if err != nil{
//		fmt.Println("json marshal filed")
//		return
//	}
//	fmt.Printf("json:%s\n", data)
//
//	// 反序列化
//	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
//	c1 := &Class{}
//	err = json.Unmarshal([]byte(str), c1)
//	if err != nil {
//		fmt.Println("json unmarshal failed!")
//		return
//	}
//	fmt.Printf("%#v\n", c1)
//}

// 8.结构体标签（tag）：tag 是结构体的元信息，可以在运行时通过反射的机制读取出来，tag在结构体中字段的后方定义。
// `key1:"value1" key2:"value2"`
// 结构体tag由一个或多个键值对组成。键与值之间使用冒号分隔。使用双引号括起来。同一个结构体字段可以设置对个键值对tag,不同的键值对之间使用空格分隔
// 注意：为结构体编写tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行都不会提示任何的错误信息，
// 通过反射也无法获取正确取值。不要在key和value之间添加空格
//type Student struct {
//	ID int `json:"id"`	// 通过tag实现json序列化该字段时的key
//	Gender string 	// json序列化是默认使用字段名作为key
//	name string 	// 私有不能被json包访问
//}
//
//func main() {
//	s1 := Student{
//		1,
//		"男",
//		"沙河娜扎",
//	}
//	data, err := json.Marshal(s1)
//	if err != nil{
//		fmt.Println("json marshal failed!")
//		return
//	}
//	fmt.Printf("json str: %s\n", data)	// json str: {"id":1,"Gender":"男"}
//}

// 9.结构体和方法补充知识点：因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制他们时要特别注意。
type Person struct {
	name string
	age int8
	dreams []string
}

func (p *Person) SetDream(dreams []string) {
	p.dreams = dreams
}

func (p *Person) SetDream1(dreams []string){
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDream1(data)

	// 想要修改p1.dreams
	data[1] = "不睡觉"
	p1.dreams[1] = "不睡觉"
	fmt.Println(data) 	// [吃饭 不睡觉 打豆豆]
	fmt.Println(p1.dreams) 	// [吃饭 不睡觉 打豆豆]
}