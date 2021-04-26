package main

import (
	"fmt"
	"reflect"
)

// 反射学习

// 1.变量的内在机制
// go语言中的变量是分为两个部分的：
	// 类型信息：预先定义好的元信息
	// 值信息：程序运行过程中可动态变化的

/* 2.反射介绍：反射是指在程序运行期对程序本身进行访问和修改的能力。程序编译的时候，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在程序
 运行时，程序无法获取自身的信息。
	支持反射的语言可以在程序编译期将变量的反射信息，比如字段名称，类型信息，结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以
在程序运行期获取尅性的反射信息，并且有能力修改他们。
	go程序在运行期使用reflect包访问程序的反射信息
*/

// 3.reflect包：在go语言的反射机制中，任何接口值都是由一个具体类型和具体类型的值两部分组成的。在go语言中反射的相关功能由内置的reflect包提供，任意
// 接口值在反射中都可以理解为reflect.Type和reflect.Value两部分组成。并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的value和Type

// 3.1.TypeOf：使用reflect.TypeOf()函数可以获得任意值得类型对象(reflect.Type)，程序通过类型对象可以访问任意值的类型信息
//func reflectOf (x interface{}) {
//	v := reflect.TypeOf(x)
//	fmt.Printf("%v \n", v)
//}
//
//func main() {
//	var a float64 = math.Pi
//	reflectOf(a)
//
//	var b int64 = 100
//	reflectOf(b)
//}

// 3.2.type name 和 type kind
/*
在反射中关于类型还划分为两种：类型(Type)和种类(Kind)。因为在go语言中我们使用type关键字构造很多自定义的类型，而种类(kind)就是指底层的类型，但在
反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类(kind)。
*/
//
//type myInt int64
//
//func reflectOf(x interface{}) {
//	v := reflect.TypeOf(x)
//	fmt.Printf("typeOf: %v, kindOf: %v\n", v.Name(), v.Kind())
//}
//
//func main() {
//	var a *float32 	// 指针
//	var b myInt		// 自定义类型
//	var c rune		// 类型别名
//
//	reflectOf(a)
//	reflectOf(b)
//	reflectOf(c)
//
//	type person struct {
//		Name string
//		Age int
//	}
//	type book struct{title string}
//
//	var d = person{
//		"哈哈",
//		15,
//	}
//	reflectOf(d)
//
//	var e = book{"神雕侠侣"}
//	reflectOf(e)
//}

// 注意：在go语言的反射中，像数组，切片，map，指针等类型的变量，他们的.Name都是返回空。

// 4.reflect.ValueOf 返回的是reflect.Value类型，其中包含原始值的信息，reflect.Value与原始值之间可以互相转换。
// 4.1.通过反射获取值
//func reflectValue(x interface{}) {
//	v := reflect.ValueOf(x)
//	k := v.Kind()
//
//	switch k {
//	case reflect.Int64:
//		fmt.Printf("type is int64, value is %v\n", int64(v.Int()))
//	case reflect.Float32:
//		fmt.Printf("type is float32, value is %v\n", float32(v.Float()))
//	case reflect.Float64:
//		fmt.Printf("type is float64, value is %v\n", float64(v.Float()))
//	}
//}
//
//func main() {
//	var a float32 = 10
//	var b float64 = 20
//
//	reflectValue(a)
//	reflectValue(b)
//
//	c := reflect.ValueOf(10)
//	fmt.Printf("type of c: %T\n", c)
//}

// 4.2.通过反射设置变量的值：想要在函数中通过反射修改变量的值，需要注意函数传递的是值拷贝，参数传递变量地址才能修改变量值。而反射中使用专有的Elam()
// 方法开获取指针对应的值

//func reflectSetValue1(x interface{}){
//	v := reflect.ValueOf(x)
//	if v.Kind() == reflect.Int64{
//		v.SetInt(200)
//	}
//}
//
//func reflectSetValue2(x interface{}){
//	v := reflect.ValueOf(x)
//	// 反射中使用 Elem()方法获取指针对应的值
//	if v.Elem().Kind() == reflect.Int64{
//		v.Elem().SetInt(200)
//	}
//}
//
//func main() {
//	var a int64 = 10
//	//reflectSetValue1(a)
//	reflectSetValue2(&a)
//	fmt.Println(a)
//}

// 4.3.isNil()和isValid()

// func (v Value) IsNil() bool
// isNil()报告v持有的值是否为nil。v持有的值得分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic

// func (v Value) IsValid() bool
// isValid()返回v是否持有一个值，如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic

// IsNil() 常用于判断指针是否为空；IsValid()常用于判定返回值是否有效
//func main() {
//	var a *int
//	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
//	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
//
//	b := struct {}{}
//	// 尝试用结构体中查找'abc'成员
//	fmt.Println("存在结构体成员'abc'：", reflect.ValueOf(b).FieldByName("abc").IsValid())
//	fmt.Println("存在结构体方法'abc'：", reflect.ValueOf(b).MethodByName("abc").IsValid())
//
//	c := map[string]int{}
//	fmt.Println("map中存在键'娜扎'：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
//}

// 5.结构体反射
// 5.1.与结构体相关的方法：任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象(reflect.Type)的NumField()
// 和Field()方法获得结构体成员的详细信息

// 5.2.structField类型 用来秒数结构体中的一个字段的信息
//type StructField struct {
//	Name string 	// Name 是字段的名字，PkgPath 是非导出字段的包路径，对导出字段该字段为""
//	PkgPath string
//	Type Type		// 字段的类型
//	Tag StructTag 	// 字段的标签
//	Offset uintptr	// 字段在结构体中的字节偏移量
//	Index []int 	// 用于Type.FieldByIndex时的索引切片
//	Anonymous bool 	// 是否匿名字段
//}

// 5.3.结构体反射示例：当使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名称去获取指定的字段信息
type student struct{
	Name string `json:"name"`
	Score int	`json:"score"`
}
//
//func main() {
//	stu1 := student{
//		Name: "小王子",
//		Score: 90,
//	}
//
//	t := reflect.TypeOf(stu1)
//	fmt.Println(t, t.Name(), t.Kind())
//
//	// 可以通过for循环遍历结构体中的所有字段信息
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		fmt.Printf("name:%s index: %d type: %v json tag:%v\n", field.Name, field.Index[0], field.Type, field.Tag.Get("json"))
//	}
//
//	// 通过字段名获取指定结构体字段信息
//	if scoreField, ok := t.FieldByName("Score"); ok{
//		fmt.Printf("name:%s index: %d type: %v json tag:%v\n", scoreField.Name, scoreField.Index[0], scoreField.Type, scoreField.Tag.Get("json"))
//	}
//}

// 写一个函数printMethod(s interface{})来遍历打印s包含的方法
func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())

	for i := 0; i < v.NumMethod(); i++{
		methodType := v.Method(i).Type()

		fmt.Printf("method name: %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", methodType)

		// 通过反射接口调用方法传递的参数必须是 []reflect.Value类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := student{
		Name: "小王子",
		Score: 100,
	}
	printMethod(stu1)
}

/*
反射是把双刃剑：
	反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码，但是反射不能被滥用，原因有三个
	1.基于反射的代码是脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能在代码写完很长时间以后了
	2.大量使用反射的代码通常难以理解
	3.反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个量级。
*/


