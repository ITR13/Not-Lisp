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

type DebugFunction struct {
	parent       Function
	charN, fileN int
}

func (debug *DebugFunction) AppendCall(Function) {
	panic("Cannot happen")
}

func (debug *DebugFunction) Call(f Function) Function {
	charN, fileN := -1, -1
	var name *Name
	if f != nil {
		charN, fileN = f.GetSourceN()
		name = f.GetName()
	}
	fmt.Printf("DebugFunction at %d:%d called with %d:%d (%v)\n",
		debug.fileN, debug.charN, charN, fileN, name)
	return debug
}

func (debug *DebugFunction) GetArgs() ([]Function, []string) {
	return []Function{}, []string{}
}

func (debug *DebugFunction) GetName() *Name {
	fmt.Printf("Got name from DebugFunction at %d:%d\n",
		debug.charN, debug.fileN)
	return &Name{debug.parent, []byte{1}, 1, charN, fileN}
}

func (debug *DebugFunction) GetParent() Function {
	fmt.Printf("Got parent from DebugFunction at %d:%d\n",
		debug.charN, debug.fileN)
	return debug.parent
}

func (debug *DebugFunction) GetSourceN() (int, int) {
	return debug.charN, debug.fileN
}

func (debug *DebugFunction) Resolve() Function {
	fmt.Printf("DebugFunction at %d:%d resolved\n", debug.fileN, debug.charN)
	return debug
}
