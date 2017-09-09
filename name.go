package main

type Name struct {
	parent    Function
	bytes     []byte
	infinitum uint64

	charN, fileN int
}

func (name *Name) GetSourceN() (int, int) {
	if name == nil {
		return 0, 0
	}
	return name.charN, name.fileN
}

func (name *Name) AddOne() {
	if len(name.bytes) == 0 {
		name.bytes = []byte{0}
		return
	}
	for i := range name.bytes {
		if (name.bytes)[i] == 255 {
			(name.bytes)[i] = 0
			if i+1 >= len(name.bytes) {
				name.bytes = append(name.bytes, 1)
				break
			}
		} else {
			name.bytes[i]++
			break
		}
	}
}

func (name *Name) SubOne() {
	if len(name.bytes) == 0 {
		name.bytes = []byte{1}
		return
	}
	for i := range name.bytes {
		if name.bytes[i] == 0 {
			name.bytes[i] = 255
			if i+1 >= len(name.bytes) {
				name.bytes = []byte{}
				break
			}
		} else {
			name.bytes[i]--
			break
		}
	}
}

func (name *Name) Equals(other *Name) bool {
	if name == nil {
		panic("Name is nil!")
	} else if other == nil {
		panic("Other is nil!")
	}
	if name.infinitum != other.infinitum {
		return false
	}

	for i := range name.bytes {
		if i < len(other.bytes) {
			if name.bytes[i] != other.bytes[i] {
				return false
			}
		} else if name.bytes[i] != 0 {
			return false
		}
	}
	for i := len(name.bytes); i < len(other.bytes); i++ {
		if other.bytes[i] != 0 {
			return false
		}
	}
	return true
}

func (name *Name) Call(f Function) Function {
	c := name.Count()
	if c == 0 {
		return (&Zero{name.parent, charN, fileN}).Call(f)
	} else if c == -1 {
		return (&NegOne{nil, name.parent, charN, fileN}).Call(f)
	}
	n := name.GetName()
	n.SubOne()
	return n
}

func (name *Name) Find(n *Name) Function {
	panic("Should not happen")
	//	return name.parent.Find(n)
}

func (name *Name) GetName() *Name {
	if name == nil {
		panic("Name is nil")
	}
	bytes := make([]byte, len(name.bytes))
	copy(bytes, name.bytes)
	return &Name{name.parent, bytes, name.infinitum, charN, fileN}
}

func (name *Name) AppendCall(f Function) {
	panic("Figure out how to do this")
}

func (name *Name) GetParent() Function {
	if name == nil {
		return nil
	}
	return name.parent
}

func (name *Name) Count() int {
	if name == nil {
		return -2
	}
	if len(name.bytes) == 0 {
		return -1
	}
	s := 0
	for i := range name.bytes {
		s += int(name.bytes[i])
	}
	return s
}

func (name *Name) GetArgs() ([]Function, []string) {
	return []Function{}, []string{}
}
