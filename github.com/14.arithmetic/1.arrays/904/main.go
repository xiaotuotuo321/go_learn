package main

import "fmt"

// 水果成篮
// https://leetcode.cn/problems/fruit-into-baskets/
/*
双指针 巧妙利用set来存储采过的树
*/

func totalFruit(fruits []int) int {
	res := 0
	cnt := map[int]int{}
	left := 0

	for right, value := range fruits {
		cnt[value]++

		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}

	return a
}

func main() {
	fruits := []int{1, 2, 3, 2, 2}

	fmt.Println(totalFruit(fruits))
}
