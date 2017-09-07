package main

type Zero struct {
	parent Function
}

func (zero *Zero) Call(name Function) Function {
	//fmt.Println("\nZERO-CALLED:", zero, name, name.GetName())

	return &NegOne{name, zero.parent}
}

func (zero *Zero) Find(name *Name) Function {
	return zero.parent.Find(name)
}

func (zero *Zero) GetName() *Name {
	return &Name{zero.parent, []byte{0}, 0}
}

func (zero *Zero) AppendCall(f Function) {
	//fmt.Println("Something pointless is happening...")
}

func (zero *Zero) GetParent() Function {
	return zero.parent
}

type NegOne struct {
	name   Function
	parent Function
}

func (nOne *NegOne) Call(body Function) Function {
	scope := &Scope{nOne.parent, body, nOne.name.GetName()}

	//fmt.Println("\nNONE-CALLED:", nOne, body, nOne.name.GetName())

	if scope.body == nil {
		panic("Find out what to do")
	}

	return scope
}

func (nOne *NegOne) Find(name *Name) Function {
	return nOne.parent.Find(name)
}

func (nOne *NegOne) GetName() *Name {
	return &Name{nOne.parent, []byte{}, 0}
}

func (nOne *NegOne) AppendCall(f Function) {
	if nOne.name == nil {
		nOne.name = f
	} else {
		panic("Wait, what do I do here again?")
		nOne.name = &Call{nOne, nOne.name, f}
	}
}

func (nOne *NegOne) GetParent() Function {
	return nOne.parent
}
