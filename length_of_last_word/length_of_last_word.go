package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	splitWord := strings.Split(s, " ")

	sLen := len(splitWord)
	if sLen < 2 {
		return 0
	}

	wLen := len(splitWord[sLen-1])

	return wLen
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Hello"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("  "))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("Hello World Once More"))
}
