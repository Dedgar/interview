package main

import "fmt"


// more math trickery.
// the faster of the two functions for larger arrays.
// by subtracting the current item's value from our target, we only need to find one other number, which will be equal to to the remainder of our subtraction operation.
// identify middle of the slice by dividing in half
// if the middle is still smaller than our target, set the middle as the new bottom, and run again.
// if the middle is still greater than our target, set the middle as the new top, and run again.
// progressing through the slice this way, we're guaranteed to be left with the target as per the lc problem constraints.
func twoSumBinarySearch(numbers []int, target int) []int {
	fmt.Println("looking for", target)
	top := len(numbers) - 1
	solution := make([]int, 2)

	fmt.Println("starting with numbers")
	fmt.Println(numbers)
	for pos := 0; pos <= top; pos++ {
		bot := pos + 1
		sub := target - numbers[pos]

		for bot <= top {
			mid := (bot + top) >> 1

			if numbers[mid] == sub {
				solution[0], solution[1] = pos+1, mid+1
				return solution
			} else if numbers[mid] < sub {
				bot = mid + 1
			} else {
				top = mid - 1
			}
		}
	}
	return solution
}

func twoSum(numbers []int, target int) []int {
	fmt.Println("looking for", target)
	bot := 0
	top := len(numbers) - 1
	solution := make([]int, 2)

	for bot < top {
		sum := numbers[bot] + numbers[top]
		if sum == target {
			fmt.Println("found sum matching target by combining: ", bot, top)
			solution[0], solution[1] = bot+1, top+1 //numbers[bot], numbers[top]
			break
		} else if sum < target {
			bot++
		} else {
			top--
		}
	}
	return solution
}

func main() {
	a := []int{2, 7, 11, 15}
	b := []int{2, 3, 4}
	c := []int{-1, 0}
	d := []int{2, 7, 11, 15, 16, 17, 18, 19, 20, 21, 22, 23, 34, 45, 56, 67, 78, 89, 99}

	/*
		fmt.Println("***starting first***")
		fmt.Println(twoSum(a, 9))
		fmt.Println("***starting second***")
		fmt.Println(twoSum(b, 6))
		fmt.Println("***starting third***")
		fmt.Println(twoSum(c, -1))
	*/
	fmt.Println("***starting fourth***")
	fmt.Println(twoSumBinarySearch(a, 9))
	fmt.Println("***starting fifth***")
	fmt.Println(twoSumBinarySearch(b, 6))
	fmt.Println("***starting sixth***")
	fmt.Println(twoSumBinarySearch(c, -1))
	fmt.Println("***starting seventh***")
	fmt.Println(twoSumBinarySearch(d, 9))
	fmt.Println("***starting eighth***")
	fmt.Println(twoSumBinarySearch(d, 45))
}
