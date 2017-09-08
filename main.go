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
	GetArgs() ([]Function, []string)
}

func main() {
	program := "()()()"
	mainScope := &MainScope{true, nil}

	scope := InterpretString(program, mainScope)
	pc := ParseCall(scope)
	for i := range pc.functions {
		fmt.Println(pc.functions[i])
	}
	fmt.Println("--------------")
	//scope = scope.Call(&Name{scope, []byte{}, 0})
	ex := scope.GetName()
	fmt.Println(string(ex.bytes))
	fmt.Println(ex.Count())
}

func InterpretString(s string, f Function) Function {
	for _, c := range []byte(s) {
		f = Interpret(c, f)
	}
	return f
}
func InterpretBytes(bytes []byte, f Function) Function {
	for _, c := range bytes {
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
