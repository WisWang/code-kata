package fizzbuzz8

import (
	"fmt"
	"strings"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzbuzz = fizz + buzz
)

type Game struct {
	input int
}

func (g Game) String() string {
	return fmt.Sprint(g.input)
}

func (g Game) Play() string {
	if g.special() != "" {
		return g.special()
	}
	return g.String()
}

func (g Game) special() (result string) {
	if g.means(3) {
		result += fizz
	}
	if g.means(5) {
		result += buzz
	}
	return result
}

func (g Game) means(input int) bool {
	return g.exactDivide(input) || strings.Contains(g.String(), fmt.Sprint(input))
}

func (g Game) exactDivide(dividend int) bool {
	return g.input%dividend == 0
}
