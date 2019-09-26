package main

import (
	"fmt"
	"strconv"
	"strings"
)

func positive(x int) string {
	var b strings.Builder

	y := strconv.Itoa(x)

	for i := len(y) - 1; i >= 0; i-- {
		fmt.Fprintf(&b, "%v", string(y[i]))
	}

	return b.String()
}

func reverse(x int) int {
	B := ""

	if x > 0 {
		B = positive(x)
	} else {
		z := -x
		B = positive(z)
	}

	res, err := strconv.Atoi(B)

	if err != nil {
		fmt.Println("strconv Error: ", err)
	}

	if x < 0 {
		res = -res
	}

	if res > 2147483647 || res < -2147483647 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(1234))
	fmt.Println(reverse(-1234))
}
