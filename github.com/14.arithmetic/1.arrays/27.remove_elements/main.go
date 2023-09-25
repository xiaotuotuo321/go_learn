package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0

	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(removeElement(nums, val))
}
