package parser

import (
	"bytes"
)

func finishClosingTag(input []byte) (Tag, []byte, error) {
	t := Tag{}
	name, index, input, err := takeUntil(input, REGULAR_NAME_TERMINATORS)
	if err != nil {
		return t, nil, err
	}
	if len(name) == 0 {
		return t, nil, SyntaxError.New("Whitespace before name after closing tag")
	}
	t.Kind = TagClosing
	inner := ClosingTag{
		Name: string(name),
	}
	t.Inner = inner

	if !bytes.Equal(REGULAR_NAME_TERMINATORS[index], []byte(REGULAR_TERM_TOKEN)) {
		// Make sure to take until a real tag terminator
		var out []byte
		out, _, input, err = takeUntil(input, REGULAR_TAG_TERMINATORS)
		// Make sure there are no extra characters after the name
		// Like -> </body foobar>
		// </body    > -> This should pass
		out = eatSpaces(out)
		if len(out) != 0 {
			return t, nil, SyntaxError.New("Extra characters between closing tag name and closing tag")
		}
	}

	input = eatWhitespace(input)
	return t, input, nil
}
