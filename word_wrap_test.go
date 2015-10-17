package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_WordWrap(t *testing.T) {

	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit sil."

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
			lineWidth:    1,
			out:          "Lorem\nipsum\ndolor\nsit\namet,\nconsectetur\nadipiscing\nelit\nsil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    2,
			out:          "Lorem\nipsum\ndolor\nsit\namet,\nconsectetur\nadipiscing\nelit\nsil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    3,
			out:          "Lorem\nipsum\ndolor\nsit\namet,\nconsectetur\nadipiscing\nelit\nsil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    4,
			out:          "Lorem\nipsum\ndolor\nsit\namet,\nconsectetur\nadipiscing\nelit\nsil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    10,
			out:          "Lorem\nipsum\ndolor sit\namet,\nconsectetur\nadipiscing\nelit sil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    20,
			out:          "Lorem ipsum dolor\nsit amet,\nconsectetur\nadipiscing elit sil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    35,
			out:          "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit sil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    47,
			out:          "Lorem ipsum dolor sit amet, consectetur\nadipiscing elit sil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    80,
			out:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit sil.",
		},
		{
			in:           text,
			addLineWidth: true,
			lineWidth:    100,
			out:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit sil.",
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
