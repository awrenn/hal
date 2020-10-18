package parser

import (
	"bytes"
)

func finishRefTag(input []byte) (Tag, []byte, error) {
	t := Tag{}

	// We have taken the first chomp of the opening tag, so we know we are getting that
	name, index, input, err := takeUntil(input, REFERENCE_NAME_TERMINATORS)
	if err != nil {
		return t, nil, err
	}
	t.Kind = TagReference
	inner := ReferenceTag{
		Name:       string(name),
		Attributes: nil,
	}
	if bytes.Equal(REFERENCE_NAME_TERMINATORS[index], []byte(REFERENCE_TERM_TOKEN)) {
		t.Inner = inner
		return t, input, nil
	}

	attrs, index, input, err := takeAttributes(input, REFERENCE_TAG_TERMINATORS)
	if err != nil {
		return t, nil, err
	}

	inner.Attributes = attrs
	t.Inner = inner
	return t, input, nil
}
