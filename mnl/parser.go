package mnl

import (
	"fmt"
	"strconv"
	"strings"
)

type metaParser struct {
	indent     int
	lineNumber int

	dictionary map[string]int
}

func CreateParser() metaParser {
	return metaParser{
		0,
		0,
		make(map[string]int),
	}
}

func (mp *metaParser) ParseLine(line string) error {
	mp.lineNumber++
	line = strings.ToUpper(strings.Replace(line, "\t", " ", -1))
	split := strings.Split(line, " ")
	var tokens []string
	for i := range split {
		if len(split[i]) > 0 {
			tokens = append(tokens, split[i])
		}
	}

	if len(tokens) == 1 {
		return fmt.Errorf("[%d] Only one token on line", line)
	} else if len(tokens) == 0 {
		return nil
	}

	switch tokens[0] {
	case "DEF":
		if len(tokens) < 3 {
			return fmt.Errorf("[%d] DEF expected number but got EOL", line)
		}
		n, err := strconv.Atoi(tokens[2])
		if err != nil {
			return fmt.Errorf(
				"[%d] DEF expected number but got \"%s\"",
				mp.lineNumber,
				tokens[2],
			)
		}
		mp.dictionary[tokens[1]] = n
	default:
		return fmt.Errorf(
			"[%d] Unknown starting token \"%s\"",
			mp.lineNumber,
			tokens[0],
		)
	}

	return nil
}
