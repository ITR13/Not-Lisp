package main

import "fmt"

type Zero struct {
	parent       Function
	charN, fileN int
}

func (zero *Zero) GetSourceN() (int, int) {
	return zero.charN, zero.fileN
}

func (zero *Zero) Call(name Function) Function {
	if name != nil {
		fmt.Println("\nZERO-CALLED:", zero, name, name.GetName().bytes)
	} else {
		fmt.Println("\nZERO-CALLED:", zero, name, nil)
	}

	return &NegOne{name, zero.parent, charN, fileN}
}

func (zero *Zero) Find(name *Name) Function {
	panic("Should not happen")
	//return zero.parent.Find(name)
}

func (zero *Zero) GetName() *Name {
	return &Name{zero.parent, []byte{0}, 0, charN, fileN}
}

func (zero *Zero) AppendCall(f Function) {
	//fmt.Println("Something pointless is happening...")
}

func (zero *Zero) GetParent() Function {
	return zero.parent
}

func (zero *Zero) GetArgs() ([]Function, []string) {
	return []Function{}, []string{}
}

type NegOne struct {
	name         Function
	parent       Function
	charN, fileN int
}

func (nOne *NegOne) GetSourceN() (int, int) {
	return nOne.charN, nOne.fileN
}

func (nOne *NegOne) Call(body Function) Function {
	var name *Name
	if nOne.name != nil {
		name = nOne.name.GetName()
	}
	scope := &Scope{nOne.parent, body, name, charN, fileN}
	fmt.Println("NOne", body, name)

	return scope
}

func (nOne *NegOne) Find(name *Name) Function {
	panic("Should not happen")
	//return nOne.parent.Find(name)
}

func (nOne *NegOne) GetName() *Name {
	return &Name{nOne.parent, []byte{}, 0, charN, fileN}
}

func (nOne *NegOne) AppendCall(f Function) {
	if nOne.name == nil {
		nOne.name = f
	} else {
		panic("Wait, what do I do here again?")
		nOne.name = &Call{nOne, nOne.name, f, charN, fileN}
	}
}

func (nOne *NegOne) GetParent() Function {
	return nOne.parent
}

func (nOne *NegOne) GetArgs() ([]Function, []string) {
	return []Function{nOne.name}, []string{"Name"}
}
