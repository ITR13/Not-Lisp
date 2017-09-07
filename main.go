// notLisp project main.go
package main

import (
	"fmt"
)

type Function interface {
	GetName() *Name
	Find(*Name) Function
	Call(Function) Function
	AppendCall(Function)
	GetParent() Function
}

func main() {
	programs := []string{
		"()()()()",
	}

	for _, program := range programs {

		mainScope := &MainScope{}

		scope := InterpretString(program, mainScope)
		ex := scope.GetName()
		fmt.Println(string(ex.bytes))

		if len(ex.bytes) == 0 {
			fmt.Println(program, " => ", -1)
		} else {
			var s uint64
			for i := range ex.bytes {
				s += uint64(ex.bytes[i])
			}

			fmt.Println(program, " => ", s)
		}
	}
}

func InterpretString(s string, f Function) Function {
	for _, c := range []byte(s) {
		f = Interpret(c, f)
	}
	return f
}

func Interpret(c byte, f Function) Function {
	//fmt.Print(string(c))
	switch c {
	case '(':
		newScope := &Scope{f, nil, nil}
		return newScope
	case ')':
		parent := f.GetParent()
		parent.AppendCall(f)
		return parent
	default:
		//panic(fmt.Errorf("Tried to interpret unknown symbol %v", c))
	}
	return f
}

func RunFunc(f Function, arg Function) Function {
	name := f.GetName()
	rf := f.Find(name)
	if rf != nil {
		return rf.Call(arg)
	}
	return f.Call(arg)
}
