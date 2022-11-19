package main

import "fmt"

func minPartitions(n string) int {
	var b byte = '0'
	for i := 0; i < len(n); i++ {
		if n[i] > b {
			b = n[i]
		}
	}
	return int(b - '0')
}

func main() {
	fmt.Println(minPartitions("12345"))
}
