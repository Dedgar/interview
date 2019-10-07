package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	cA := make(map[int]int, len(arr))
	fM := make(map[int][]int, len(arr))

  // Increment a key in this map every time we see the key
	for _, i := range arr {
		cA[i]++
	}

	// Could potentially check the length of the int arrays here as an alternative.
	// pros: we drop out early if we find dupes.
	// cons: we check every time we run through the loop. for arrays with 999 1s and
	// only 1 2, we're checking 1000 times instead of 2 times during the next loop.
	for k, v := range cA {
		fM[v] = append(fM[v], k)
	}

  // Loop through the keys, which represent item counts.
  // If value's slice contains more than 1 item, it's not unique.
	for _, j := range fM {
		if len(j) > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
