package mnl

import (
	"strings"
	"testing"

	"github.com/ITR13/Not-Lisp/interpreter"
)

func TestParseLine_ERROR(t *testing.T) {
	tests := [][2]interface{}{
		{"DEF", true},
		{"DEF CAT", true},
		{"DEF CAT 123", false},
		{"DEF CAT 123A123", true},
		{"DEF CAT 123 123", true},
		{"DEF CAT -1", false},
		{"DEF CAT -2", true},

		{"SET", true},
		{"SET CAT", true},
		{"SET CAT 123", false},
		{"SET CAT 123A123", true},
		{"SET DOG CAT", true},
		{"SET 123 CAT", false},
		{"SET 123 123", false},
		{"SET CAT CAT", false},
		{"SET CAT CAT CAT", true},
		{"SET -1 -1", false},
		{"SET -1 -2", true},
		{"SET -2 -1", true},
		{"SET -2 -2", true},

		{"SET CALLN", true},
		{"SET CALLN 1", true},
		{"SET CALLN 1 1", true},
		{"SET CALLN 1 1 1", false},
		{"SET CALLN 1 1 1 1", true},
		{"SET CALLF", true},
		{"SET CALLF 1", true},
		{"SET CALLF 1 1", true},
		{"SET CALLF 1 1 1", false},
		{"SET CALLF 1 1 1 1", true},
		{"SET CALLB", true},
		{"SET CALLB 1", true},
		{"SET CALLB 1 1", true},
		{"SET CALLB 1 1 1", false},
		{"SET CALLB 1 1 1 1", true},

		{"SET ZCALLN", true},
		{"SET ZCALLN 1", true},
		{"SET ZCALLN 1 1", false},
		{"SET ZCALLN 1 1 1", true},
		{"SET ZCALLN 1 1 1 1", true},
		{"SET ZCALLF", true},
		{"SET ZCALLF 1", true},
		{"SET ZCALLF 1 1", false},
		{"SET ZCALLF 1 1 1", true},
		{"SET ZCALLF 1 1 1 1", true},
		{"SET ZCALLB", true},
		{"SET ZCALLB 1", true},
		{"SET ZCALLB 1 1", false},
		{"SET ZCALLB 1 1 1", true},
		{"SET ZCALLB 1 1 1 1", true},

		{"FUN", true},
		{"FUN CAT", true},
		{"FUN CAT 123", false},
		{"FUN CAT 123A123", true},
		{"FUN DOG CAT", true},
		{"FUN 123 CAT", false},
		{"FUN 123 123", false},
		{"FUN CAT CAT", false},
		{"FUN CAT CAT CAT", true},

		{"START", true},
		{"START CAT", true},
		{"START CAT 123A123", true},
		{"START DOG CAT", true},
		{"START CAT CAT CAT", true},
		{"START 123 123", false},
		{"START 123 123", true},
	}

	metaParser := CreateParser()
	var err error

	for i := range tests {
		err = metaParser.AddLine(tests[i][0].(string))
		if (err != nil) != tests[i][1].(bool) {
			t.Errorf(
				"Expected %v for \"%s\", but got \"%v\"",
				tests[i][1].(bool),
				tests[i][0].(string),
				err,
			)
		}
	}
}

func TestSimple(t *testing.T) {
	tests := [][2]interface{}{
		{`
			SET 5 2
			START 5 0
		`, 2}, {`
			FUN 5 2
			START 5 0
		`, 3}, {`
			SET 5 2
			SET 6 5
			START 6 0
		`, 5},

		{`
			FUN 5 TSET 5 2 ZCALLN 5
			START ZCALLN 5 0
		`, 2}, {`
			FUN 5 TSET 5 2 CALLN 5 0
			START CALLN 5 0 0
		`, 2},

		{`
			DEF RETSAME 10
			DEF M 11
			
			SET RETSAME LAM M ZCALLN M			
			START RETSAME 20
		`, 2},

		{`
DEF V 4
DEF ADDFUNC 5
DEF M 6

SET ADDFUNC LAM M TSET V ADD 1 ZCALLN V ZCALLF M
SET V 0

START LAM ZERO CALLF ADDFUNC LAM ZERO CALLF ADDFUNC LAM ZERO ZCALLN V ZERO
		`, 2},

		/*{`
		DEF ADDFUNC 8
		DEF REPEAT 9
		DEF CALLADD 10
		DEF HEAL 11
		DEF V 12
		DEF M 13

		SET V 0
		SET ADDFUNC LAM M TSET V ADD 1 V ZCALLF M
		FUN CALLADD CALLF ADDFUNC ZCALLN REPEAT

		SET HEAL LAM M TSET 6 5 TSET ADD 1 ZCALLN V ZCALLN V ZCALLF M
		FUN REPEAT TSET 6 ZCALLN CALLADD TSET ADD 1 ZCALLN V V ZCALLF 6

		START ZCALLN REPEAT ZERO
				`, 5},*/
	}

testLoop:
	for i := range tests {
		mp := CreateParser()
		lines := strings.Split(tests[i][0].(string), "\n")
		for j := range lines {
			err := mp.AddLine(lines[j])
			if err != nil {
				t.Errorf("Test %d Line %d gave an error:\n%v", i, j, err)
				continue testLoop
			}
		}
		wanted := tests[i][1].(int)

		v, err := mp.Run(EnvSettings{false, false})
		if err != nil {
			t.Errorf("Test %d [R] gave an error:\n%v", i, err)
			continue testLoop
		}
		if v != wanted {
			t.Errorf("Test %d [R] gave %d but wanted %d", i, v, wanted)
		}

		s, _ := mp.Convert()
		v = interpreter.RunForInt(s)
		if v != wanted {
			t.Errorf("Test %d [C] gave %d but wanted %d:\n%s", i, v, wanted, s)
		}
	}
}
