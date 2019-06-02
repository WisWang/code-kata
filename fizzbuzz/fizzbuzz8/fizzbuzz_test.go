package fizzbuzz8

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input  int
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{6, fizz},
		{13, fizz},
		{23, fizz},
		{5, buzz},
		{52, buzz},
		{59, buzz},
		{51, fizzbuzz},
		{10, buzz},
		{15, fizzbuzz},
		{30, fizzbuzz},
	}
	for _, test := range tests {
		g := Game{test.input}
		assert.Equal(t, g.Play(), test.result,fmt.Sprint(test))
	}
}
