package fizzbuzz4

import (
	"testing"
	"gotest.tools/assert"
	"fmt"
)

func TestFizzBuzz(t *testing.T) {
	testcases := []struct {
		number int
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{6, fizz},
		{5, buzz},
		{10, buzz},
		{15, fizzbuzz},
		{13, fizz},
		{52, buzz},
	}
	for _, testcase := range testcases {
		f := FizzBuzz{testcase.number}
		assert.Equal(t, f.Run(), testcase.result, fmt.Sprint(testcase))
	}
}
