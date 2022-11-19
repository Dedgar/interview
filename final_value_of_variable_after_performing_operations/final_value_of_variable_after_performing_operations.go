package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	fVal := 0

	//opLen := len(operations)
	for _, i := range operations {
		fmt.Println(i)
		// alternatively,
		// if i == "--X" || i == "X--" {
		if i[1] == '-' {
			fVal--
			continue
		}
		fVal++
	}
	/*switch {
	case i == "--X" || i == "X--":
		fVal--
	case i == "++X" || i == "X++":
		fVal++
	}*/

	return fVal
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
	fmt.Println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}
