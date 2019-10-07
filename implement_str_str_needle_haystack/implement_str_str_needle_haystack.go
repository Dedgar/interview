package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) bool {
	//fmt.Println("comparing", a, b)
	if a < b {
		//fmt.Println("a smaller than b", "a:", a, "b:", b)
		return false
	}

	//if len(b) == 0 {
	//	fmt.Println("b is zero though")
	//	return true
	//}

	if len(b) > 1 {
		if a[0] == b[0] {
			res := compare(a[1:], b[1:])
			if res {
				//fmt.Println("nested")
				return true
			}
		}
	} else if len(b) == 1 && a[0] == b[0] {
		return true
	}
	return false
}

func strStr(haystack string, needle string) int {
	//fmt.Println("on items", haystack, needle)
	switch {
	case haystack == "" && len(needle) > 0 || len(haystack) < len(needle):
		return -1
	case needle == "":
		return 0
	}

	if strings.Contains(haystack, needle) {
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				//fmt.Println("starting a comparison")
				res := compare(haystack[i:], needle)
				if res {
					return i
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("1234", "tests"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("mississippi", "issip"))
}
