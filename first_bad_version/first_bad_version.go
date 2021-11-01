package main

import (
	"fmt"
)

// lc testcase mockup
func isBadVersion(version int) bool {
	fmt.Println("is", version, "bad: equal or greater than 6?")
	if version >= 6 {
		fmt.Println("yes it is")
		return true
	}
	fmt.Println("safe, is not greater than or equal to 6: ", version)
	return false
}

// lc testcase mockup
func isBadVersion2(version int) bool {
	fmt.Println("is", version, "bad: equal or greater than 2?")
	if version >= 2 {
		fmt.Println("yes it is")
		return true
	}
	fmt.Println("safe, is not greater than or equal to 2: ", version)
	return false
}

// find the number in the middle, then ignore either the entire lower or upper half
// keep inching toward the solution in a binary fashion by narrowing the range this way
//
// mid is 10
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
//
// top is now 10
// mid is now 5
// 1 2 3 4 5 6 7 8 9 10
//
// bot is now 5
// mid is now 7
// 5 6 7 8 9 10
//
// top is now 7
// mid is now 6
// 5 6 7
//
// 6 is now top
// mid is now 5
// 5 6
//
// left evaluates tp false, and nothing remains to the right, so we've hit the first of the Bad versions
func firstBadVersion(n int) int {
	bot := 1
	top := n
	fmt.Println("starting number is: ", n)

	for bot < top {
		mid := (bot + top) >> 1
		fmt.Println("mid is: ", mid)
		if mid == 1 {
			return mid
		}
		if isBadVersion(mid) {
			top = mid
			fmt.Println("top is now: ", top)
		} else {
			bot = mid + 1
			fmt.Println("bot is now: ", bot)
		}
	}

	return bot
}

func firstBadVersion2(n int) int {
	bot := 1
	top := n
	fmt.Println("starting number is: ", n)

	for bot < top {
		mid := (bot + top) >> 1
		fmt.Println("mid is: ", mid)

		if isBadVersion2(mid) {
			top = mid
			fmt.Println("top is now: ", top)
		} else {
			bot = mid + 1
			fmt.Println("bot is now: ", bot)
		}
	}

	return bot
}

func main() {
	fmt.Println("***starting first***")
	fmt.Println(firstBadVersion(1))
	fmt.Println("***starting second***")
	fmt.Println(firstBadVersion(20))
	fmt.Println("***starting third***")
	fmt.Println(firstBadVersion(100))
	fmt.Println("***starting fourth***")
	fmt.Println(firstBadVersion2(2))
}
