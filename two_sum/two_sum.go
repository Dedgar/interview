package main

import "fmt"

// Solving this somewhat backwards with subtraction instead of addition.
// The position map's key is a number from the nums array, its value is its index.
// We'll keep checking the position map since its a quick constant time operation.
// If there's a match in the map for [target - num], that means that (num+that num) = target.
// In that case, we return a new int slice containing our position and the value returned from the map.
// If there's no match this time, store the num as a key in posMap, with its position in nums as its value.
func twoSum(nums []int, target int) []int {
	fmt.Println("Starting with array nums:", nums)
	posMap := make(map[int]int)

	for pos, num := range nums {
		fmt.Println("target minus our current num would be:", target, num, target-num)
		if index, ok := posMap[target-num]; ok {
			return []int{pos, index}
		}

		posMap[num] = pos
		fmt.Println("indexmap is now:", posMap)
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
