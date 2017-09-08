package main

import (
	"fmt"
	"reflect"
)

type ParsedCall struct {
	found     map[Function]int
	functions []string
}

func ParseCall(f Function) *ParsedCall {
	pc := &ParsedCall{make(map[Function]int), make([]string, 1)}
	pc.functions[0] = "NIL"
	pc.found[nil] = 0

	mainCall := 0
	for f != nil {
		pc.Add(f, mainCall)
		f = f.Call(nil)
	}

	return pc
}

func (pc *ParsedCall) Add(f Function, callID int) {
	id := pc.GetIDOf(f)
	pc.functions[id] += fmt.Sprintf(" => %d", callID)
}

func (pc *ParsedCall) GetIDOf(f Function) int {
	id, ok := pc.found[f]
	if ok {
		return id
	}

	id = len(pc.functions)
	pc.functions = append(pc.functions, "[Not Set]")
	pc.found[f] = id

	s := ""
	args, names := f.GetArgs()
	for i := range args {
		s += fmt.Sprintf("[%s: %d] ", names[i], pc.GetIDOf(args[i]))
	}
	name := f.GetName()
	pc.functions[id] = fmt.Sprintf("%v:\t[ %s] in %d, named %d.%s (%d)",
		reflect.TypeOf(f), s, pc.GetIDOf(f.GetParent()),
		name.infinitum, name.bytes, name.Count(),
	)

	return id
}
