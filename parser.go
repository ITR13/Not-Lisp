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
				prev = Call(prev, subString)
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

	if prev == nil {
		return &Data{[]byte{}, []byte{}, Encapsulated}
	}

	return prev
}

func Call(data *Data, arg []byte) *Data {
	if data == nil {
		return &Data{arg, []byte{}, Encapsulated}
	}

	//c := Count(data)

	switch data.state {
	case Encapsulated:
		//fmt.Println("IE:", data.bytes)
		if len(data.bytes) == 0 {
			return &Data{[]byte{}, arg[1 : len(arg)-1], HasName}
		}

		return Parse(data.bytes[1 : len(data.bytes)-1])
	case HasName:
		//fmt.Println("HN:", data.name, data.bytes)
		return &Data{arg[1 : len(arg)-1], data.name, HasBody}
	case HasBody:
		//fmt.Println("HB:", data.bytes)
		//Add Scope here
		data = Parse(data.bytes)
		//fmt.Println("HB<-", data)
		//Remove Scope here
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
