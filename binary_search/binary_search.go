package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	numCount := len(nums)

	// Don't waste time iterating slices that exclusively contain bigger or smaller numbers than the target
	if nums[0] > target || nums[numCount-1] < target {
		return -1
	}

	bot := 0
	top := numCount - 1

	// identify middle of the slice by dividing in half
	// if the middle is still greater than our target, set the middle as the new top, and run again.
	// if the middle is still smaller than our target, set the middle as the new bottom, and run again.
	// when we've iterated through the whole slice this way, either we'll be left with the target, or break and return -1
	for bot <= top {
		mid := (bot + top) >> 1
		if nums[mid] > target {
			top = mid - 1
		} else if nums[mid] < target {
			bot = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
