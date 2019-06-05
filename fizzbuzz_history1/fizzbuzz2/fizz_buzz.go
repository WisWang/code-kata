package fizzbuzz2

import (
	"fmt"
	"strings"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzbuzz = "fizzbuzz"
)

type FizzBuzz int

func (f FizzBuzz) String() string {
	return fmt.Sprintf("%d", f)
}

func (f FizzBuzz) isFizz() bool {
	return f%3 == 0
}

func (f FizzBuzz) isBuzz() bool {
	return f%5 == 0
}

func (f FizzBuzz) isFizzBuzz() bool {
	return f.isFizz() && f.isBuzz()
}

func (f FizzBuzz) contains(sub string) bool {
	return strings.Contains(f.String(), sub)
}


func (f FizzBuzz) Run() string {
	if f.isFizzBuzz() {
		return fizzbuzz
	}
	if f.isFizz() {
		return fizz
	}
	if f.isBuzz() {
		return buzz
	}
	if f.contains("3") {
		return fizz
	}
	if f.contains("5") {
		return buzz
	}
	return f.String()
}
