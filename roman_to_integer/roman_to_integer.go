package main

import "fmt"

var (
	nums = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"L":    50,
		"C":    100,
		"D":    500,
		"M":    1000,
	}
)

func romanToInt(s string) int {
	converted := 0
	for _, n := range s {
		if i, ok := nums[n]; ok {
		}
		converted += i
	}
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
