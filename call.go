package main

type Call struct {
	parent Function
	body   Function
	arg    Function
}

func (call *Call) Call(arg Function) Function {
	return RunFunc(RunFunc(call.body, call.arg), arg)
}

func (call *Call) Find(name *Name) Function {
	return call.parent.Find(name)
}

func (call *Call) GetName() *Name {
	return RunFunc(call.body, call.arg).GetName()
}

func (call *Call) AppendCall(f Function) {
	panic("Should not happen")
}

func (call *Call) GetParent() Function {
	panic("Should not happen")
}
