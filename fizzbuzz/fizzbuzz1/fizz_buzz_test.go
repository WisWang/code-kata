package fizzbuzz1

import (
	"testing"
	"gotest.tools/assert"
)


type testCase struct {
	number int
	magic string
}

func TestMagic(t *testing.T){
	testcases := []testCase{
		{1,"1"},
		{2,"2"},
		{3,"fizz"},
		{6,"fizz"},
		{5,"buzz"},
		{10,"buzz"},
		{15,"fizzbuzz"},
		{30,"fizzbuzz"},
	}
	for _, testcase := range testcases{
		assert.Equal(t, Magic(testcase.number), testcase.magic)
	}
}



