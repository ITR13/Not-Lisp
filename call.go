package main

type Call struct {
	parent       Function
	body         Function
	arg          Function
	charN, fileN int
}

func (call *Call) GetSourceN() (int, int) {
	return call.charN, call.fileN
}

func (call *Call) Call(arg Function) Function {
	return RunFunc(RunFunc(call.body, call.arg), arg)
}

func (call *Call) Find(name *Name) Function {
	panic("Should not happen")
	//	return call.parent.Find(name)
}

func (call *Call) GetName() *Name {
	return call.Resolve().GetName()
	//return RunFunc(call.body, call.arg).GetName()
	//name := call.body.GetName()
	//name.SubOne()
	//return name
	//return call.body.Call(call.arg).GetName()
}

func (call *Call) AppendCall(f Function) {
	panic("Should not happen")
}

func (call *Call) GetParent() Function {
	//panic("Should not happen")
	return nil
}
func (call *Call) GetArgs() ([]Function, []string) {
	return []Function{call.body, call.arg}, []string{"Body", "Args"}
}

func (call *Call) Resolve() Function {
	return RunFunc(call.body, call.arg)
}
