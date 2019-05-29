package fizzbuzz4

import (
	"fmt"
	"strings"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzbuzz = "fizzbuzz"
)

type FizzBuzz struct {
	number int
}

func (f FizzBuzz) exactDivide(dividend int) bool {
	return f.number%dividend == 0
}

func (f FizzBuzz) Strings() string {
	return fmt.Sprint(f.number)
}

func (f FizzBuzz) contains(sub string) bool {
	return strings.Contains(f.Strings(), sub)
}

func (f FizzBuzz) Run() string {
	if f.exactDivide(3) && f.exactDivide(5) {
		return fizzbuzz
	}
	if f.exactDivide(3) {
		return fizz
	}
	if f.exactDivide(5) {
		return buzz
	}
	if f.contains("3") {
		return fizz
	}
	if f.contains("5") {
		return buzz
	}
	return f.Strings()
}

