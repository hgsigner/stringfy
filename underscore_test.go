package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_Undercore(t *testing.T) {

	tests := []struct {
		in, out string
	}{
		{in: "    Fizz    ", out: "fizz"},
		{in: "Fizz", out: "fizz"},
		{in: "Fizz Buzz", out: "fizz_buzz"},
		{in: "fizz_buzz", out: "fizz_buzz"},
		{in: "Fázz Bôzz", out: "fazz_bozz"},
		{in: "São Paulo", out: "sao_paulo"},
		{in: "São França Ávido", out: "sao_franca_avido"},
	}

	for _, test := range tests {
		und := stringfy.Underscore(test.in)
		if und != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, und)
		}
	}

}
