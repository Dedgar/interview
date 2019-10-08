package main

import "fmt"

func checkSlice(n *[]int, i, val int) {
	fmt.Println("looks like", *n)
	if len(*n) == 0 {
		return
	}

	if i < len(*n) {
		if (*n)[i] == val {
			fmt.Println("match, gotta remove", (*n)[i])
			// The new slice = everything up until this point +
			// everything +1 position from this point.
			if len(*n) == 1 {
				*n = []int{}
				return
			}

			if i+1 < len(*n) {
				*n = append((*n)[:i], (*n)[i+1:]...)
				checkSlice(n, i, val)
			} else if i == len(*n)-1 {
				*n = (*n)[:i]
			}
		}
	}
}

func removeElement(nums []int, val int) int {
	for i := range nums {
		checkSlice(&nums, i, val)
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 0, 0, 0}, 0))
}
