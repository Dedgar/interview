package main

import "fmt"

func rotate(nums []int, k int) {
	aLen := len(nums)
	if aLen <= 1 {
		fmt.Println("aLen is short")
		fmt.Println(nums)
		return
	}

	// get the modulus since we only need to move by the remainder.
	// it saves us from having to do math around where the numbers should be
	// after shifting an array of items a hundred times.
	mod := aLen - (k % aLen)
	fmt.Println("aLen, k, mod", aLen, k, mod)

	// start making new arrays with everything before the shift number
	// and everything after the shift number
	begin := nums[mod:]
	end := nums[:mod]

	// make a new array with the items shifted by the correct amount
	begin = append(begin, end...)

	// copy because of the weird challenge restrictions, or else we could just
	// return the begin var at this point.
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
