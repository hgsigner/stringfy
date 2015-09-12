package escaper_test

import (
	"testing"

	"github.com/hgsigner/stringfy/escaper"
)

func Test_PerformOn(t *testing.T) {

	tests := []struct {
		in, out string
	}{
		{in: "São França Ávido", out: "Sao Franca Avido"},
	}

	for _, test := range tests {
		es := escaper.PerformOn(test.in)
		if es != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, es)
		}
	}
}
