package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// empty struct for size 0 vs bool 1
	// match length of input for preallocation
	dupes := make(map[int]struct{}, len(nums))

	// exit immediately if we already have the number in our dupes array
	for _, i := range nums {
		if _, ok := dupes[i]; ok {
			return true
		}
		dupes[i] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 5, 87, 3, 7, 2, 56, 87, 43, 87, 4, 5, 8}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1}))
	fmt.Println(containsDuplicate([]int{0, 0, 0, 0, 0}))
}
