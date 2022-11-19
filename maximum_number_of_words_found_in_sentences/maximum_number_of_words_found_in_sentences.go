package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	c := 0

	for _, i := range sentences {
		f := strings.Fields(i)
		l := len(f)
		if l > c {
			c = l
		}
		/* alternatively

		s := strings.Split(i, " ")
		l := len(s)
		if l > c {
			c = l
		}
		*/
	}
	return c
}

func main() {
	sl1 := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	sl2 := []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(sl1))
	fmt.Println(mostWordsFound(sl2))
}
