package fizzbuzz7

import (
	"fmt"
	"strings"
)

const (
	fizz = "fizz"
	buzz = "buzz"
	fizzbuzz = fizz + buzz
)

type Game struct {
	input int
}

func (g Game) String() string {
	return fmt.Sprint(g.input)
}

func (g Game) Play() string {
	if g.isSpecial(){
		return g.special()
	}
	return g.String()
}

func (g Game) exactDivide(dividend int) bool {
	return g.input%dividend == 0
}

func (g Game) isSpecial() bool {
	return g.exactDivide(3) ||
		g.exactDivide(5) ||
		g.contains(3) ||
		g.contains(5)
}

func (g Game) special() (result string) {
	if g.exactDivide(3) || g.contains(3) {
		result += fizz
	}
	if g.exactDivide(5) || g.contains(5) {
		result += buzz
	}
	return
}

func (g Game) contains(input int) bool {
	return strings.Contains(g.String(), fmt.Sprint(input))
}
