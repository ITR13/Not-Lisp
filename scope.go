package main

import (
	"fmt"
)

var (
	CurrentScope map[int]*Data   = map[int]*Data{}
	Overwritten  map[*Data]*Data = map[*Data]*Data{}
)

func EnterScope(arg *Data, name int) {
	old, ok := CurrentScope[name]
	fmt.Printf("%d << %s (%v)\n", name, arg, old)
	if ok {
		Overwritten[arg] = old
	}
	CurrentScope[name] = arg
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
