package fizzbuzz6

import (
	"gotest.tools/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{6, fizz},
		{5, buzz},
		{10, buzz},
		{15, fizzbuzz},
		{13, fizz},
		{23, fizz},
		{52, buzz},
		{53, fizzbuzz},
		{51, fizzbuzz},
	}
	for _, test := range tests {
		g := Game{test.input}
		assert.Equal(t, g.Run(), test.output)
	}
}
