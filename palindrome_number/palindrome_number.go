package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func isPalindrome(x int) bool {
	b := strconv.Itoa(x)
	r := reverse(b)

	if b == r {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(1234321))
}
