package main

import (
	"fmt"
)

const (
	Encapsulated State = iota
	Named        State = iota
	Bodied       State = iota
)

type State uint8

type Data struct {
	bytes, name []byte
	state       State
}

func Parse(s []byte) *Data {
	fmt.Println(s)
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
			case '*':
				fmt.Println("Called Debug")
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

	fmt.Println(data)

	switch data.state {
	case Encapsulated:
		bytes := data.bytes[1 : len(data.bytes)-1]
		if len(bytes) == 0 {
			return &Data{[]byte{}, arg[1 : len(arg)-1], Named}
		} else {
			return Parse(bytes)
		}
	case Named:
		return &Data{arg, data.name, Bodied}
	case Bodied:
		//Add Scope here
		data = Parse(data.bytes)
		//Remove Scope here
		return data
	}
	panic("Cannot Happen")
}

func Count(data *Data) int {
	if data.state == Bodied {
		return -1
	}
	c := 0
	for data.state != Named {
		data = Parse(data.bytes[1 : len(data.bytes)-1])
		if data == nil {
			return c + 1
		}
		if data.state == Bodied {
			return c
		}
		c++
	}
	return c
}
