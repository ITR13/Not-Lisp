package main

import (
	"fmt"
)

func EnterScope(arg []byte, name int) {
	old, ok := CurrentScope[name]
	fmt.Printf("%d << %s (%v)\n", name, arg, old)
	if ok {
		Overwritten[&arg] = old
	}
	CurrentScope[name] = &arg
}

func ExitScope(name int) {
	arg, ok := CurrentScope[name]
	if !ok {
		panic("Tried to exit scope that doesn't exist")
	}
	old := Overwritten[arg]
	CurrentScope[name] = old
	fmt.Printf("%d >> %v (%s)\n", name, old, arg)
}
