package main

// log包实现了简单的日志服务。

// 1.使用logger
/*
log包定义了logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的"标准"logger,可以通过调用Print，fatal和panic
*/

//func main() {
//	log.Println("普通日志")
//	v := "很普通的"
//	log.Printf("这是一条%s日志\n", v)
//	log.Println("这是一条会触发fatal的日志")
//	log.Println("这是一条会触发panic的日志")
//}

// 2.配置logger 默认情况下logger只会提供日志的时间信息，如果想获得日志的文件名和行号等。log标准库中的Flags函数会返回标准logger的输出配置
// 	而SetFlags函数是用来设置标准logger的输出配置
// 2.1.flag选项
/*
log 标准库提供了如下的flag选项，他们是一系列定义好的常量
const(
	// 控制日志信息的细节，但是不能控制输出的顺序和格式
	// 输出的日志在每一项后会有一个冒号分隔
)
*/
//const (
//	Ldate = 1<<iota		// 日期：2009/01/23
//	Ltime				// 时间：01：23：23
//	Lmicroseconds		// 微妙级别的时间：01：23：23.123123
//	Llongfile			// 文件全路径名+行号
//	Lshortfile			// 文件名+行号
//	LUTC				// 使用UTC时间
//	LstdFlags = Ldate | Ltime	// 标准logger的初始值
//)
//
//func main() {
//	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
//	log.Println("这是一条日志")
//}

// 2.2.配置日志的前缀：Prefix函数用来查看logger的输出前缀，SetPrefix函数用来设置输出前缀
//func main() {
//	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
//	log.Println("这是一条很普通的日志")
//	log.SetPrefix("[小王子]")
//	log.Println("这是一条很普通的日志")
//}

// 2.3.配置日志的输出位置：SetOutput函数用来设置标准logger的输出目的地，默认的是标准错误输出
//func main() {
//	logFile, err := os.OpenFile("./xxx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil{
//		fmt.Println("open log file faild, err:", err)
//	}
//	log.SetOutput(logFile)
//	log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
//	log.Println("这是一条很普通的日志")
//	log.SetPrefix("[小王子]")
//	log.Println("这是一条很普通的日志")
//}

// ** 如果使用标准的logger，一般会把上面的配置写到init函数中

// 2.4.创建logger：标准库中还提供了一个新建logger对象的构造函数-New,支持创建自己的logger示例
// func New(out io.Writer, prefix string, flag int) *Logger
//func main() {
//	logger := log.New(os.Stdout, "<New>", log.Llongfile|log.Lmicroseconds|log.Ldate)
//	logger.Println("这是自定义的logger记录的日志")
//}

// 备注：go内置的log库功能有限，无法满足不同级别日志的情况，可以使用第三方的日志库：logrus,zap等


