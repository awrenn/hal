package parser

const (
	REGULAR_OPEN_TOKEN  = "<"
	REGULAR_CLOSE_TOKEN = "</"
	REGULAR_TERM_TOKEN  = ">"
	COMPLETE_TERM_TOKEN = "/>"

	REFERENCE_OPEN_TOKEN = "<%"
	REFERENCE_TERM_TOKEN = "/%>"

	WHITESPACE = " \n\t"
	SPACES     = " \t"
)

var (
	// This list needs to match precedence
	// COMPLETE_TERM_TOKEN must come before REGULAR_TERM_TOKEN
	REGULAR_TAG_TERMINATORS = [][]byte{
		[]byte(COMPLETE_TERM_TOKEN),
		[]byte(REGULAR_TERM_TOKEN),
	}
	REFERENCE_TAG_TERMINATORS = [][]byte{
		[]byte(REFERENCE_TERM_TOKEN),
	}
	WHITE_SPACE_TERMINATORS = [][]byte{
		[]byte(" "),
		[]byte("\t"),
		[]byte("\n"),
	}

	REGULAR_NAME_TERMINATORS = [][]byte{}

	REFERENCE_NAME_TERMINATORS = [][]byte{}

	// Precendence matters again - shortest needs to come last
	TEXT_TERMINATORS = [][]byte{
		[]byte(REFERENCE_OPEN_TOKEN),
		[]byte(REGULAR_CLOSE_TOKEN),
		[]byte(REGULAR_OPEN_TOKEN),
	}
)

func init() {
	REGULAR_NAME_TERMINATORS = append(
		WHITE_SPACE_TERMINATORS,
		REGULAR_TAG_TERMINATORS...,
	)

	REFERENCE_NAME_TERMINATORS = append(
		WHITE_SPACE_TERMINATORS,
		REFERENCE_TAG_TERMINATORS...,
	)
}
