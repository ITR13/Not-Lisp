// notLisp project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `
		()()( (()) )(
			(())()( (()) )
		)(
			()()( (*) )(
				()()( (*)() )(
					(*)()()
				)()
			)
		)
		`
	s = strings.Replace(s, "*", "((((()))))", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	data := Parse([]byte(s))
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
