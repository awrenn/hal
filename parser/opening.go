package parser

import (
	"bytes"
)

func finishOpeningTag(input []byte) (Tag, []byte, error) {
	t := Tag{}
	// We have taken the first chomp of the opening tag, so we know we are getting that
	name, index, input, err := takeUntil(input, REGULAR_NAME_TERMINATORS)
	if err != nil {
		return t, nil, err
	}
	t.Kind = TagOpening
	inner := RegularTag{
		Name:       string(name),
		Attributes: nil,
	}

	// TODO acwrenn
	// Make this faster, without doing constant int matching...?
	if bytes.Equal(REGULAR_NAME_TERMINATORS[index], []byte(COMPLETE_TERM_TOKEN)) {
		t.Kind = TagClosing
		t.Inner = inner
		return t, input, nil
	} else if bytes.Equal(REGULAR_NAME_TERMINATORS[index], []byte(REGULAR_TERM_TOKEN)) {
		t.Inner = inner
		return t, input, nil
	}

	attrs, index, input, err := takeAttributes(input, REGULAR_TAG_TERMINATORS)
	if err != nil {
		return t, nil, err
	}

	inner.Attributes = attrs
	if bytes.Equal(REGULAR_TAG_TERMINATORS[index], []byte(COMPLETE_TERM_TOKEN)) {
		t.Kind = TagComplete
	}
	t.Inner = inner
	return t, input, nil
}
