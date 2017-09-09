package main

import (
	"testing"
)

func TestInterpretString(t *testing.T) {
	tests := [][2]interface{}{
		{"", 0},
		{"()", 1},
		{"(())", 2},
		{"((()))", 3},
		{"(((())))", 4},
		{"(())()", 1},
		{"((()))()", 2},
		{"(((())))()", 3},
		{"((((()))))()", 4},
		{"((()))()()", 1},
		{"(((())))()()", 2},
		{"((((()))))()()", 3},

		{"()()", 0},
		{"()(())", 0},
		{"(())(())", 1},
		{"(())((()))", 1},

		{"()()()", -1},
		{"()()(())", -1},
		{"()()((()))", -1},
		{"()(())()", -1},
		{"()(())(())", -1},

		{"()()()()", 1},
		{"()()()()()", 0},
		{"(()()()())", 2},

		{"(()())", 1},
		{"((()()))", 2},
		{"((())())", 2},
		{"((()))()", 2},

		{"()()()(())()", 1},
		{"()()()((()))()", 2},
		{"()()()(((())))()", 3},
		{"()()()(())(())", 1},
		{"()()()((()))(())", 2},
		{"()()()(((())))(())", 3},
		{"()()()(())((()))", 1},
		{"()()()((()))((()))", 2},
		{"()()()(((())))((()))", 3},
	}
	for _, pair := range tests {
		p, r := pair[0].(string), pair[1].(int)
		mainScope := &MainScope{}
		scope := InterpretString(p, mainScope)
		ex := scope.GetName()
		c := ex.Count()
		if r != c {
			t.Errorf("%s: Expected %d but got %d ( %v )", p, r, c, ex.bytes)
		} else {
			t.Logf("%s: Succeeded with %d ( %v )", p, c, ex.bytes)
		}
	}
}

func TestAddOne(t *testing.T) {
	s := []byte{255, 0}
	n := (&Name{nil, s, 0, charN, fileN})
	n.AddOne()
	s = n.bytes
	if len(s) != 2 || s[1] != 1 || s[0] != 0 {
		t.Fatal("Expected [1,0], got ", s)
	}
	n = (&Name{nil, s, 0, charN, fileN})
	n.SubOne()
	s = n.bytes
	if len(s) != 2 || s[0] != 255 {
		t.Fatal("Expected [255], got ", s)
	}
	n = (&Name{nil, s, 0, charN, fileN})
	n.AddOne()
	s = n.bytes
	if len(s) != 2 || s[1] != 1 || s[0] != 0 {
		t.Fatal("Expected [1,0], got ", s)
	}

	s = []byte{255, 255, 0, 255, 255}
	n = (&Name{nil, s, 0, charN, fileN})
	n.AddOne()
	s = n.bytes
	if len(s) != 5 {
		t.Fatal("Expected [255,255,1,0,0], got ", s)
	}
	for i := range s {
		if i < 2 {
			if s[i] != 0 {
				t.Fatal("Expected [255,255,1,0,0], got ", s)
			}
		} else if i < 3 {
			if s[i] != 1 {
				t.Fatal("Expected [255,255,1,0,0], got ", s)
			}
		} else if s[i] != 255 {
			t.Fatal("Expected [255,255,1,0,0], got ", s)
		}
	}

	s = []byte{1, 0, 0, 0, 0}
	n = (&Name{nil, s, 0, charN, fileN})
	n.SubOne()
	s = n.bytes
	if len(s) != 5 {
		t.Fatal("Expected [0,0,0,0,0], got ", s)
	}
	for i := range s {
		if s[i] != 0 {
			t.Fatal("Expected [0,0,0,0,0], got ", s)
		}
	}
	n = (&Name{nil, s, 0, charN, fileN})
	n.SubOne()
	s = n.bytes
	if len(s) != 0 {
		t.Fatal("Expected [], got ", s)
	}
	n = (&Name{nil, s, 0, charN, fileN})
	n.SubOne()
	s = n.bytes
	if len(s) != 1 || s[0] != 1 {
		t.Fatal("Expected [1], got ", s)
	}
}

func TestAdvancedMethods(t *testing.T) {
	tests := [][2]interface{}{
		{"()()(())(()())()", 0},
		{"()()(())(()())(())", 1},
		{"()()(())(()())((()))", 2},
		{"()()(())(()())(((())))", 3},
		{"()()()(())()()", 0},
		{"()()()(())(())()", 0},
		{"()()()(())((()))()", 0},
	}
	for _, pair := range tests {
		p, r := pair[0].(string), pair[1].(int)
		mainScope := &MainScope{}
		scope := InterpretString(p, mainScope)
		ex := scope.GetName()
		c := ex.Count()
		if r != c {
			t.Errorf("%s: Expected %d but got %d ( %v )", p, r, c, ex.bytes)
		} else {
			t.Logf("%s: Succeeded with %d ( %v )", p, c, ex.bytes)
		}
	}
}
