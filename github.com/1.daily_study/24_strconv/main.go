package main

// go语言内置的strconv：实现了基本数据类型和其字符串表示的相互转换

// strconv包实现了基本数据类型与其字符串表示的转换，只要有以下常用的函数：Atoi(),Itia(),parse系列，format系列，append系列

// 1.string与int类型转换
// 1.1.Atoi 函数用于将字符串类型的整数转换为int类型
// func Atoi(s string) (i int, err error)
// 如果传入的字符串无法转换为int类型，就会返回错误
//func main() {
//	s1 := "100"
//	i1, err := strconv.Atoi(s1)
//	if err != nil{
//		fmt.Println("can't convert to int")
//	} else {
//		fmt.Printf("type:%T value:%#v\n", i1, i1) // type: int value:100
//	}
//}

// 1.2.Itoa()函数：Itoa()函数用于将int类型数据转换为对应的字符串表示， func Itoa(i int) string
//func main() {
//	i2 := 200
//	s2 := strconv.Itoa(i2)
//
//	fmt.Printf("type: %T value: %#v\n", s2, s2)
//}

// 2.parse系列函数： Parse类函数用于转换字符串为给定类型的值：ParseBool(),ParseFloat(),ParseInt(),ParseUint()
// 2.1.ParseBool() func ParseBool(str string) (value bool, err error)
// 返回字符串表示的bool值，它接收1，0，t，f，T，F，true，false，True，False，TRUE，FALSE；否则返回错误
//func main() {
//	a := "True"
//	b, _ := strconv.ParseBool(a)
//	fmt.Printf("type: %T, value: %v\n", b, b)
//}

// 2.2.ParseInt() func ParseInt(s string, base int, bitSize int)(i int64, err error)
//func main() {
//	a := "-0x100"
//	b, _ := strconv.ParseInt(a, 0, 0)
//	fmt.Printf("type: %T, value: %v\n", b, b)
//}
/*
返回字符串表示的整数值，接收正负号。
base 指定进制（2到26），如果base为0，则会从字符串前置判断；"0x"是16进制，"0"是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的证书类型，0，8，16，32，64分别代表int，int8， int16， int32， int64；
返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
*/

// 2.3.ParseUint 和ParseInt类似，但不接受正负号，用于无符号整型 func ParseUint(s string, base int, bitSize int) (n uint64, err error)

// 2.4.ParseFloat()
//func ParseFloat(s string, bitSize int) (f float64, err error)
/*
解析一个表示浮点类型的字符串并返回其值
如果字符串合乎语法，函数会返回最为接近s表示值得一个浮点数（使用IEEE754规范舍入）
bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32）， 64是float64
返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
*/
//func main() {
//	b, err := strconv.ParseBool("true")
//	f, err := strconv.ParseFloat("3.1415", 64)
//	i, err := strconv.ParseInt("-2", 10, 64)
//	u, err := strconv.ParseUint("2", 10, 64)
//	fmt.Println(err)
//	fmt.Println(b, f, i, u)
//}

// 3.Format系列函数： Format系列函数实现了将给定类型数据格式化为string类型数据的功能
// 3.1.FormatBool()
//func FormatBool(b bool) string  根据b的值返回"true"或"false"

// 3.2.FormatInt()
// func FormatInt(i int64, base int) string
// 返回i的base进制的字符串表示，base必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字

// 3.3.FormatUint()
// func FormatUint(i uint64, base int) string 无符号的FormatInt()

// 3.4.FormatFloat()
// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
/*
函数将浮点数表示为字符串并返回
bitSize表示f的来源类型（32：float32, 64: float64）,会据此进行四舍五入
fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
*/

//func main() {
//	s1 := strconv.FormatBool(true)
//	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
//	s3 := strconv.FormatInt(-2, 16)
//	s4 := strconv.FormatUint(2, 16)
//	fmt.Println(s1, s2, s3, s4)
//}

// 4.1. isPrint(r rune) bool 返回一个字符串是否可打印的，和unicode.IsPrint一样，r必须是：字母、数字、标点、符号、ASCII空格

// 4.2. CanBackQuote(s string) bool
// 返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串


// 除此之外，strconv包中还有Append系列、quote系列等函数


