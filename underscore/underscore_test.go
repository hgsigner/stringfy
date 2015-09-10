package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Undercore(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		in, out string
	}{
		{in: "    Fizz    ", out: "fizz"},
		{in: "Fizz", out: "fizz"},
		{in: "Fizz Buzz", out: "fizz_buzz"},
		{in: "fizz_buzz", out: "fizz_buzz"},
		{in: "Fázz Bôzz", out: "fazz_bozz"},
	}

	for _, t := range tests {
		a.Equal(t.out, PerformOn(t.in))
	}

}
