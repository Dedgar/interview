package main

import(
  "fmt"
)

func defangIPaddr(address string) string {
    dfIP := ""

    for _, c := range address {
        if string(c) == "." {
            dfIP += "[.]"
        } else {
            dfIP += string(c)
        }
    }
    return dfIP
}

func main() {
  fmt.Println(defangIPaddr("127.0.0.1"))
}
