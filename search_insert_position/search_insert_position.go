package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l := len(nums)

	if target > nums[l-1] {
		return l
	}

	for i, j := range nums {
		if j == target {
			return i
		} else if j > target {
			if i > 0 {
				return i
			}
			return 0
		}
	}
	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{2, 3, 4, 5}, 1))
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, 50))
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, -50))
	fmt.Println(searchInsert([]int{-1, -2, -3, -4, -5, 53}, -50))
	fmt.Println(searchInsert([]int{-1, -2, -3, -4, -5}, 50))
}
