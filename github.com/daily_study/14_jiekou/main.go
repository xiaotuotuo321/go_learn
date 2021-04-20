package main

import "fmt"

// 接口学习

// 1.接口类型：在go语言中接口是一种类型，一种抽象类型。interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是
// 定义一个协议（规则），interface是一种类型

// 2.为什么要使用接口
//type Cat struct {}
//
//func(c Cat) Say() string {
//	return "喵喵喵"
//}
//
//type Dog struct{}
//
//func (d Dog) Say() string {
//	return "汪汪汪"
//}
//
//func main() {
//	c := Cat{}
//	fmt.Println("猫", c.Say())
//	d := Dog{}
//	fmt.Println("狗", d.Say())
//}
// 接口区别于我们之前所有的数据类型，接口是一种抽象的类型。当看到一个接口类型的值时，唯一知道的是通过它的方法能做什么

// 3.接口的定义：go语言提倡面向接口编程；每个接口由数个方法组成，接口的定义如下：
/*
type 接口类型名 interface{
	方法名1（参数列表1）返回值列表1
	方法名2（参数列表2）返回值列表2
	...
}
接口名：使用type将接口定义为自定义的类型名。go语言的接口在命名时，一般会在单词后面添加er，如有操作的接口叫做Writer，有字符串功能的接口叫做Stringer
	等。接口名最好要能突出该接口的类型含义。
方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包之外的代码访问
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略

type writer interface {
	Write([]byte) error
}
*/

// 4.实现接口的条件：一个对象只要全部实现了接口中的方法，那么就实现了这个接口。接口就是一个需要实现的方法列表
//type Sayer interface {
//	say()
//}
//
//type Cat struct {}
//type Dog struct {}
//
//func (c Cat) say() {
//	fmt.Println("喵喵喵")
//}
//
//func (d Dog) say() {
//	fmt.Println("汪汪汪")
//}
//
//func main() {
//	var x Sayer
//	a := Cat{}
//	b := Dog{}
//	x = a
//	x.say()
//	x = b
//	x.say()
//}

// 5.接口类型变量：接口类型变量能够存储所有实现了该接口的实例。示例见 4

// 6.值接收者和指针接收者实现接口的区别
//type Mover interface {
//	move()
//}
//
//type dog struct {}
// 6.1.值接收者实现接口
//func (d dog) move() {
//	fmt.Println("狗会动")
//}
//// 此时实现的接口是dog类型
//func main() {
//	var x Mover
//	var wangcai = dog{}
//	x = wangcai
//	var fugui = &dog{}
//	x = fugui
//	x.move()
//}

// 从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给接口变量。

// 6.2.指针接收实现接口
//func (d *dog) move() {
//	var x Mover
//	var wangcai = dog{}
//	x = wangcai
//	var fugui = &dog{}
//	x = fugui
//}
// 此时Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai

// 7.面试题：我觉得不能；答对了
//type People interface {
//	Speak(string) string
//}
//
//type Student struct {}
//
//func (s *Student) Speak(think string) (talk string){
//	if think == "ls" {
//		talk = "你是ls"
//	} else {
//		talk = "您好"
//	}
//	return
//}
//
//func main() {
//	var peo People = &Student{}
//	think := "ll"
//	fmt.Println(peo.Speak(think))
//}

// 8.类型和接口的关系
// 8.1.一个类型实现多个接口：一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现
//type Sayer interface {
//	say()
//}
//
//type Mover interface {
//	move()
//}
//
//// dog 既可以实现Sayer接口，也可以实现Mover接口。
//type dog struct {
//	name string
//}
//
//func (d dog) say() {
//	fmt.Printf("%s会汪汪叫\n", d.name)
//}
//
//func (d dog) move() {
//	fmt.Printf("%s会动\n", d.name)
//}
//
//func main() {
//	var x Sayer
//	var y Mover
//	var a = dog{name: "旺财"}
//	x = a
//	y = a
//	x.say()
//	y.move()
//}

// 8.2.多个类型实现同一个接口
//type Mover interface {
//	move()
//}
//
//type dog struct {
//	name string
//}
//
//type car struct {
//	brand string
//}
//
//func (d dog) move() {
//	fmt.Printf("%s会动\n", d.name)
//}
//
//func (c car) move() {
//	fmt.Printf("%s速度70迈\n", c.brand)
//}
//
//func main() {
//	var x Mover
//	var a = dog{"旺财"}
//	var b = car{"保时捷"}
//
//	x = a
//	x.move()
//	x = b
//	x.move()
//}

// 8.3.一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
//type WashingMachine interface {
//	wash()
//	dry()
//}
//
//type dryer struct {}
//
//func (d dryer) dry(){
//	fmt.Println("甩一甩")
//}
//
//type haier struct {
//	dryer
//}
//
//func (h haier) wash() {
//	fmt.Println("洗一洗")
//}
//
//func main() {
//	var h = haier{}
//
//	h.wash()
//	h.dry()
//}

// 9.接口嵌套：接口与接口间可以通过嵌套创造出新的接口
//type Sayer interface {
//	say()
//}
//
//type Mover interface {
//	move()
//}
//
//type animal interface {
//	Sayer
//	Mover
//}
//
//type cat struct {
//	name string
//}
//
//func (c cat) say() {
//	println("喵喵喵")
//}
//
//func (c cat) move() {
//	fmt.Println("猫会动")
//}

//func main() {
//	var x animal
//	x = cat{"花花"}
//	x.say()
//	x.move()
//}

// 10.空接口：空接口是指没有定义任何方法的接口，因此任何类型都实现了空接口。空接口类型的变量可以存储任意类型的变量
//func main() {
//	// 定义一个空接口x
//	var x interface{}
//	s := "Hello 沙河"
//	x = s
//	fmt.Printf("type of s: %T, value: %v\n", x, x)
//	i := 100
//	x = i
//	fmt.Printf("type: %T, value: %v\n", x, x)
//	b := true
//	x = b
//	fmt.Printf("type: %T, value: %v\n", x, x)
//}

// 10.1.空接口的应用
// 10.1.1.空接口做为函数的参数：使用空接口实现可以接受任意类型的函数参数
//func show(a interface{}) {
//	fmt.Printf("type: %T, value: %v", a, a)
//}

// 10.1.2.空接口作为map的值:使用空接口实现可以保存任意值的字典
//func main() {
//	var studentInfo = make(map[string]interface{})
//	studentInfo["name"] = "沙河娜扎"
//	studentInfo["age"] = 18
//	studentInfo["married"] = false
//	fmt.Println(studentInfo)
//}

// 11.类型断言：接口值：一个接口的值是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
//var w io.Writer
//w = os.Stdout
//w = new(bytes.Buffer)
//w = nil

// 11.1.想要判断接口中的值这个时候可以使用类型断言：x.(T)
// x:表示类型为interface{}的变量；T：表示断言 x可能是的类型
// 该语法返回两个参数，第一个参数是x转化为T类型的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败
//func main() {
//	var x interface{}
//	x = "Hello 沙河"
//	v, ok := x.(string)
//	if ok{
//		fmt.Println(v)
//	}else {
//		fmt.Println("类型断然失败")
//	}
//}

// 11.2.上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现
func justityType(x interface{}){
	switch v:=x.(type) {
	case string:
		fmt.Println(v,"是一个字符串类型")
	case int:
		fmt.Println(v, "x是一个int类型")
	case bool:
		fmt.Println(v, "是一个bool类型")
	default:
		fmt.Println("unsupport type!")
	}
}

// 应为空接口可以存储任意值得特点，所以空接口在go语言中的使用十分广泛
// 关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。










