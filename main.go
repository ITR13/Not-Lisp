// notLisp project main.go
package main

import (
	"fmt"
)

func main() {
	program := []byte("()()()()")

	data := Parse(program)
	//fmt.Print(program, " => ")
	c := Count(data)
	fmt.Println(c)
	fmt.Println("Done")
}
