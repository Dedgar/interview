package main

import "fmt"

var (
	am = map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "d",
		"E": "e",
		"F": "f",
		"G": "g",
		"H": "h",
		"I": "i",
		"J": "j",
		"K": "k",
		"L": "l",
		"M": "m",
		"N": "n",
		"O": "o",
		"P": "p",
		"Q": "q",
		"R": "r",
		"S": "s",
		"T": "t",
		"U": "u",
		"V": "v",
		"W": "w",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}
)

func toLowerCase(str string) string {
	lwr := ""

	for _, c := range str {
		if l, ok := am[string(c)]; ok {
			lwr += l
		} else {
			lwr += string(c)
		}
	}
	return lwr
}

func main() {
	fmt.Println(toLowerCase("Hello"))
}
