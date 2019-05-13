package main

import (
	"fmt"

)

type FizzBuzz int

func (f FizzBuzz) String() string {
	return fmt.Sprintf("hello %d", f)
}

func main()  {
	f := FizzBuzz(5)
	f1 := FizzBuzz(6)
	fmt.Println(f)
	fmt.Println(f1)
	fmt.Println(f.String())

}
