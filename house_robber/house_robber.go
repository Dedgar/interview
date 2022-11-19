package main

import "fmt"

// faster than 100%, memory beats 74.35%
func rob(nums []int) int {
	m := make([]int, len(nums)+2)

	for i := 2; i < len(m); i++ {
		m[i] = max(m[i-2]+nums[i-2], m[i-1])
	}
	return m[len(m)-1]
}

func robVerbose(nums []int) int {
	// avoid having to make special logic for our edge case of 1-2 item arrays.
	memo := make([]int, len(nums)+2)

	// start at index 0 of nums if len(nums) >=2, and always start at index 2 of memo:
	//  |
	// [5 7 3]  |
	// [0 0 0 0 0]
	// if len(nums) is less than 2, we're returning index 0 of memo.
	// if len(nums) == 0, return 0
	// if len(nums) == 1, return len(memo)-1 which is equal to nums[0]
	// else the logic will be the same, and we traverse nums 2 places behind memo

	// this block doesn't execute if i<2 to begin with.
	for i := 2; i < len(memo); i++ {
		fmt.Println("nums is: ", nums)
		//if len(memo) > 2 {
		//	fmt.Println("memo is: ", memo[2:])
		//} else {
		fmt.Println("memo is: ", memo)
		//}

		fmt.Println("current index of i: ", i)
		fmt.Println("memo[i-2]", memo[i-2])
		fmt.Println("nums[i-2]", nums[i-2])
		fmt.Println("memo[i-1]", memo[i-1])

		// is the current nums item plus the previous memo item greater than the most recent memo item?
		// track if nums current item is greater than the memo item that we saved last
		memo[i] = max(memo[i-2]+nums[i-2], memo[i-1])
	}
	/*A runthrough of case7 would go as follows:
	nums starts:  [100 101 5]
	memo starts:  [0 0 0 0 0]
	current index of i:  2
	memo[i-2] 0
	nums[i-2] 100
	memo[i-1] 0
	is 100 more than 0?

	nums is now:  [100 101 5]
	memo is now:  [0 0 100 0 0]
	current index of i:  3
	memo[i-2] 0
	nums[i-2] 101
	memo[i-1] 100
	is 101 more than 100?

	nums is now:  [100 101 5]
	memo is now:  [0 0 100 101 0]
	current index of i:  4
	memo[i-2] 100
	nums[i-2] 5
	memo[i-1] 101
	is 105 more than 101?

	nums is now:  [100 101 5]
	memo is now:  [0 0 100 101 105]
	and we've iterated the entire nums array now, so return the last item from memo: 105
	*/

	// always in-bounds and safe to return, even when we're evaluating empty or single length nums arrays.
	// this is because we specify the length of our own array to be at least 2.
	fmt.Println("returning memo[len(memo)-1]", memo[len(memo)-1])

	// return the last item in our memo, which will be the combination of numbers (even if it's only 0).
	return memo[len(memo)-1]
}

func max(a int, b int) int {
	fmt.Printf("is %v more than %v?\n", a, b)
	if a > b {
		return a
	}
	return b
}

func main() {
	case1 := []int{1, 2, 3, 1}
	case2 := []int{1, 2, 3, 1, 6, 7, 8, 3, 4, 5}
	case3 := []int{1, 50, 15, 9, 3}
	case4 := []int{1, 2, 16, 1, 6, 7, 8, 3, 4, 5}
	case5 := []int{100, 2, 1, 1, 6, 7, 8, 3, 4, 5}
	case6 := []int{100, 2, 1, 1, 1, 1, 1, 900, 4, 5}
	case7 := []int{100, 101, 5}
	case8 := []int{100, 5}
	case9 := []int{5}
	case10 := []int{}
	fmt.Println(rob(case1))
	fmt.Println(rob(case2))
	fmt.Println(rob(case3))
	fmt.Println(rob(case4))
	fmt.Println(rob(case5))
	fmt.Println(rob(case6))
	fmt.Println(rob(case7))
	fmt.Println(rob(case8))
	fmt.Println(rob(case9))
	fmt.Println(rob(case10))
}
