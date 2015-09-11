package escaper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PerformOn(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		in, out string
	}{
		{in: "São França Ávido", out: "Sao Franca Avido"},
	}

	for _, t := range tests {
		a.Equal(t.out, PerformOn(t.in))
	}
}
