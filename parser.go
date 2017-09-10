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
				if prev == nil {
					prev = &Data{subString, []byte{}, Encapsulated}
				} else {
					prev = SCall(prev, subString)
				}
			}
		} else {
			switch c {
			case '(':
				indent++
				subString = []byte{'('}
			case ')':
				panic("Tried to exit scope")
			}
		}
	}

	if indent > 0 {
		panic("Unbalanced indents")
	}

	return prev
}

func SCall(data *Data, arg []byte) *Data {
	//c := Count(data)

	switch data.state {
	case Encapsulated:
		bytes := data.bytes[1 : len(data.bytes)-1]
		if len(bytes) == 0 {
			return &Data{[]byte{}, arg[1 : len(arg)-1], HasName}
		} else {
			return Parse(bytes)
		}
	case HasName:
		return &Data{arg, data.name, HasBody}
	case HasBody:
		//Add Scope here
		data = Parse(data.bytes)
		//Remove Scope here
		return data
	}
	panic("Cannot Happen")
}

func Count(data *Data) int {
	if data == nil {
		return 0
	}

	if data.state == HasBody {
		return -1
	}
	c := 0
	for data.state != HasName {
		data = Parse(data.bytes[1 : len(data.bytes)-1])
		if data == nil {
			return c + 1
		}
		if data.state == HasBody {
			return c
		}
		c++
	}
	return c
}
