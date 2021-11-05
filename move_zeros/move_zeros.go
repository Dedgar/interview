package main

import "fmt"

func moveZeroesVerbose(nums []int) {
	fmt.Println("***starting new run***")
	fmt.Println(nums)
	numLen := len(nums)

	if numLen <= 1 {
		return
	}

	nonZeroPos := 0

	for pos := 0; pos <= numLen-1; pos++ {
		fmt.Println("pos is currently", pos)
		fmt.Println("nonZeroPos is currently", nonZeroPos)
		if nums[pos] == 0 {
			fmt.Println("hit zero, skipping")
			continue
		}
		fmt.Printf("nums[pos] value: %v nums[nonZeroPos] value: %v\n", nums[pos], nums[nonZeroPos])
		// this if block skips the position if it's the same as the last non-zero value's position
		if nums[pos] != nums[nonZeroPos] {
			fmt.Println("they are not equal ")
			nums[nonZeroPos] = nums[pos]
			fmt.Println("now it changed to: ")
			fmt.Println(nums)
			fmt.Printf("nums[pos]: %v nums[nonZeroPos]: %v\n", nums[pos], nums[nonZeroPos])
		} else {
			fmt.Println("equal! skipping")
		}
		nonZeroPos++
	}

	fmt.Println("first for loop completed, nums is now: ")
	fmt.Println(nums)

	fmt.Printf("for nonZeroPos %v < numLen %v\n", nonZeroPos, numLen)

	for nonZeroPos < numLen {
		fmt.Printf("if nums[nonZeroPos] %v != 0\n", nums[nonZeroPos])
		if nums[nonZeroPos] != 0 {
			fmt.Printf("nums[nonZeroPos] %v will be replaced with 0\n", nums[nonZeroPos])
			nums[nonZeroPos] = 0
			fmt.Printf("nums[nonZeroPos] %v switched to 0\n", nums[nonZeroPos])
		}
		nonZeroPos++
	}
	fmt.Println(nums)
}

// "now here's a little lesson, in trickery -"
// With the aim of minimizing write operations (not speed or memory), start iterating the array from the beginning.
// skip over array positions where the value is either 0 or equal to the last non-zero position.
// after we finish iterating, overwrite every non-zero value that's in a position greater than what's tracked by nonZeroPos.
// example:
// [1 0 0 0 0 0 0 0 0 1] start with this, iterate it in order, and move any non-zero numbers found to the last non-zero position as we go.
// [1 1 0 0 0 0 0 0 0 1] after moving the non-zero values as far down as we can (in order), we see we still have to clean the non-zeros from their original positions.
// [1 1 0 0 0 0 0 0 0 0] we're left with the final result after replacing any extraneus non-zero values greater than nonZeroPos
func moveZeroes(nums []int) {
	numLen := len(nums)

	if numLen <= 1 {
		return
	}

	nonZeroPos := 0

	for pos := 0; pos <= numLen-1; pos++ {
		if nums[pos] == 0 {
			continue
		}

		// skip the position if its value is the same as the last non-zero value.
		// otherwise, replace it with the value of our current position.
		if nums[pos] != nums[nonZeroPos] {
			nums[nonZeroPos] = nums[pos]
		}
		nonZeroPos++
	}

	// start at the last non-zero number, and replace every value from here on out with a zero.
	for nonZeroPos < numLen {
		if nums[nonZeroPos] != 0 {
			nums[nonZeroPos] = 0
		}
		nonZeroPos++
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	b := []int{0}
	c := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	d := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := []int{1}

	moveZeroes(a)
	moveZeroes(b)
	moveZeroes(c)
	moveZeroes(d)
	moveZeroes(e)
	moveZeroes(f)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
