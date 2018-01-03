package mnl

import (
	"fmt"
	"strings"
)

type CallType uint8

const (
	CallN  CallType = iota
	CallF  CallType = iota
	CallB  CallType = iota
	ZCallN CallType = iota
	ZCallT CallType = iota
	ZCallB CallType = iota
)

type Environment struct {
}

type Action struct {
	name        string
	expressions []Expression
}

func (a Action) Call(e Environment) error {
	panic("Not yet Implemented")
}

func (a Action) Convert() string {
	switch a.name {
	case "SET":
		return "()()(" +
			a.expressions[0].Convert() +
			")(*)(" +
			a.expressions[1].Convert() +
			")"
	case "FUN":
		return "()()(" +
			a.expressions[0].Convert() +
			")(*)(()()()(" +
			a.expressions[1].Convert() +
			"))"
	case "START":
		return a.expressions[0].Convert() +
			"(" + a.expressions[1].Convert() + ")"
	}
	panic(fmt.Errorf("Missing Convert for action {0}", a.name))
}

type Expression interface {
	Call(Environment) (Expression, error)
	Convert() string
}

type Number struct {
	value int
}

func (n Number) Call(e Environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (n Number) Convert() string {
	if n.value == 0 {
		return "()()"
	} else if n.value == -1 {
		return "()()()"
	}
	return strings.Repeat("(", n.value) + strings.Repeat(")", n.value)
}

type Call struct {
	callType   CallType
	expression []Expression
}

func (c Call) Call(e Environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (c Call) Convert() string {
	if c.callType > CallB {
		return c.expression[0].Convert()
	}
	return c.expression[0].Convert() + "(" + c.expression[1].Convert() + ")"
}

type TSET struct {
	isFunc      bool
	expressions []Expression
}

func (s TSET) Call(e Environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (s TSET) Convert() string {
	if s.isFunc {
		return "()()(" +
			s.expressions[0].Convert() +
			")(" +
			s.expressions[2].Convert() +
			")(()()()(" +
			s.expressions[1].Convert() +
			"))"
	}
	return "()()(" +
		s.expressions[0].Convert() +
		")(" +
		s.expressions[2].Convert() +
		")(" +
		s.expressions[1].Convert() +
		")"
}

type ADD struct {
	n           int
	expressions []Expression
}

func (a ADD) Call(e Environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (a ADD) Convert() string {
	return strings.Repeat("(", a.n) +
		a.expressions[0].Convert() +
		strings.Repeat(")", a.n)
}
