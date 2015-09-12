package camelcase_test

import (
	"testing"

	"github.com/hgsigner/stringfy/camelcase"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		in, out string
	}{
		{in: "fizz buzz bazz", out: "FizzBuzzBazz"},
		{in: "fizz_buzz_bazz", out: "FizzBuzzBazz"},
		{in: "fizz-buzz-bazz", out: "FizzBuzzBazz"},
		{in: "Fizz Buzz Bazz", out: "FizzBuzzBazz"},
		{in: "São Ñino França Ávido Caça", out: "SaoNinoFrancaAvidoCaca"},
	}

	for _, t := range tests {
		cc := camelcase.PerformOn(t.in)
		a.Equal(t.out, cc)
	}
}
