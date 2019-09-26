package main

import (
  "fmt"
)

func twoSumLowMem(nums []int, target int) []int {
    var intSlice []int

    for index1, num := range nums {
        for index2, num2 := range nums {
            if index1 != index2 && (num + num2) == target {
                intSlice = append(intSlice, index1, index2)
                return intSlice
            }
        }
    }
    return intSlice
}

func twoSumLowLatency(nums []int, target int) []int {
    lookup := map[int]int{}

    for i, n := range nums {
        fmt.Printf("checking %v minus %v equals %v\n", target, n, target-n)
        if j, ok := lookup[target - n]; ok {
            return []int{j, i}
        }
        fmt.Printf("not found, making an index for %v at position %v\n", n, i)
        lookup[n] = i
    }
    return []int{}
}

func twoSumLowLat2(nums []int, target int) []int {
	n := len(nums)
	records := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if j, ok := records[target-nums[i]]; ok {
			return []int{j, i}
		}
		records[nums[i]] = i
	}
	return nil
}

func main() {
    //intSlice := []int{2,7,11,15}
    intSlice2 := []int{2,7,11,15,4}

    //target := 9
    target2 := 19

    //test := twoSumLowLatency(intSlice, target)
    test2 := twoSumLowLatency(intSlice2, target2)

    fmt.Println(test2)
}
