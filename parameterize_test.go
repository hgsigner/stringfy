package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_Parameterize(t *testing.T) {

	tests := []struct {
		in, out string
	}{
		{in: "    Fizz    ", out: "fizz"},
		{in: "Fizz", out: "fizz"},
		{in: "Fizz Buzz", out: "fizz-buzz"},
		{in: "fizz-buzz", out: "fizz-buzz"},
		{in: "Fázz Bôzz", out: "fazz-bozz"},
		{in: "São Paulo", out: "sao-paulo"},
		{in: "São França Ávido", out: "sao-franca-avido"},
	}

	for _, test := range tests {
		par := stringfy.Parameterize(test.in)
		if par != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, par)
		}
	}

}
