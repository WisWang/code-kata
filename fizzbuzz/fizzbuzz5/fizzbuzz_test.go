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
		{3, Fizz},
		{6, Fizz},
		{5, Buzz},
		{10, Buzz},
		{15, FizzBuzz},
		{51, FizzBuzz},
		{35, FizzBuzz},
		{13, Fizz},
		{31, Fizz},
		{59, Buzz},
	}
	for _, test := range tests {
		f := Game{test.input}
		assert.Equal(t, f.Magic(), test.output,fmt.Sprint(test))
	}
}
