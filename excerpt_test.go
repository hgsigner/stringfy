package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_Excerpt(t *testing.T) {

	text1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	tests := []struct {
		text         string
		phrase       string
		radious      int
		addRadious   bool
		omission     string
		addOmission  bool
		separator    string
		addSeparator bool
		out          string
	}{
		{
			text:       text1,
			phrase:     "boo",
			addRadious: true,
			radious:    5,
			out:        text1,
		},
		{
			text:       text1,
			phrase:     "sit",
			addRadious: true,
			radious:    5,
			out:        "...olor sit amet...",
		},
		{
			text:       text1,
			phrase:     "em",
			addRadious: true,
			radious:    5,
			out:        "Lorem ipsu...",
		},
		{
			text:       text1,
			phrase:     "elit",
			addRadious: true,
			radious:    5,
			out:        "...cing elit.",
		},
		{
			text:        text1,
			phrase:      "sit",
			addRadious:  true,
			radious:     5,
			addOmission: true,
			omission:    "(...)",
			out:         "(...)olor sit amet(...)",
		},
		{
			text:         text1,
			phrase:       "dolor",
			addRadious:   true,
			radious:      1,
			addSeparator: true,
			separator:    " ",
			out:          "...ipsum dolor sit...",
		},
	}

	for _, test := range tests {
		exc := stringfy.NewExcerpt()

		if test.addRadious {
			exc.Options(stringfy.AddRadious(test.radious))
		}

		if test.addOmission {
			exc.Options(stringfy.AddOmission(test.omission))
		}

		if test.addSeparator {
			exc.Options(stringfy.AddSeparator(test.separator))
		}

		excf := exc.Perform(test.text, test.phrase)
		if excf != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, excf)
		}
	}
}
