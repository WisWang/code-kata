package fizzbuzz2

import (
	"testing"
	"gotest.tools/assert"
	"fmt"
)

func TestMagic(t *testing.T) {
	testcases := []struct {
		number int
		magic  string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{6, fizz},
		{5, buzz},
		{10, buzz},
		{15, fizzbuzz},
		{30, fizzbuzz},
		{13, fizz},
		{52, buzz},
	}
	for _, testcase := range testcases {
		f := FizzBuzz(testcase.number)
		assert.Equal(t, f.Run(), testcase.magic, fmt.Sprint(testcase))
	}

}
