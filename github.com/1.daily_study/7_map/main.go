package main

import "fmt"

// go map 学习

// map是一种无序的基于key-value的数据结构，go语言中的map是引用类型，必须初始化才能使用

// 1. map的定义 map[KeyType]ValueType
// keytype 表示键的类型，valuetype表示键对应的值的类型
// map类型的变量默认初始值为nil,需要使用make()函数来分配内存
// make(map[KeyType]ValueType, [cap])
// cap表示map的容量,该参数虽然不是必须的，但在初始化map时,就为其指定一个合适的容量

// 2.map的基本使用
// func main() {
// 	scoreMap := make(map[string]int, 8)
// 	scoreMap["小明"] = 80
// 	scoreMap["小红"] = 65

// 	fmt.Println(scoreMap)
// 	fmt.Println(scoreMap["小明"])
// 	fmt.Printf("type of scoreMap: %T \n", scoreMap)
// }

// 2.1.在声明变量时，填充元素
// func main() {
// 	userInfo := map[string]string{
// 		"username": "沙河小王子",
// 		"password": "123456",
// 	}
// 	fmt.Println(userInfo)
// }

// 3.判断某个键是否存在  value, ok := map[key]
// func main() {
// 	scoreMap := make(map[string]int, 8)

// 	scoreMap["小明"] = 80
// 	scoreMap["小红"] = 90

// 	// 如果key存在ok为true，v为对应的值；不存在ok为false,v为值类型的零值
// 	v, ok := scoreMap["小红"]

// 	if ok {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("查无此人！")
// 	}
// }

// 4.map的遍历：遍历map时的元素顺序与添加键值对的顺序无关
// func main() {
// 	scoreMap := make(map[string]int, 8)
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	scoreMap["娜扎"] = 60

// 	for k := range scoreMap { //当前情况为只遍历key
// 		fmt.Println(k)
// 	}
// }

// 5.使用delete()函数删除键值对 delete(map, key)
// map:表示要删除键值对的map,key表示要删除的键值对的键
// func main() {
// 	scoreMap := make(map[string]int, 8)

// 	scoreMap["a"] = 1
// 	scoreMap["b"] = 2
// 	scoreMap["c"] = 3

// 	delete(scoreMap, "a")
// 	fmt.Println(scoreMap)
// }

// 6.按照指定顺序遍历map
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	var scoreMap = make(map[string]int, 200)
//
//	for i := 0; i < 100; i++ {
//		key := fmt.Sprintf("stu%02d", i)
//		value := rand.Intn(100)
//		scoreMap[key] = value
//	}
//	// 取出map中的所有key存入切片keys
//	var keys = make([]string, 0, 200)
//	for key := range scoreMap {
//		fmt.Printf("key的值为: %s, key的类型为: %T, %v, %+v\n", key, key, key, key)
//		keys = append(keys, key)
//	}
//	// 对切片进行排序
//	sort.Strings(keys)
//
//	// 按照排序之后的key遍历map
//	for _, key := range keys {
//		fmt.Println(key, scoreMap[key])
//	}
//}

// 7.元素为map类型的切片
//func main() {
//	var mapSlice = make([]map[string]string, 3)
//
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//
//	fmt.Println("after init")
//	// 对切片中的map元素进行初始化
//	mapSlice[0] = make(map[string]string, 10)
//	mapSlice[0]["name"] = "小王子"
//	mapSlice[0]["password"] = "123456"
//	mapSlice[0]["address"] = "沙河"
//
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//}

// 8.值为切片类型的map
//func main() {
//	var sliceMap = make(map[string][]string, 3)
//
//	fmt.Println(sliceMap)
//	fmt.Println("after init")
//
//	key := "中国"
//
//	value, ok := sliceMap[key]
//	if !ok {
//		value = make([]string, 0, 2)
//	}
//	value = append(value, "北京", "上海")
//	value = append(value, "深圳", "重庆")
//	sliceMap[key] = value
//	fmt.Println(sliceMap)
//}

// 练习题 1
// func main() {
// 	str := "how do you do"
// 	var strMap = make(map[string]int, 3)

// 	strArr := strings.Split(str, " ")
// 	fmt.Println(strArr)

// 	for _, value := range strArr {
// 		strMap[value] = strMap[value] + 1
// 	}

// 	fmt.Println(strMap)

// }

// 练习题 2 运行结果
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	fmt.Printf("s的类型为：%T\n", s)
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
