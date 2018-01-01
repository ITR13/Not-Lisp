package mnl

import (
	"testing"
)

func TestParseLine_DEF_ERROR(t *testing.T) {
	tests := [][2]interface{}{
		{"DEF", true},
		{"DEF CAT", true},
		{"DEF CAT 123", false},
		{"DEF CAT 123A123", true},
	}

	metaParser := CreateParser()
	var err error

	for i := range tests {
		err = metaParser.ParseLine(tests[i][0].(string))
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
