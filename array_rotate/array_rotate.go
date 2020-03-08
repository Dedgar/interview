package main

import "fmt"

func rotate(nums []int, k int) {
	rotated := make([]int, len(nums))
	aLen := len(nums) - 1
	diff := aLen - k + 1

	// if each array item index = array[index + shift]
	// if the item is over len(array): (item + k) - len(array) - 1
	switch {
	case aLen == 0:
		fmt.Println(nums)
		return
	case k == 0 || diff == 1:
		fmt.Println(nums)
		return
	}

	for i, v := range nums {
		switch {
		case i+k <= aLen:
			rotated[i+k] = v
			continue
		case i+k > aLen:
			rotated[i+k-aLen-1] = v
			continue
		}
	}
	nums = rotated
	fmt.Println(nums)
}

func main() {
	rotate([]int{1}, 1)
	rotate([]int{1, 2, 3, 4, 5}, 1)
	rotate([]int{1, 2, 3, 4, 5}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
