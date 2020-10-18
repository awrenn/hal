package parser

import ()

func finishTextTag(input []byte) (Tag, []byte, error) {
	t := Tag{}
	t.Kind = TagText
	text, _, input, err := takeUpTo(input, TEXT_TERMINATORS)
	if err != nil {
		return t, nil, SyntaxError.Wrap(err, "Text never terminates")
	}

	t.Inner = TextTag{
		Text: string(text),
	}
	return t, input, nil
}
