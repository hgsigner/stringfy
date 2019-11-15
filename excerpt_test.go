package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

var text string = `Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.

The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.`

func Test_Excerpt(t *testing.T) {

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
		errorOut     string
	}{
		{
			text:       text,
			phrase:     "bhtk",
			addRadious: true,
			radious:    5,
			out:        "",
			errorOut:   "phrase (bhtk) not found",
		},
		{
			text:       text,
			phrase:     "to",
			addRadious: true,
			radious:    5,
			out:        "...rary to popu...",
			errorOut:   "",
		},
		{
			text:       text,
			phrase:     "tra",
			addRadious: true,
			radious:    5,
			out:        "Contrary to...",
			errorOut:   "",
		},
		{
			text:       text,
			phrase:     "Rackham",
			addRadious: true,
			radious:    5,
			out:        "...y H. Rackham.",
			errorOut:   "",
		},
		{
			text:        text,
			phrase:      "sit",
			addRadious:  true,
			radious:     5,
			addOmission: true,
			omission:    "(...)",
			out:         "(...)olor sit amet(...)",
			errorOut:    "",
		},
		{
			text:         text,
			phrase:       "popular",
			addRadious:   true,
			radious:      1,
			addSeparator: true,
			separator:    " ",
			out:          "...to popular belief,...",
			errorOut:     "",
		},
		{
			text:         text,
			phrase:       "to",
			addRadious:   true,
			radious:      2,
			addSeparator: true,
			separator:    " ",
			out:          "Contrary to popular belief,...",
			errorOut:     "",
		},
		{
			text:         text,
			phrase:       "bhtk",
			addRadious:   true,
			radious:      1,
			addSeparator: true,
			separator:    " ",
			out:          "",
			errorOut:     "phrase (bhtk) not found",
		},
		{
			text:         text,
			phrase:       "to",
			addRadious:   true,
			radious:      5,
			addSeparator: true,
			separator:    "-",
			out:          "...rary to popu...",
			errorOut:     "",
		},
		{
			text:         text,
			phrase:       "to popular",
			addRadious:   true,
			radious:      5,
			addSeparator: true,
			separator:    " ",
			out:          "",
			errorOut:     "when composing the excerpt, your phrase should not contain more than one word",
		},
		{
			text:         text,
			phrase:       "to popular",
			addRadious:   true,
			radious:      5,
			addSeparator: true,
			separator:    "-",
			out:          "",
			errorOut:     "when composing the excerpt, your phrase should not contain more than one word",
		},
	}

	for _, test := range tests {
		exc := stringfy.NewExcerpt()

		if test.addRadious {
			exc.Options(stringfy.AddRadius(test.radious))
		}

		if test.addOmission {
			exc.Options(stringfy.AddOmission(test.omission))
		}

		if test.addSeparator {
			exc.Options(stringfy.AddSeparator(test.separator))
		}

		excf, err := exc.Perform(test.text, test.phrase)
		if err != nil {
			if err.Error() != test.errorOut {
				t.Errorf("\nError Expected: %s\nError Got:      %s", test.errorOut, err.Error())
			}
		}

		if excf != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, excf)
		}
	}
}

func BenchmarkExcerptPerform_Radious(b *testing.B) {
	exc := stringfy.NewExcerpt()
	exc.Options(stringfy.AddRadius(5))
	for i := 0; i < b.N; i++ {
		exc.Perform(text, "to")
	}
}

func BenchmarkExcerptPerform_RadiosSeparator(b *testing.B) {
	exc := stringfy.NewExcerpt()
	exc.Options(stringfy.AddRadius(1), stringfy.AddSeparator(" "))
	for i := 0; i < b.N; i++ {
		exc.Perform(text, "popular")
	}
}

func BenchmarkExcerptPerform_RadiosOmission(b *testing.B) {
	exc := stringfy.NewExcerpt()
	exc.Options(stringfy.AddRadius(5), stringfy.AddOmission("(...)"))
	for i := 0; i < b.N; i++ {
		exc.Perform(text, "sit")
	}
}
