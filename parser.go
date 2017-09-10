package main

const (
	Encapsulated State = iota
	HasName      State = iota
	HasBody      State = iota
)

type State uint8

type Data struct {
	bytes, name []byte
	state       State
}

var (
	CurrentScope map[int]*[]byte     = map[int]*[]byte{}
	Overwritten  map[*[]byte]*[]byte = map[*[]byte]*[]byte{}
)

func Parse(s []byte) *Data {
	var prev *Data
	indent := 0
	subString := []byte{}
	for _, c := range s {
		if indent > 0 {
			subString = append(subString, c)
			switch c {
			case '(':
				indent++
			case ')':
				indent--
			}
			if indent == 0 {
				prev = Call(prev, subString)
			}
		} else {
			switch c {
			case '(':
				indent++
				subString = []byte{'('}
			case ')':
				panic("Unbalanced Parenthesises")
			}
		}
	}

	if indent > 0 {
		panic("Unbalanced Parenthesises")
	}

	if prev == nil {
		return &Data{[]byte{}, []byte{}, Encapsulated}
	}

	return prev
}

func Call(data *Data, arg []byte) *Data {
	if data == nil {
		return &Data{arg, []byte{}, Encapsulated}
	}

	c := Count(data)
	override, ok := CurrentScope[c]
	if ok && override != nil {
		return Parse(*override)
	}

	switch data.state {
	case Encapsulated:
		if len(data.bytes) == 0 {
			return &Data{[]byte{}, arg[1 : len(arg)-1], HasName}
		}

		return Parse(data.bytes[1 : len(data.bytes)-1])
	case HasName:
		return &Data{arg[1 : len(arg)-1], data.name, HasBody}
	case HasBody:
		c = Count(Parse(data.name))
		EnterScope(arg[1:len(arg)-1], c)
		data = Parse(data.bytes)
		ExitScope(c)
		return data
	}
	panic("Cannot Happen")
}

func Count(data *Data) int {
	if data == nil {
		return 0
	}
	if data.state == HasName {
		return -1
	}

	c := 0
	for data.state != HasName {
		if data.state == HasBody {
			data = Parse(data.bytes)
			c++
		}
		if len(data.bytes) == 0 {
			return c
		}
		data = Parse(data.bytes[1 : len(data.bytes)-1])
		c++
	}
	return c
}

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
