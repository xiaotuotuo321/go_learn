package main

import "fmt"

// 练习题： 数据数组中只出现一次的数字

func main() {
	nums := [...]int{1, 2, 3, 4, 3, 2, 1}

	n := nums[0]

	for i := 1; i < len(nums); i++ {
		n ^= nums[i]
		fmt.Printf("nums[i]的值为%d \n", nums[i])
		fmt.Printf("n的值为%d n", n)
	}

	fmt.Println(n)
}
