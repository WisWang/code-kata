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

func (f FizzBuzz) contains3() bool {
	return strings.Contains(f.String(), "3")
}

func (f FizzBuzz) contains5() bool {
	return strings.Contains(f.String(), "5")
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
	if f.contains3() {
		return fizz
	}
	if f.contains5() {
		return buzz
	}
	return f.String()
}

func Magic(i int) string {
	f := FizzBuzz(i)
	return f.Run()
}
