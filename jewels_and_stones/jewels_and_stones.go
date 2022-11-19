package main

import "fmt"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Jewels and Stones.
Memory Usage: 2.1 MB, less than 58.72% of Go online submissions for Jewels and Stones.
*/
func numJewelsInStones(jewels string, stones string) int {
	c := 0
        // computing the len first increases runtime to (5 ms, faster than 9.88%)
        // r := make(map[rune]struct{}, len(jewels))
	r := map[rune]struct{}{}

	for _, j := range jewels {
		if _, ok := r[j]; !ok {
			r[j] = struct{}{}
		}
	}

	for _, s := range stones {
		if _, ok := r[s]; ok {
			c++
		}
	}
	return c
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Jewels and Stones.
Memory Usage: 1.9 MB, less than 95.93% of Go online submissions for Jewels and Stones.
*/
func numJewelsInStones2(jewels string, stones string) int {
	c := 0

	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				c++
			}
		}
	}

	return c
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
