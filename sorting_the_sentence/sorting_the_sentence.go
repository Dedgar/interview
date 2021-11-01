package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	wordSlice := make([]string, 10)
	sent := strings.Split(s, " ")

	// take word, trim last character, use last character as the word's position in the new slice
	for _, word := range sent {
		strLen := len(word)
		lastDigit, err := strconv.Atoi(word[strLen-1:])
		if err != nil {
			fmt.Println("Error converting to int: ", word[strLen-1])
		}

		baseWord := word[:strLen-1]

		wordSlice[lastDigit] = baseWord
	}

	// Join each item in the slice into a string, separated by spaces.
	// Then get rid of leading and trailing spaces resulting from that operation.
	return strings.Trim(strings.Join(wordSlice, " "), " ")
}

func main() {
	//fmt.Println("sorting")
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
	fmt.Println(sortSentence("QcGZ4 TFJStgu3 HvsRImLBfHd8 PaGqsPNo9 mZwxlrYzanuVOUZoyNjt1 fzhdtYIen6 mV7 LKuaOtefsixxo5 pwdEK2"))
}
