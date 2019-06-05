package fizzbuzz

import (
	"testing"
	"fmt"
)

func TestMagicFizz(t *testing.T)  {
	testCases := []int{
		3,
		6,
		9,
	}
	for _,testcase:= range testCases{
		if "fizz" != Magic(testcase){
			t.Errorf("testcase: %d should be fizz\n",testcase)
		}
	}
}

func TestMagicBuzz(t *testing.T)  {
	testCases := []int{
		5,
		10,
		20,
	}
	for _,testcase:= range testCases{
		if "buzz" != Magic(testcase){
			t.Errorf("testcase: %d should be buzz\n",testcase)
		}
	}
}

func TestMagicFizzBuzz(t *testing.T)  {
	testCases := []int{
		15,
		30,
		45,
	}
	for _,testcase:= range testCases{
		if "fizzbuzz" != Magic(testcase){
			t.Errorf("testcase: %d should be buzz\n",testcase)
		}
	}
}

func TestMagicMain(t *testing.T) {
	//expextResults := []string{
	//	"1",
	//	"2",
	//	"fizz",
	//	"4",
	//	"buzz",
	//}
	fmt.Println(MagicMain())
}
