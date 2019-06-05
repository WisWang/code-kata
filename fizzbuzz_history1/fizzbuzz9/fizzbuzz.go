package fizzbuzz9

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

func (g Game) Play() (result string) {
	result = g.specail()
	if result != "" {
		return
	}
	return g.String()
}

func (g Game) specail() (result string) {
	if g.means(3) {
		result += fizz
	}
	if g.means(5){
		result += buzz
	}
	return
}

func (g Game) means(specialNumber int) bool {
	return g.exactDivide(specialNumber) || g.contains(specialNumber)
}

func (g Game) contains(specialNumber int) bool {
	return strings.Contains(g.String(), fmt.Sprint(specialNumber))
}

func (g Game) exactDivide(dividend int) bool {
	return g.input%dividend == 0
}
