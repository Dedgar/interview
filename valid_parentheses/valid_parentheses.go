package main

import (
	"fmt"
)

var (
	pMap = map[rune]rune{']': '[', '}': '{', ')': '('}
)

type Stack []rune

func isValid(s string) bool {
	ns := make(Stack, 0)

	for _, i := range s {
		l := len(ns)
		if l == 0 {
			ns = append(ns, i)
			continue
		}
		if pMap[i] == ns[l-1] {
			ns = ns[:l-1]
		} else {
			ns = append(ns, i)
		}
	}
	return len(ns) == 0
}

func main() {
	//fmt.Println("result is:", isValid("()"))
	fmt.Println("result is:", isValid("(())"))
	//fmt.Println("result is:", isValid("{}()"))
	//fmt.Println("result is:", isValid("(]"))

	//fmt.Println(isValid("()()[[[]]{}]"))
	//fmt.Println(isValid("()()[[]]{}]"))
	//fmt.Println(isValid(")"))
	//fmt.Println(isValid("("))
}
