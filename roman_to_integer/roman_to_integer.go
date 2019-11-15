package main

import "fmt"

var (
	nums = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	converted := 0
	l := len(s) - 1

	for i := 0; i < l; i++ {
		if nums[s[i]] < nums[s[i+1]] {
			converted -= nums[s[i]]
		} else {
			converted += nums[s[i]]
		}
	}
	return converted + nums[s[l]]
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("V"))
}
