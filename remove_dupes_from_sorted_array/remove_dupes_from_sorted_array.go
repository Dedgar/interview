package main

import "fmt"

func checkSlice(n *[]int, i, j int) {
	// So we don't get an out of bounds error,
	// check that there even is an i+1 first.
	if i+1 < len(*n) {
		if j == (*n)[i+1] {
			// The new slice = everything up until this point +
			// everything +1 position from this point.
			*n = append((*n)[:i], (*n)[i+1:]...)
			checkSlice(n, i, j)
		}
	}
}

func removeDuplicates(nums []int) int {
	for i, j := range nums {
		fmt.Println("on item", i, j)
		checkSlice(&nums, i, j)
	}
	return len(nums)
}

func main() {
	//fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
