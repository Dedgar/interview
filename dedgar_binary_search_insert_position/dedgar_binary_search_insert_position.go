package main

import "fmt"

func searchInsert(nums []int, target int) int {
	bot := 0
	top := len(nums) - 1

	// identify middle of the slice by dividing in half
	// if the middle is still greater than our target, set the middle as the new top, and run again.
	// if the middle is still smaller than our target, set the middle as the new bottom, and run again.
	// when we've iterated through the whole slice this way, we'll be left with the target
	for bot <= top {
		mid := (bot + top) >> 1
		fmt.Println("mid is: ", mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			top = mid - 1
			fmt.Println("top is now: ", top)
		} else {
			bot = mid + 1
			fmt.Println("bot is now: ", bot)
		}
	}
	return bot
}

func main() {
	fmt.Println("***starting first***")
	fmt.Println(searchInsert([]int{-1, 0, 3, 5, 9, 12}, 15))
	fmt.Println("***starting second***")
	fmt.Println(searchInsert([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println("***starting third***")
	fmt.Println(searchInsert([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println("***starting fourth***")
	fmt.Println(searchInsert([]int{1, 3, 5}, 3))

}
