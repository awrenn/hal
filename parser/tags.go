package parser

import (
	"bytes"
	"io"
	"strings"
)

type TagType int

const (
	// An Opening Tag is the opening tag for a partent node
	// <body> is an example of an opening tag
	TagOpening TagType = iota
	// A complete tag is a self-contained HTML tag
	// <script src="" /> is an example
	TagComplete
	// Just text:
	// <div> <- TagOpening
	//  Hello World! <- TagText
	// </div> <- TagClosing
	// We aren't building a tree, so this text isn't just a weird child
	TagText
	// A special reference Tag, that means we need to switch to parsing a different file until EOF
	TagReference
	// A closing tag of an opening tag. </div>
	TagClosing
)

type Tag struct {
	Kind  TagType
	Raw   []byte
	Inner TagInner
}

type TagInner interface {
	// Return the file being referenced
	// If not a RegularTag, panics
	Reference() string
}

type RegularTag struct {
	Name       string
	Attributes map[string][]string
}

func (r RegularTag) Reference() string {
	panic("Not ref tag")
}

type TextTag struct {
	Text string
}

func (t TextTag) Reference() string {
	panic("Not ref tag")
}

type ReferenceTag struct {
	Name       string
	Attributes map[string][]string
}

func (r ReferenceTag) Reference() string {
	return r.Name
}

type ClosingTag struct {
	Name string
}

func (c ClosingTag) Reference() string {
	panic("Not ref tag")
}

func (t Tag) Reader(indentCount int) io.Reader {
	if indentCount < 0 {
		indentCount = 0
	}
	buf := bytes.NewBuffer(make([]byte, 0, indentCount+len(t.Raw)+1))
	buf.WriteString(strings.Repeat("\t", indentCount))
	buf.Write(t.Raw)
	// TextTags bring their own terminal new lines
	// Reference tags end with their own newline as well
	if t.Kind != TagText && t.Kind != TagReference {
		buf.WriteString("\n")
	}
	return buf
}
