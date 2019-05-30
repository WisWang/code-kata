package fizzbuzz5

import (
	"fmt"
	"strings"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzbuzz = "fizzbuzz"
)

type Game struct {
	number int
}


func (f Game) Magic() string {
	if f.isSpecial() {
		return f.special()
	}
	return f.String()
}

func (f Game) contains(sub int) bool {
	return strings.Contains(f.String(), fmt.Sprint(sub))
}

func (f Game) String() string {
	return fmt.Sprint(f.number)
}

func (f Game) isSpecial() bool {
	return f.exactDivid(3) ||
		f.exactDivid(5) ||
		f.contains(3) ||
		f.contains(5)
}

func (f Game) special() (result string) {
	if f.exactDivid(3) || f.contains(3) {
		result += fizz
	}
	if f.exactDivid(5) || f.contains(5) {
		result += buzz
	}
	return result
}

func (f Game) exactDivid(dividend int) bool {
	return f.number%dividend == 0
}
