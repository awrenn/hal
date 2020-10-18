package parser

import (
	"io"
	"testing"
)

// TODO this needs way more testing
func TestParserBasic(t *testing.T) {
	inpt := `
        <body>
        <%reference.html /%>
        <complete />
        </body>
    `
	tagKeys := []TagType{TagOpening, TagReference, TagComplete, TagClosing}
	tagReferences := []string{"", "reference.html", "", ""}
	b := []byte(inpt)
	i := 0
	for {
		var tag Tag
		var err error
		tag, b, err = NextTag(b)
		if err == io.EOF {
			break
		}
		if tag.Kind != tagKeys[i] {
			t.Fatalf("Incorrect tag kind parsed: %d vs %d", tag.Kind, tagKeys[i])
		}
		if tag.Kind == TagReference {
			if tag.Inner.Reference() != tagReferences[i] {
				t.Fatal("Tag reference does not match")
			}
		}
		i += 1
	}
}
