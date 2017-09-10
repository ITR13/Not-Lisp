// notLisp project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "()()((*)):(  ()()(((*))):( (*)( ((*))() ) )  )-( ()()((((*)))):( (((*)))() ))()"
	//s = strings.Replace(s, "*", "((((()))))", -1)
	s = strings.Replace(s, " ", "", -1)
	program := []byte(s)

	fmt.Print(string(program), " => ")
	data := Parse(program)
	fmt.Printf("%s, %s, %d => ", data.bytes, data.name, data.state)
	c := Count(data)
	fmt.Println(c)
	fmt.Println("Done")
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
