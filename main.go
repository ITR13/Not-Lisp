// notLisp project main.go
package main

import (
	"fmt"
)

func main() {
	program := "()()()()"

	fmt.Print(program, " => ")
	fmt.Println(Count(Parse([]byte(program))))
	fmt.Println("Done")
}
