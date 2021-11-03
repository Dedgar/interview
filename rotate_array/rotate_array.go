package main

import "fmt"

// not surprisingly, we can solve this with more math.
// we don't need to actually rotate the array multiple times, just figure out where the numbers should be.
// since any rotation past the length of the array would be duplicate effort, we only need to find the final difference.
// include everything [after:] the mod in one slice, and everything [:until] the mod in another.
func rotate(nums []int, k int) {
	aLen := len(nums)
	if aLen <= 1 {
		return
	}

	mod := aLen - (k % aLen)
	fmt.Printf("%v - ( %v %% %v )\n", aLen, k, aLen)
	fmt.Printf("remainder/modulo %v, minus len(nums), after k shift amount %v divided by len(nums) %v\n", mod, k, aLen)

	begin := nums[mod:]
	end := nums[:mod]

	fmt.Println("so our new slice should begin with:")
	fmt.Println(begin)
	fmt.Println("and end with:")
	fmt.Println(end)

	begin = append(begin, end...)

	copy(nums, begin)
}

func main() {
	testCase1 := []int{1, 2, 3, 4, 5, 6, 7}
	testCase2 := []int{-1, -100, 3, 99}
	testCase3 := []int{-1, -100, 3, 99, 8, 4, 2, 7, 0, 3}

	fmt.Println("starting with:        ", testCase1)
	rotate(testCase1, 3)
	fmt.Println("after rotating by 3:  ", testCase1)
	fmt.Println("starting with:        ", testCase2)
	rotate(testCase2, 2)
	fmt.Println("after rotating by 2:  ", testCase2)
	fmt.Println("starting with:        ", testCase3)
	rotate(testCase3, 33)
	fmt.Println("after rotating by 33: ", testCase3)
}
