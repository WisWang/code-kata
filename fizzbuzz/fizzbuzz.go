package fizzbuzz

import "fmt"

func Magic(i int) (magicStr string) {
	magicStr = ""
	if i%3 == 0 {
		magicStr += "fizz"
	}
	if i%5  == 0 {
		magicStr += "buzz"
	}
	if magicStr == "" {
		magicStr = fmt.Sprintf("%d", i)
	}
	return
}

func MagicMain() (magics []string){
	number := 1
	for number < 20{
		magic := Magic(number)
		magics = append(magics, magic)
		number ++
	}
	return
}

