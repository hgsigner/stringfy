package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_Truncate(t *testing.T) {

	tests := []struct {
		in           string
		length       int
		addLength    bool
		omission     string
		addOmission  bool
		separator    string
		addSeparator bool
		out          string
	}{
		{
			in:  "Lorem ipsum.",
			out: "Lorem ipsum.",
		},
		{
			in:        "Lor",
			addLength: true,
			length:    3,
			out:       "Lor",
		},
		{
			in:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			out: "Lorem ipsum dolor sit amet,...",
		},
		{
			in:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			addOmission: true,
			omission:    "... (continued)",
			out:         "Lorem ipsum dol... (continued)",
		},
		{
			in:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			addLength:    true,
			length:       13,
			addSeparator: true,
			separator:    " ",
			out:          "Lorem ipsum...",
		},

		{
			in:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			addLength:    true,
			length:       13,
			addSeparator: true,
			separator:    " ",
			addOmission:  true,
			omission:     "... (continued) ...",
			out:          "Lorem ipsum... (continued) ...",
		},
		{
			in:           "Lorem ipsum dolor sit amet, c",
			addLength:    true,
			length:       13,
			addSeparator: true,
			separator:    " ",
			addOmission:  true,
			omission:     "... (continued) ...",
			out:          "Lorem ipsum dolor sit amet, c",
		},
	}

	for _, test := range tests {
		tr := stringfy.NewTruncate()

		if test.addLength {
			tr.Options(stringfy.AddLength(test.length))
		}

		if test.addOmission {
			tr.Options(stringfy.AddOmission(test.omission))
		}

		if test.addSeparator {
			tr.Options(stringfy.AddSeparator(test.separator))
		}

		trf := tr.Perform(test.in)
		if trf != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, trf)
		}
	}

}
