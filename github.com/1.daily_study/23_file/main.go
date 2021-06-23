package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	//file, err := os.Open("/Users/whp/go/src/go_learn/github.com/1.daily_study/23_file/main.go")
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
//	content, err := ioutil.ReadFile("/Users/whp/go/src/go_learn/github.com/1.daily_study/23_file/main.go")
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
//func main() {
//	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
//	if err != nil{
//		fmt.Println("open file failed, err:", err)
//	}
//	defer file.Close()
//	str := "hello 沙河"
//	file.Write([]byte(str))		// 写入字节切片数据
//	file.WriteString("hello 小王子")	// 直接写入字符串数据
//}

// 2.2.bufio.NewWriter
//func main() {
//	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
//	if err != nil{
//		fmt.Println("open file failed, err:", err)
//		return
//	}
//
//	defer file.Close()
//	writer := bufio.NewWriter(file)
//	for i := 0; i < 10; i++ {
//		writer.WriteString("Hello 小红\n")		// 将数据先写入缓存
//		if i == 5{
//			break
//		}
//	}
//	writer.Flush()	// 将缓存中的内容写入文件
//}

// 2.3.ioutil.WriteFile  直接生成一个新文件并写入内容
//func main() {
//	str := "hello 沙河"
//	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
//	if err != nil{
//		fmt.Println("write file failed, err:", err)
//		return
//	}
//}

// 3.练习
// 3.1.借助io.Copy()实现一个拷贝文件函数
// 拷贝文件函数
//func CopyFile(dstName, srcName string) (written int64, err error){
//	// 以读方式打开源文件
//	src, err := os.Open(srcName)
//	if err != nil{
//		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
//		return
//	}
//	defer src.Close()
//	// 以写|创建的方式打开目标文件
//	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
//	if err != nil{
//		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
//		return
//	}
//	defer dst.Close()
//	return io.Copy(dst, src)	// 调用io.Copy()拷贝内容
//}
//
//func main() {
//	_, err := CopyFile("dst.txt", "src.txt")
//	if err != nil{
//		fmt.Println("copy file failed, err:", err)
//		return
//	}
//	fmt.Println("copy done!")
//}

// 3.2.实现cat命令
// 使用文件操作相关知识，模拟实现linux平台cat命令的功能
// cat命令实现
func cat(r *bufio.Reader){
	for{
		buf, err := r.ReadBytes('\n')	// 字符
		if err == io.EOF{
			// 退出之前将读取到的内容输出
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse()	// 解析命令行参数
	if flag.NArg() == 0{
		// 如果没有参数默认从标准输入读取到内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i:=0; i < flag.NArg(); i++{
		f, err := os.Open(flag.Arg(i))
		if err != nil{
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}