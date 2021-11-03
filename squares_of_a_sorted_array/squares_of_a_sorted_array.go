package main

import "fmt"

// simple abs implementation. returns positive ints
// we need this because negative numbers can return larger values than positive ones
// example: -4 ^ -4 = 16, which is greater than 3 ^ 3 = 9, despite 3 having a larger starting value than -4.
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// since the array is already sorted, this allows us to do something particularly important:
// seeing if the last number of the array squared will actually be larger than the first number after multiplying.
// the sorting gives us assurance that no "smaller" negative numbers come later in the list, since they would end up larger after multiplying.
func sortedSquares(nums []int) []int {
	numLen := len(nums)
	newNums := make([]int, numLen)

	bot := 0
	top := numLen - 1

	fmt.Printf("starting with %v of len %v \n", nums, numLen)
	for pos := numLen - 1; pos >= 0; pos-- {
		fmt.Printf("abs nums bot %v vs abs num top %v \n", abs(nums[bot]), abs(nums[top]))
		if abs(nums[bot]) > abs(nums[top]) {
			newNums[pos] = nums[bot] * nums[bot]
			bot++
		} else {
			newNums[pos] = nums[top] * nums[top]
			top--
		}
		fmt.Println(newNums)
	}

	return newNums
}

func main() {
	fmt.Println("***starting first***")
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println("***starting second***")
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
