package parser

import (
	"bytes"
	"io"
)

func NextTag(input []byte) (Tag, []byte, error) {
	t := Tag{}
	input = eatWhitespace(input)
	if len(input) == 0 {
		return t, input, io.EOF
	}
	tagType, output, err := startTag(input)
	if err != nil {
		return t, nil, SyntaxError.Wrap(err, "Error parsing tag input")
	}

	switch tagType {
	case TagOpening:
		t, output, err = finishOpeningTag(output)
	case TagText:
		t, output, err = finishTextTag(output)
	case TagReference:
		t, output, err = finishRefTag(output)
	case TagClosing:
		t, output, err = finishClosingTag(output)
	default:
		panic("Error; invalid tag type. Curse you golang")
	}
	t.Raw = input[:len(input)-len(output)]
	return t, output, err
}

func eatWhitespace(input []byte) []byte {
	return bytes.TrimLeft(input, WHITESPACE)
}

func eatSpaces(input []byte) []byte {
	return bytes.TrimLeft(input, SPACES)
}

func startTag(input []byte) (TagType, []byte, error) {
	// Reference needs to take precedence over regular
	if bytes.HasPrefix(input, []byte(REFERENCE_OPEN_TOKEN)) {
		return TagReference, input[len(REFERENCE_OPEN_TOKEN):], nil
	}
	// Same here - REGULAR_OPEN needs to come last
	if bytes.HasPrefix(input, []byte(REGULAR_CLOSE_TOKEN)) {
		return TagClosing, input[len(REGULAR_CLOSE_TOKEN):], nil
	}

	// This tag might eventually become a TagComplete, depending on how the tag ends
	if bytes.HasPrefix(input, []byte(REGULAR_OPEN_TOKEN)) {
		return TagOpening, input[len(REGULAR_OPEN_TOKEN):], nil
	}

	// For now, just eat all the text until the next tag...
	return TagText, input, nil
}
