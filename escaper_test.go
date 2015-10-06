package stringfy_test

import (
	"testing"

	"github.com/hgsigner/stringfy"
)

func Test_Escape(t *testing.T) {

	tests := []struct {
		in, out string
	}{
		{in: "São França Ávido", out: "Sao Franca Avido"},
	}

	for _, test := range tests {
		es := stringfy.Escape(test.in)
		if es != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, es)
		}
	}
}
