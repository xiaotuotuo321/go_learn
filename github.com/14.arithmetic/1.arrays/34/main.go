package main

import "fmt"

// 34: 找出有序数组中目标值的第一个和最后一个值
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func searchRange(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)

	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}

	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}

	return []int{-1, -1}
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2

	for left <= right {
		mid := left + ((right - left) >> 2)
		if target <= nums[mid] {
			right = mid - 1
			border = right
		} else {
			left = mid + 1
		}
	}

	return border
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2

	for left <= right {
		mid := left + ((right - left) >> 2)
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
			border = left
		}
	}

	return border
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 5, 6}
	target := 3

	fmt.Println(searchRange(nums, target))
}
