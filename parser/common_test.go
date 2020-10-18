package parser

import (
	"bytes"
	"testing"
)

func TestTakeUntil(t *testing.T) {
	var testCases = []struct {
		text   []byte
		terms  [][]byte
		result []byte
		input  []byte
		index  int
		err    error
	}{
		{
			text:   []byte("Hello, I am text!"),
			terms:  [][]byte{[]byte(",")},
			result: []byte("Hello"),
			input:  []byte(" I am text!"),
			index:  0,
			err:    nil,
		},
		{
			text:   []byte("Hello, I am text!"),
			terms:  [][]byte{[]byte(" "), []byte(", I"), []byte(",")},
			result: []byte("Hello"),
			input:  []byte(" am text!"),
			index:  1,
			err:    nil,
		},
		{
			text:   []byte("Hello, I am text!"),
			terms:  [][]byte{[]byte("Z")},
			result: nil,
			input:  nil,
			index:  0,
			err:    NoTerminator,
		},
	}

	for i, tc := range testCases {
		result, index, input, err := takeUntil(tc.text, tc.terms)
		if !bytes.Equal(result, tc.result) {
			t.Fatalf("Results do not match in case %d", i)
		}
		if !bytes.Equal(input, tc.input) {
			t.Fatalf("Inputs do not match in case %d", i)
		}
		if index != tc.index {
			t.Fatalf("Indexes do not match in case %d", i)
		}
		if err != tc.err {
			t.Fatalf("Errors do not match in case %d", i)
		}
	}
}
