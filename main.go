// notLisp project main.go
package main

import (
	"fmt"
)

type Function interface {
	GetName() *Name
	//Find(*Name) Function
	Call(Function) Function
	Resolve() Function
	AppendCall(Function)
	GetParent() Function
	GetArgs() ([]Function, []string)
	GetSourceN() (charN int, fileN int)
}

var (
	fileN                             = 1
	charN                             = 1
	Overwritten map[Function]Function = map[Function]Function{}
	Functions   map[int]Function      = map[int]Function{}
)

func main() {
	program := "()()()()" //"()()(())(()())(*)"

	fmt.Print(program, " => ")
	fmt.Println(Count(Parse([]byte(program))))
	fmt.Println("Done")
	return
	mainScope := &MainScope{true, nil, fileN}

	scope := InterpretString(program, mainScope)
	pc := ParseCall(scope)
	for i := range pc.functions {
		fmt.Println(pc.functions[i])
	}
	fmt.Println("--------------")
	//scope = scope.Call(&Zero{mainScope, charN, fileN})
	ex := scope.GetName()
	fmt.Println(string(ex.bytes))
	fmt.Println(ex.Count())
}

func InterpretString(s string, f Function) Function {
	defer func() {
		fileN++
		charN = 1
	}()
	for _, c := range []byte(s) {
		f = Interpret(c, f)
	}
	return f
}
func InterpretBytes(bytes []byte, f Function) Function {
	defer func() {
		fileN++
		charN = 1
	}()
	for _, c := range bytes {
		f = Interpret(c, f)
	}
	return f
}

func Interpret(c byte, f Function) Function {
	defer func() { charN++ }()
	//fmt.Print(string(c))
	switch c {
	case '(':
		newScope := &Scope{f, nil, nil, charN, fileN}
		return newScope
	case ')':
		parent := f.GetParent()
		parent.AppendCall(f)
		return parent
	case '*':
		f.AppendCall(&DebugFunction{f, charN, fileN})
	default:
		//panic(fmt.Errorf("Tried to interpret unknown symbol %v", c))
	}
	return f
}

func RunFunc(f Function, arg Function) Function {
	//B, A := f.GetSourceN()
	//D, C := arg.GetSourceN()
	//fmt.Printf("Running %d:%d with %d:%d\n", A, B, C, D)

	_, isScope := f.(*Scope)
	if isScope {
		//EnterScope(scope)
		//defer ExitScope(scope)
	}

	name := f.GetName()
	c := name.Count()

	rf, ok := Functions[c]
	if ok {
		/*D, C = rf.GetSourceN()
		fmt.Printf("Function with name %d found, running "+
			"%d:%d instead of %d:%d\n", c, C, D, A, B)*/
		return rf.Call(arg)
	}
	//fmt.Printf("No function with name %d found, running "+
	//	"%d:%d\n", c, A, B)
	return f.Call(arg)
}

func EnterScope(scope *Scope) {
	if scope.name != nil {
		c := scope.name.Count()
		f, ok := Functions[c]
		if ok {
			Overwritten[scope.body] = f
		}
		fmt.Printf("Entered scope %v as %d\n", scope.body, scope.GetName())
		Functions[c] = scope.body
	}
}

func ExitScope(scope *Scope) {
	if scope.name != nil {
		c := scope.name.Count()
		body := Functions[c]
		f, ok := Overwritten[body]
		if ok {
			Functions[c] = f
			delete(Overwritten, body)
		} else {
			delete(Functions, c)
		}
	}
}
