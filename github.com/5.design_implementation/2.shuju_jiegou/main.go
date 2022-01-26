package main

import "fmt"

//  两种初始化方式
//func main() {
//	arr1 := [3]int{1, 2, 3}
//	arr2 := [...]int{4, 5, 6}
//	fmt.Println(arr1, arr2)
//}

// 在不考虑逃逸分析的情况下，数组中的个数小于或等于4个所有的变量会直接在栈上初始化，如果大于四个变量就会在静态存储区初始化后然后拷贝到栈上

// 切片
// 1. 初始化的三种方式
// 通过下标的方式获得数组或者切片的一部分； 使用字面量初始化新的切片； 使用关键字make创建切片
//func main() {
//	arr := [3]int{1, 2, 3}
//	slice := arr[0:1]
//	fmt.Println(slice)
//}

//func main() {
//	var vstat [3]int
//	vstat[0] = 1
//	vstat[1] = 2
//	vstat[2] = 3
//
//	var vauto *[3]int = new([3]int)
//	*vauto = vstat
//	slice := vauto[:]
//	fmt.Println(slice)
//}

// 2.切片的扩容：
// 期望容量大于当前容量的两倍就会使用期望容量；当切片的长度小于1024就会将容量翻倍；如果当前切片的长度大于1024就会增加每次增加25%的容量，直到新容量大于期望容量
// 申请内存的时候，还需要根据切片中的元素大小对其内存，当数组中元素所占的字节大小为1，8或者2的倍数时。

//func main() {
//	var arr []int64
//	arr = append(arr, 1, 2, 3, 4, 5)
//
//	fmt.Println(cap(arr), len(arr))
//}

// 3.切片的拷贝 整块拷贝会占用非常多的资源，在大切片上进行拷贝时一定要注意对性能的影响

// 哈希结构
// 为了解决哈希碰撞问题：最常用的是开放寻址法和拉链法
// 1.开放寻址法：依次探测和比较数组中的元素以判断目标键值对是否存在于哈希表中。 装载率超过70%之后。性能急剧下降
// 2.拉链法：先找桶 再从桶中取出键值对。装载因子 ：= 元素数量 ➗ 桶数量。一般情况下装载因子不会超过1

// go语言运行时同时使用了多个数据结构组合表示哈希表，其中runtime.hmap是最核心的结构体
//type hmap struct {
//	count     int		当前哈希表中的元素数量
//	flags     uint8
//	B         uint8		表示当前哈希表中的buckets的数量，桶的数量都是2的倍数。该字段会存储对数。len(buckets) == 2^B
//	noverflow uint16
//	hash0     uint32	哈希的种子，他能为哈希函数的结果引入随机性。在创建哈希表的时确定，调用哈希函数时作为参数传入
//
//	buckets    unsafe.Pointer
//	oldbuckets unsafe.Pointer	在扩容时保存之前buckets的字段，它的大小是当前buckets的一半
//	nevacuate  uintptr
//
//	extra *mapextra
//}
//
//type mapextra struct {
//	overflow    *[]*bmap
//	oldoverflow *[]*bmap
//	nextOverflow *bmap
//}

// 哈希的初始化：通过字面量和运行时。
func main() {
	dict := map[string]int{
		"小明": 100,
		"小红": 90,
		"小兰": 80,
	}

	score := dict["小明"]
	x := dict["小黑"]
	fmt.Println(score, x)
}