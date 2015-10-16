package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_WordWrap(t *testing.T) {

	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	tests := []struct {
		in           string
		addLineWidth bool
		lineWidth    int
		out          string
	}{
		{
			in:  text,
			out: text,
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    8,
			out:          "Lorem\nipsum\ndolor\nsit\namet,\nconsectetur\nadipiscing\nelit.",
		},
	}

	for _, test := range tests {
		wwr := stringfy.NewWordWrap()
		if test.addLineWidth {
			wwr.Options(stringfy.AddLineWidth(test.lineWidth))
		}
		wwrf := wwr.Perform(test.in)
		if wwrf != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, wwrf)
		}
	}
}
