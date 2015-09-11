package parameterize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parameterize(t *testing.T) {
	a := assert.New(t)

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

	for _, t := range tests {
		a.Equal(t.out, PerformOn(t.in))
	}

}
