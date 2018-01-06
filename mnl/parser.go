package mnl

import (
	"fmt"
	"strconv"
	"strings"
)

type metaParser struct {
	lineNumber int
	hasStart   bool

	dictionary map[string]int
	actions    []action
}

func CreateParser() metaParser {
	return metaParser{
		0, false,
		make(map[string]int),
		make([]action, 0),
	}
}

func (mp *metaParser) AddLine(line string) error {
	action, err := mp.ParseLine(line)
	if err != nil {
		return err
	}
	if action != nil {
		if mp.hasStart {
			return fmt.Errorf(
				"Tried adding line to program that has been started",
			)
		}
		mp.actions = append(mp.actions, *action)
		if action.name == "START" {
			mp.hasStart = true
		}
	}
	return nil
}

func (mp *metaParser) ParseLine(line string) (*action, error) {
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
		return nil, fmt.Errorf("[%d] Only one token on line", line)
	} else if len(tokens) == 0 {
		return nil, nil
	}

	neededExpressions := -1
	switch tokens[0] {
	case "DEF":
		if len(tokens) < 3 {
			return nil, fmt.Errorf("[%d] DEF expected number but got EOL", line)
		} else if len(tokens) > 3 {
			return nil, fmt.Errorf("[%d] DEF expected 2 tokens but got more", line)
		}
		n, err := strconv.Atoi(tokens[2])
		if err != nil {
			return nil, fmt.Errorf(
				"[%d] DEF expected number but got \"%s\"",
				mp.lineNumber,
				tokens[2],
			)
		} else if n < -1 {
			return nil,
				fmt.Errorf(
					"[%d] Number %d is lower than -1",
					mp.lineNumber,
					n,
				)
		}

		mp.dictionary[tokens[1]] = n
		return nil, nil
	case "SET":
		neededExpressions = 2
	case "FUN":
		neededExpressions = 2
	case "START":
		neededExpressions = 2
	default:
		return nil, fmt.Errorf(
			"[%d] Unknown starting token \"%s\"",
			mp.lineNumber,
			tokens[0],
		)
	}
	if neededExpressions == -1 {
		panic("Cannot happen")
	}

	expressions, cToken, err := mp.ParseExp(tokens, 1, neededExpressions)
	if err != nil {
		return nil, err
	}
	if cToken != len(tokens) {
		return nil, fmt.Errorf(
			"[%d] Expected a total of %d tokens, but got %d",
			mp.lineNumber,
			cToken,
			len(tokens),
		)
	}

	return &action{tokens[0], expressions}, nil
}

func (mp *metaParser) ParseExp(
	tokens []string,
	offset int,
	count int,
) ([]Expression, int, error) {
	expressions := make([]Expression, count)

	var err error
	var innerExp []Expression

expLoop:
	for currentExp := range expressions {
		if offset >= len(tokens) {
			return nil,
				-1,
				fmt.Errorf(
					"[%d] Expected expression but got EOL",
					mp.lineNumber,
				)
		}

		val, ok := mp.dictionary[tokens[offset]]
		if ok {
			expressions[currentExp] = Number{val}
			offset++
			continue
		}

		val, err = strconv.Atoi(tokens[offset])
		if err == nil {
			if val < -1 {
				return nil,
					-1,
					fmt.Errorf(
						"[%d] Number %d is lower than -1",
						mp.lineNumber,
						val,
					)
			}
			expressions[currentExp] = Number{val}
			offset++
			continue
		}

		legalCallT := []string{"N", "F", "B"}
		legalCallP := []string{"CALL", "ZCALL"}

		for i := 0; i < 6; i++ {
			if tokens[offset] == legalCallP[i/3]+legalCallT[i%3] {
				innerExp, offset, err = mp.ParseExp(tokens, offset+1, 2-i/3)
				if err != nil {
					return nil, -1, err
				}
				expressions[currentExp] = Call{CallType(i), innerExp}
				continue expLoop
			}
		}

		switch tokens[offset] {
		case "ADD":
			if offset+1 >= len(tokens) {
				return nil,
					-1,
					fmt.Errorf(
						"[%d] Expected number but got EOL",
						mp.lineNumber,
					)
			}
			n, err := strconv.Atoi(tokens[offset+1])
			if err != nil {
				return nil,
					-1,
					fmt.Errorf(
						"[%d] ADD expected number but got \"%s\"",
						mp.lineNumber,
						tokens[2],
					)
			} else if n < 0 {
				return nil,
					-1,
					fmt.Errorf(
						"[%d] Number %d is lower than 0",
						mp.lineNumber,
						n,
					)
			}

			innerExp, offset, err = mp.ParseExp(tokens, offset+2, 1)
			if err != nil {
				return nil, -1, err
			}
			expressions[currentExp] = ADD{n, innerExp}
		case "TSET":
			innerExp, offset, err = mp.ParseExp(tokens, offset+1, 3)
			if err != nil {
				return nil, -1, err
			}
			expressions[currentExp] = TSET{false, innerExp}
		case "TFUN":
			innerExp, offset, err = mp.ParseExp(tokens, offset+1, 3)
			if err != nil {
				return nil, -1, err
			}
			expressions[currentExp] = TSET{true, innerExp}
		case "LAM":
			innerExp, offset, err = mp.ParseExp(tokens, offset+1, 2)
			if err != nil {
				return nil, -1, err
			}
			expressions[currentExp] = LAMBDA{innerExp}
		case "ZERO":
			offset += 1
			expressions[currentExp] = ZERO{}
		default:
			return nil, -1, fmt.Errorf(
				"[%d] Unknown token \"%s\"",
				mp.lineNumber,
				tokens[offset],
			)
		}
	}

	return expressions, offset, nil
}

func (mp *metaParser) Convert() (string, error) {
	s := "*"
	for i := range mp.actions {
		s = strings.Replace(s, "*", mp.actions[i].Convert(), -1)
	}
	if mp.hasStart {
		return s, nil
	} else {
		return s, fmt.Errorf("Converted program without a start")
	}
}

func (mp *metaParser) Run() (int, error) {
	e := environment{ZERO{}, make(map[int]Expression)}
	for i := range mp.actions {
		err := mp.actions[i].Call(&e)
		if err != nil {
			return -2, err
		}
	}
	return Count(e.value, &e)
}
