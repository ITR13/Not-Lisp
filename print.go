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
	pc.found[(*Name)(nil)] = 0

	mainCall := 0

	for f != nil {
		pc.Add(f, mainCall)
		mainCall++
		_, ok := f.(*Zero)
		if ok {
			break
		}
		_, ok = f.(*NegOne)
		if ok {
			break
		}
		f = f.Call(nil)
		if mainCall > 10000 {
			fmt.Println("Maximum depth reached")
			f = nil
		}
		charN++
	}
	fileN++
	return pc
}

func (pc *ParsedCall) Add(f Function, callID int) {
	id := pc.GetIDOf(f)
	pc.functions[id] += fmt.Sprintf(" => %d", callID)
}

func (pc *ParsedCall) GetIDOf(f Function) int {
	fmt.Println(f, reflect.TypeOf(f), f == nil)
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
	charN, fileN := f.GetSourceN()
	pc.functions[id] = fmt.Sprintf("%v:\t[ %s] in %d, named %v.%v (%v) - %d:%d",
		reflect.TypeOf(f), s, pc.GetIDOf(f.GetParent()),
		name.infinitum, name.bytes, name.Count(), fileN, charN,
	)

	return id
}
