// notLisp project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	data := &Data{[]byte("()()(())(()())((()))"), []byte{}, HasBody}

	data = Call(data, []byte("()"))
	fmt.Println(data, Count(data))
}

func PrintWithIndent(s []byte, I []int) {
	max := 0
	for _, i := range I {
		if i > max {
			max = i
		}
	}

	for i := range s {
		if i >= len(I) {
			fmt.Printf("--:%s\n", string(s[i:]))
			break
		} else {
			fmt.Printf("%02x:%s%c\n", I[i], strings.Repeat("  ", I[i]), s[i])
		}
	}
}
