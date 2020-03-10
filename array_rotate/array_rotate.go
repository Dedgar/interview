package main

import "fmt"

func rotate(nums []int, k int) {
	aLen := len(nums)
	if aLen <= 1 {
		fmt.Println("aLen is short")
		fmt.Println(nums)
		return
	}
	mod := aLen - (k % aLen)
	fmt.Println("aLen, k, mod", aLen, k, mod)

	begin := nums[mod:]
	end := nums[:mod]

	begin = append(begin, end...)

	copy(nums, begin)

	fmt.Println(nums)
}

func main() {
	rotate([]int{1}, 1)
	rotate([]int{1, 2}, 1)
	rotate([]int{1, 2}, 5)
	rotate([]int{1, 2, 3, 4, 5}, 1)
	rotate([]int{1, 2, 3, 4, 5}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
