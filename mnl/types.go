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
	ZCallF CallType = iota
	ZCallB CallType = iota
)

type environment struct {
	value     Expression
	variables map[int]Expression
}

type action struct {
	name        string
	expressions []Expression
}

func (a action) Call(e *environment) error {
	fmt.Println("Calling " + a.name)
	switch a.name {
	case "SET":
		n, err := Count(a.expressions[0], e)
		if err != nil {
			return err
		}
		e.variables[n] = a.expressions[1]
	case "FUN":
		n, err := Count(a.expressions[0], e)
		if err != nil {
			return err
		}
		e.variables[n] = LAMBDA{[]Expression{ZERO{}, a.expressions[1]}}
	case "START":
		par, err := Count(a.expressions[1], e)
		if err != nil {
			return err
		}

		exp, err := a.expressions[0].Call(par, e)
		fmt.Println(par, exp)
		if err != nil {
			return err
		}
		e.value = exp
	default:
		panic("CANNOT HAPPEN")
	}
	return nil
}

func Count(exp Expression, e *environment) (int, error) {
	var err error
	outside := 0
	for true {
		switch t := exp.(type) {
		case Number:
			return t.value + outside, nil
		case ZERO:
			return outside, nil
		default:
			exp, err = exp.Call(0, e)
			if err != nil {
				return -2, err
			}
			outside++
		}
	}
	panic("Cannot happen")
}

func (a action) Convert() string {
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
	Call(int, *environment) (Expression, error)
	Evaluate(*environment) (Expression, error)
	Convert() string
}

type Number struct {
	value int
}

func (number Number) Call(n int, e *environment) (Expression, error) {
	exp, ok := e.variables[number.value]
	if ok {
		return exp, nil
	}
	if number.value == -1 {
		return Number{1}, nil
	}
	return Number{number.value - 1}, nil
}

func (number Number) Evaluate(e *environment) (Expression, error) {
	return number, nil
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

func (c Call) Call(n int, e *environment) (Expression, error) {
	exp, err := c.Evaluate(e)
	if err != nil {
		return nil, err
	}
	return exp.Call(n, e)
}

func (c Call) Evaluate(e *environment) (Expression, error) {
	var err error
	exp := c.expression[0]

	if c.callType%3 == 1 {
		exp, err = exp.Call(0, e)
		if err != nil {
			return nil, err
		}
	}

	if c.callType > CallB {
		return exp.Call(0, e)
	}
	n, err := Count(c.expression[1], e)
	if err != nil {
		return nil, err
	}

	return c.expression[0].Call(n, e)
}

func (c Call) Convert() string {
	if c.callType > CallB {
		if c.callType == ZCallF {
			return c.expression[0].Convert() + "()()"
		}
		return c.expression[0].Convert() + "()"
	} else if c.callType == CallF {
		return c.expression[0].Convert() +
			"()(" + c.expression[1].Convert() + ")"
	}
	return c.expression[0].Convert() + "(" + c.expression[1].Convert() + ")"
}

type TSET struct {
	isFunc      bool
	expressions []Expression
}

func (s TSET) Call(nextN int, e *environment) (Expression, error) {
	exp, err := s.Evaluate(e)
	if err != nil {
		return nil, err
	}
	return exp.Call(nextN, e)
}

func (s TSET) Evaluate(e *environment) (Expression, error) {
	switch s.expressions[0].(type) {
	case ZERO:
		return s.expressions[2].Evaluate(e)
	}

	n, err := Count(s.expressions[0], e)
	if err != nil {
		return nil, err
	}

	oldval, ok := e.variables[n]
	if s.isFunc {
		e.variables[n] = LAMBDA{[]Expression{ZERO{}, s.expressions[1]}}
	} else {
		exp, err := s.expressions[1].Evaluate(e)
		if err != nil {
			return nil, err
		}
		e.variables[n] = exp
	}
	exp, err := s.expressions[2].Evaluate(e)
	if ok {
		e.variables[n] = oldval
	} else {
		delete(e.variables, n)
	}
	if err != nil {
		return nil, err
	}
	return exp, nil
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

type LAMBDA struct {
	expressions []Expression
}

func (l LAMBDA) Call(n int, e *environment) (Expression, error) {
	switch l.expressions[0].(type) {
	case ZERO:
		return l.expressions[1], nil
	default:
		panic("Not yet Implemented")
	}
}

func (l LAMBDA) Evaluate(e *environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (l LAMBDA) Convert() string {
	return "()()(" +
		l.expressions[0].Convert() +
		")(" +
		l.expressions[1].Convert() +
		")"
}

type ADD struct {
	n           int
	expressions []Expression
}

func (a ADD) Call(n int, e *environment) (Expression, error) {
	if a.n == 0 {
		return a.expressions[0].Call(n, e)
	}
	exp, err := a.expressions[0].Evaluate(e)
	if err != nil {
		return nil, err
	}

	if a.n == 1 {
		return exp, nil
	}

	return ADD{a.n - 1, []Expression{exp}}, nil
}

func (a ADD) Evaluate(e *environment) (Expression, error) {
	exp, err := a.expressions[0].Evaluate(e)
	if err != nil {
		return nil, err
	}
	return ADD{a.n, []Expression{exp}}, nil
}

func (a ADD) Convert() string {
	return strings.Repeat("(", a.n) +
		a.expressions[0].Convert() +
		strings.Repeat(")", a.n)
}

type ZERO struct{}

func (a ZERO) Call(n int, e *environment) (Expression, error) {
	panic("Not yet Implemented")
}

func (a ZERO) Evaluate(e *environment) (Expression, error) {
	return a, nil
}

func (a ZERO) Convert() string {
	return ""
}
