package main

import (
	"log"
)

type Scope struct {
	parent Function
	body   Function
	name   *Name
}

func (scope *Scope) Call(arg Function) Function {
	if scope.body == nil {
		return &Zero{scope}
	}
	return scope.body
}

func (scope *Scope) Find(name *Name) Function {
	if scope.name != nil && scope.name.Equals(name) {
		return scope.body
	}

	return scope.parent.Find(name)
}

func (scope *Scope) GetName() *Name {
	if scope.body == nil {
		return &Name{scope.body, []byte{1}, 0}
	}

	name := scope.body.GetName()
	name.AddOne()
	return name
}

func (scope *Scope) AppendCall(f Function) {
	if scope.body == nil {
		scope.body = f
	} else {
		oldBody := scope.body
		scope.body = &Call{scope, oldBody, f.Call(nil)}
	}
}

func (scope *Scope) GetParent() Function {
	return scope.parent
}

func (scope *Scope) GetArgs() ([]Function, []string) {
	return []Function{scope.body, scope.name}, []string{"Body", "Name"}
}

type MainScope struct {
	allowExit bool
	body      Function
}

func (scope *MainScope) Call(arg Function) Function {
	return RunFunc(scope.body, arg)
}

func (scope *MainScope) Find(name *Name) Function {
	return nil
}

func (scope *MainScope) GetName() *Name {
	if scope.body == nil {
		return &Name{scope.body, []byte{0}, 0}
	}
	return scope.body.GetName()
}

func (scope *MainScope) AppendCall(f Function) {
	if scope.body == nil {
		scope.body = f
	} else {
		oldBody := scope.body
		scope.body = &Call{scope, oldBody, f.Call(nil)}
	}
}

func (scope *MainScope) GetParent() Function {
	if !scope.allowExit {
		log.Fatal("Tried to exit main-scope")
	}
	return nil
}

func (scope *MainScope) GetArgs() ([]Function, []string) {
	return []Function{scope.body}, []string{"Body"}
}
