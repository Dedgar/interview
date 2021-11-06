package main

import "fmt"

func reverseString(s []byte) {
	l1 := len(s)

	// bail if we don't need to reverse the array.
	if l1 < 2 {
		fmt.Println(string(s))
		return
	}

	// keep some reference numbers of where we're at in the array,
	// working both ends toward the middle.
	p1 := 0
	p2 := l1 - 1

	// stop at the middle of the array,
	// works for arrays with both even and odd numbers of items.
	for p1 < p2 {
		// take note of the current items
		b1 := s[p1]
		b2 := s[p2]

		// and then swap them
		s[p1] = b2
		s[p2] = b1

		p1++
		p2--
	}
	fmt.Println(string(s[:]))
}

// a slightly cleaner version that only swaps values if they're different.
// reduces the total number of write operations for arrays that have lots of similar values.
func reverseString2(s []byte) {
	l1 := len(s)

	// bail if we don't need to reverse the array.
	if l1 < 2 {
		fmt.Println(string(s))
		return
	}

	// keep some reference numbers of where we're at in the array,
	// as we work from both ends toward the middle.
	bot := 0
	top := l1 - 1

	for bot < top {
		// swap the current items if they aren't the same already
		if s[bot] != s[top] {
			s[bot], s[top] = s[top], s[bot]
		}

		bot++
		top--
	}
	fmt.Println(string(s[:]))
}

func main() {
	reverseString2([]byte{'s', 'o', 'm', 'e', 's', 't', 'r', 'i', 'n', 'g', 't', 'o', 'r', 'e', 'v', 'e', 'r', 's', 'e'})
	reverseString2([]byte{'s', 'o', 'u', 'p'})
	reverseString2([]byte{'s'})
}
