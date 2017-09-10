package main

func EnterScope(arg []byte, name int) {
	old, ok := CurrentScope[name]
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
}
