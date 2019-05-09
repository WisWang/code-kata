package fizzbuzz1

import (
	"testing"
	"gotest.tools/assert"
)


func TestMagicFizz(t *testing.T){
	testcases := []int{
		3,
		6,
		9,
	}
	for _, testcase := range testcases{
		assert.Equal(t, Magic(testcase),"fizz","the testcase should be fizz")
	}
}

func TestMagicBuzz(t *testing.T) {
	testcases := []int{
		5,
		10,
	}
	for _, testcase := range testcases{
		assert.Equal(t, Magic(testcase),"buzz","the testcase should be fuzz")
	}
}

func TestMagicFizzBuzz(t *testing.T){
	testcases := []int{
		15,
		30,
	}
	for _, testcase := range testcases{
		assert.Equal(t, Magic(testcase),"fizzbuzz",)
	}
}