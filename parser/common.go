package parser

import (
	"bytes"

	"github.com/joomcode/errorx"
)

var (
	SyntaxError    = errorx.CommonErrors.NewType("Synxtax Error")
	ReferenceCycle = errorx.CommonErrors.NewType("Reference Cycle Detected")
	NoTerminator   = SyntaxError.New("Terminator never arrived")
)

func takeAttributes(input []byte, terminators [][]byte) (map[string][]string, int, []byte, error) {
	result, index, input, err := takeUntil(input, terminators)
	if err != nil {
		return nil, 0, nil, err
	}
	_ = result

	// TODO acwrenn - actually grab tag attributes from result -> rules are simple
	// https://html.spec.whatwg.org/multipage/syntax.html
	return make(map[string][]string), index, input, nil
}

func takeUntil(input []byte, terminators [][]byte) ([]byte, int, []byte, error) {
	result, index, input, err := takeUpTo(input, terminators)
	if err != nil {
		return nil, 0, nil, err
	}
	return result, index, input[len(terminators[index]):], nil
}

func takeUpTo(input []byte, terminators [][]byte) ([]byte, int, []byte, error) {
	result := make([]byte, 0)
	for {
		if len(input) == 0 {
			return nil, 0, nil, NoTerminator
		}
		for index := range terminators {
			// We have eaten input until 1 of the terminators arrived
			if bytes.HasPrefix(input, terminators[index]) {
				return result, index, input, nil
			}
		}
		result = append(result, input[0])
		input = input[1:]
	}
}
