package main

import (
	"fmt"
	"time"
)

//func main() {
//	now := time.Now()
//	fmt.Printf("现在时间：%v\n", now)
//
//	year := now.Year()
//	month := now.Month()
//	day := now.Day()
//	hour := now.Hour()
//	minute := now.Minute()
//	second := now.Second()
//
//	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)
//}

// 2.时间戳
// 2.1.根据时间对象获取时间戳
//func main() {
//	now := time.Now()
//
//	t1 := now.Unix()
//	t2 := now.UnixNano()
//
//	fmt.Printf("秒级时间戳：%d\n", t1)
//	fmt.Printf("纳秒级时间戳：%d\n", t2)
//}

// 2.2.将函数转换为时间格式
//func main() {
//	now := time.Now().UnixNano()
//
//	timeObj := time.Unix(now / 1e9, 0)	// 将时间戳转为时间格式
//	fmt.Println(timeObj)
//
//	year := timeObj.Year()
//	fmt.Println(year)
//}

// 3.时间运算
// 3.1.时间相加
//func main() {
//	now := time.Now()
//	later := time.Now()
//	duration := later.Sub(now)
//	fmt.Println(duration)
//}

// 4.定时器： 定时器本质上是一个通道
//func main() {
//	ticker := time.Tick(time.Second)
//	for i := range ticker{
//		fmt.Println(i)
//	}
//}

// 5.时间格式化：时间类型有一个自带的Format进行格式化: 格式化的方式是go的诞生时间：2006-01-02 15:04:05
//func main() {
//	now := time.Now()
//	// 24小时制
//	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))
//	// 12小时制
//	fmt.Println(now.Format("2006-01-02 03:04:05 Mon Jan"))
//
//	fmt.Println(now.Format("2006/01/02"))
//}

// 6.解析字符串格式的时间
func main() {
	now := time.Now()
	fmt.Println(now)

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
	}

	// 按照指定的时区和指定的格式来解析时间
	timeStr, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-05-27 10:43:30", loc)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(timeStr)
	fmt.Printf("%T\n", timeStr)
}


