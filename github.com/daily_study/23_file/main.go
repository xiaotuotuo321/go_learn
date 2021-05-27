package main

import (
	"fmt"
	"os"
)

// go语言文件操作：文件是存储在外部介质（通常是磁盘）上的数据集合，文件分为文本文件和二进制文件

// 1.打开和关闭文件
// os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件示例调用close()方法能够关闭文件
//func main() {
//	// 只读的方式打开当前目录下的main.go文件
//	file, err := os.Open("./main.go")
//	if err != nil{
//		fmt.Println("open file failed! err:", err)
//	}
//
//    // 关闭文件
//    file.Close()
//}
// 关闭文件时通常用defer注册文件关闭语句

// 2.读取文件
// func(f *File) Read(b []byte) (n int, err error)
// 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾会返回0和io.EOF
//func main() {
	// 只读方式打开当前目录下的main.go文件
	//file, err := os.Open("/Users/whp/go/src/go_learn/github.com/daily_study/23_file/main.go")
	//if err != nil{
	//	fmt.Println("open file failed! err:", err)
	//}
	//defer file.Close()
    // 使用read方法读取数据
    //var tmp = make([]byte, 128)
    //n, err := file.Read(tmp)
    //if err == io.EOF{
    //	fmt.Println("文件读完了")
	//}
	//if err != nil{
	//	fmt.Println("read file failed, err:", err)
	//}
	//fmt.Printf("读取了%d字节数据\n", n)
    //fmt.Println(string(tmp[:n]))

    // 循环读取文件
    //var content []byte
	//var tmp = make([]byte, 128)
	//for {
	//	n, err := file.Read(tmp)
	//	if err == io.EOF{
	//		fmt.Println("文件读完了")
	//		break
	//	}
	//	if err != nil{
	//		fmt.Println("read file failed, err:", err)
	//	}
	//	content = append(content, tmp[:n]...)
	//}
	//fmt.Println(string(content))

	// bufio是在file的基础上封装了一层API，支持更对的功能
	//reader := bufio.NewReader(file)
	//for{
	//	line, err := reader.ReadString('\n')  	// 字符
	//	if err == io.EOF{
	//		if len(line) != 0{
	//			fmt.Println(file)
	//		}
	//		fmt.Println("文件读完了")
	//		break
	//	}
	//	if err != nil{
	//		fmt.Println("read file failed, err:", err)
	//	}
	//	fmt.Println(line)
	//}

	// ioutil.ReadFile读取整个文件
//	content, err := ioutil.ReadFile("/Users/whp/go/src/go_learn/github.com/daily_study/23_file/main.go")
//	if err != nil{
//		fmt.Println("read file failed, err:", err)
//	}
//	fmt.Println(string(content))
//}

// 2.写入操作 os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能
// func OpenFile(name string, flag int, perm FildMode) (*File, error)

/*
打开文件的集中模式
模式					含义
os.O_WRONLY			创建文件
os.O_CREATE			只读
os.O_RDONLY			只写
os.O_TRUNC			清空
os.O_APPEND    		追加

perm: 文件追加，一个八进制数，r(读)04，w(写)02，x(执行)01
*/

// 2.1.write和writeString
func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println("open file failed, err:", err)
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))		// 写入字节切片数据
	file.WriteString("hello 小王子")
}

