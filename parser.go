package main

import (
	"fmt"
	"strings"
)

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

var AAA int = 0

func Parse(s []byte) *Data {
	I := make([]int, len(s))

	aaa := AAA
	AAA++

	var prev *Data
	depth := 0
	subString := []byte{}
	for i, c := range s {
		I[i] = depth
		if depth > 0 {
			subString = append(subString, c)
			switch c {
			case '(':
				depth++
			case ')':
				depth--
				I[i]--
			}
			if depth == 0 {
				if /*aaa == 0 || aaa == 6 || aaa == 19 ||*/ aaa == 33 {
					fmt.Println(aaa, string(subString))
				}

				prev = Call(prev, subString)
				if prev == nil {
					fmt.Println("Call returned nil: ", string(s))
					PrintWithIndent(s, I[:i+1])
					return nil
				}
			}
		} else {
			switch c {
			case '(':
				depth++
				subString = []byte{'('}
			case ')':
				fmt.Println("Unbalanced Parenthesises: ", string(s))
				PrintWithIndent(s, I[:i+1])
				return nil
			}
		}
	}

	if depth > 0 {
		fmt.Println("Unbalanced Parenthesises: ")
		PrintWithIndent(s, I)
		return nil
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

	if data.state != HasBody {
		c := Count(data)

		if c == -2 {
			return nil
		}

		override, ok := CurrentScope[c]
		if ok && override != nil {
			return override
		}
	}

	switch data.state {
	case Encapsulated:
		if len(data.bytes) == 0 {
			return &Data{[]byte{}, Strip(arg), HasName}
		}

		return Parse(Strip(data.bytes))
	case HasName:
		return &Data{Strip(arg), data.name, HasBody}
	case HasBody:
		if len(data.name) != 0 {
			c := Count(Parse(data.name))
			if c == -2 {
				return nil
			}
			EnterScope(ReScope(Parse(Strip(arg))), c)
			data = Parse(data.bytes)
			ExitScope(c)
		} else {
			data = Parse(data.bytes)
		}
		return data
	}
	panic("Cannot Happen")
}

func Count(data *Data) int {
	if data == nil {
		return -2
	}
	if data.state == HasName {
		return -1
	}

	c := 0
	for data.state != HasName {
		for data.state == HasBody {
			data = Parse(data.bytes)
			c++
			if data == nil {
				return -2
			}
		}
		if len(data.bytes) == 0 {
			return c
		}
		data = Parse(Strip(data.bytes))
		c++
		if data == nil {
			return -2
		}
	}
	return c
}

func Strip(bytes []byte) []byte {
	depth := 0
	exited := false
	for _, c := range bytes {
		switch c {
		case '(':
			depth++
		case ')':
			depth--
		}
		if depth == 0 {
			if exited {
				panic("Tried to strip unencapsulated bytes")
			}
			exited = true
		}
	}
	return bytes[1 : len(bytes)-1]
}

func ReScope(data *Data) *Data {
	depth := 0

	if data.state == HasBody || data.state == HasName {
		return data
	}

	for data.state != HasBody {
		if len(data.bytes) == 0 {
			return &Data{
				[]byte(strings.Repeat("(", depth) + strings.Repeat(")", depth)),
				[]byte{}, Encapsulated}

		}
		data = Parse(Strip(data.bytes))
		depth++
	}
	return &Data{
		[]byte(strings.Repeat("(", depth) +
			string(data.bytes) +
			strings.Repeat(")", depth)),
		[]byte{}, Encapsulated}
}
