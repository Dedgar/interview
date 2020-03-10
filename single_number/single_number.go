package main

import "fmt"

func singleNumberXOR(nums []int) int {
	r := 0
	for _, v := range nums {
		r = r ^ v
	}
	return r
}

func singleNumberMap(nums []int) int {
	odd := make(map[int]struct{}, len(nums))
	var single int

	// if we've seen the number before, remove it from the map
	for _, n := range nums {
		if _, ok := odd[n]; ok {
			delete(odd, n)
			continue
		}
		odd[n] = struct{}{}
	}

	// loop through the map that was promised to only have 1 element
	for k := range odd {
		single = k
	}
	return single
}

func main() {
	fmt.Println(singleNumberXOR([]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 6}))
}
