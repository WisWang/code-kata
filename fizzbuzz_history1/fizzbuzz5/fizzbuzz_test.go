package fizzbuzz5

import (
	"fmt"
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
		{51, fizzbuzz},
		{35, fizzbuzz},
		{13, fizz},
		{31, fizz},
		{59, buzz},
	}
	for _, test := range tests {
		f := Game{test.input}
		assert.Equal(t, f.Magic(), test.output,fmt.Sprint(test))
	}
}
