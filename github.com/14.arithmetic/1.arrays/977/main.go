package main

import "fmt"

// 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	res := make([]int, n)

	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]

		if lm > rm {
			res[k] = lm
			i++
		} else {
			res[k] = rm
			j--
		}
		k--
	}

	return res
}

func main() {
	a := []int{-4, -1, 0, 3, 10}
	for _, val := range sortedSquares(a) {
		fmt.Println(val)
	}
}
