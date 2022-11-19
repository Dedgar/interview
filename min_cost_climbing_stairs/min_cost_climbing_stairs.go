package main

import "fmt"

// 7ms faster than 40.38, 3.1mb memory beats 45.15%
func minCostClimbingStairs(cost []int) int {
	m := make([]int, len(cost))
	m[0], m[1] = cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		m[i] = min(m[i-1], m[i-2]) + cost[i]
	}

	return min(m[len(cost)-1], m[len(cost)-2])
}

func minCostClimbingStairsVerbose(cost []int) int {
	memo := make([]int, len(cost))
	memo[0], memo[1] = cost[0], cost[1]

	// start on index 2, compare both the previous number and the one before it, one step at a time.
	for i := 2; i < len(cost); i++ {
		fmt.Printf("cost and memo arrays: \n%v\n%v\n", cost, memo)
		fmt.Println("m[i]", memo[i])
		fmt.Println("m[i-1]", memo[i-1])
		fmt.Println("m[i-2]", memo[i-2])
		fmt.Println("cost[i]", cost[i])

		// compare the previous two items in memo, take whichever one is less, and add it to the current item in the cost array.
		memo[i] = min(memo[i-1], memo[i-2]) + cost[i]
		fmt.Printf("Added %v to %v to make %v\n", min(memo[i-1], memo[i-2]), cost[i], memo[i])
	}

	/* cost and memo arrays:
	[100 5 19 3]
	[100 5 0 0]
	m[i] 0
	m[i-1] 5
	m[i-2] 100
	cost[i] 19
	is 5 less than 100?
	is 5 less than 100?
	Added 5 to 19 to make 24
	cost and memo arrays:
	[100 5 19 3]
	[100 5 24 0]
	m[i] 0
	m[i-1] 24
	m[i-2] 5
	cost[i] 3
	is 24 less than 5?
	is 24 less than 5?
	Added 5 to 3 to make 8
	is 8 less than 24?
	returning the lowest combination of adjacent numbers, which is: 8
	*/

	// return the last item in our memo, which will be the lowest combination of numbers (even if it's only 0).
	// the length of our array matches that of cost, which has to be at least 2 as per lc constraints.
	return min(memo[len(cost)-1], memo[len(cost)-2])
}

func minCostClimbingStairsInt(cost []int) int {
	p1, p2 := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		p1, p2 = p2, min(p2+cost[i], p1+cost[i])
	}

	return min(p1, p2)
}

func min(a, b int) int {
	fmt.Printf("is %v less than %v?\n", a, b)
	if a < b {
		return a
	}
	return b
}

// Runtime: 13 ms, faster than 9.14%
// Memory Usage: 5 MB, less than 7.45%
func minCostClimbingStairsSlower(cost []int) int {
	memo = make(map[int]int)
	return min(cacher(cost, len(cost)-1), cacher(cost, len(cost)-2))
}

var memo map[int]int

func cacher(cost []int, i int) int {
	if i < 2 {
		return cost[i]
	}

	if _, ok := memo[i]; ok {
		return memo[i]
	}

	result := cost[i] + min(cacher(cost, i-1), cacher(cost, i-2))

	memo[i] = result

	return result
}

func main() {
	case1 := []int{10, 15, 20}
	case2 := []int{1, 2, 3, 1, 6, 7, 8, 3, 4, 5}
	case3 := []int{1, 50, 15, 9, 3}
	case4 := []int{1, 2, 16, 1, 6, 7, 8, 3, 4, 5}
	case5 := []int{100, 2, 1, 1, 6, 7, 8, 3, 4, 5}
	case6 := []int{100, 2, 1, 1, 1, 1, 1, 900, 4, 5}
	case7 := []int{100, 101, 5}
	case8 := []int{100, 5}
	case9 := []int{100, 5, 19, 3}
	fmt.Println(minCostClimbingStairs(case1))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case2))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case3))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case4))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case5))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case6))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case7))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case8))
	fmt.Println()
	fmt.Println(minCostClimbingStairs(case9))
}
