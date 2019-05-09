package fizzbuzz1

import "fmt"


func Magic(i int) string{
	if i%3 == 0 && i%5 == 0 {
		return "fizzbuzz"
	}
	if i%3 == 0{
		return "fizz"
	}
	if i%5 == 0{
		return "buzz"
	}
	return fmt.Sprintf("%d", i)
}

