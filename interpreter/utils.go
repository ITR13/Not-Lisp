package interpreter

import (
	"fmt"
	"strings"
)

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
