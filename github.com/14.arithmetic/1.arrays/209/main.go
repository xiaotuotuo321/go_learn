package main

import "fmt"

// 209: 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
func minSubArrayLen(target int, nums []int) int {
	i := 0
	length := len(nums)
	sum := 0

	result := length + 1

	for j := 0; j < length; j++ {
		sum += nums[j]

		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}

			sum -= nums[i]
			i++
		}
	}

	if result == length+1 {
		return 0
	}
	return result
}

func main() {
	a := []int{2, 3, 1, 2, 4, 3}
	target := 7

	fmt.Println(minSubArrayLen(target, a))
}
