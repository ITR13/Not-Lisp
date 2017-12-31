package interpreter

var (
	CurrentScope map[int]*Data   = map[int]*Data{}
	Overwritten  map[*Data]*Data = map[*Data]*Data{}
)

func EnterScope(arg *Data, name int) {
	//fmt.Println("Scope: ", name, string(arg.bytes))
	old, ok := CurrentScope[name]
	if ok {
		Overwritten[arg] = old
	}
	CurrentScope[name] = arg
}

func ExitScope(name int) {
	//fmt.Println("EScope: ", name)
	arg, ok := CurrentScope[name]
	if !ok {
		panic("Tried to exit scope that doesn't exist")
	}
	old := Overwritten[arg]
	CurrentScope[name] = old
}
