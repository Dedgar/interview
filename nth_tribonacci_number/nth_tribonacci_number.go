package main

import "fmt"

func tribonacci(n int) int {
	if n < 3 {
		return n
//	} else if n == 2 {
//		return 1
//	}
	m := make([]int, n+1)
	m[0], m[1], m[2] = 0, 1, 1

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2] + m[i-3]
	}
	return m[n]
}

func main() {
	fmt.Println(tribonacci(25))
}
