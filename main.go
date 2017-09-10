// notLisp project main.go
package main

import (
	"fmt"
)

func main() {
	program := []byte("()()(())(()())(())")

	fmt.Print(string(program), " => ")
	data := Parse(program)
	fmt.Printf("%s, %s, %d => ", data.bytes, data.name, data.state)
	c := Count(data)
	fmt.Println(c)
	fmt.Println("Done")
}
