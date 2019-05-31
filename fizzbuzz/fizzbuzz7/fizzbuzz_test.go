package fizzbuzz7

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
		{10, buzz},
		{52, buzz},
		{15, fizzbuzz},
		{30, fizzbuzz},
		{53, fizzbuzz},
		{51, fizzbuzz},
	}
	for _, test := range tests {
		g := Game{test.input}
		assert.Equal(t, g.Play(), test.result,fmt.Sprint(test))
	}
}
