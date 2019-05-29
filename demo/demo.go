package main

import (
	"fmt"
	"reflect"
)

type m struct {
	lines []string
}


func ms() (result string) {
	return
}

func main()  {
	fmt.Println("hello world")
	s := [][]string{{},{"12","13"}}
	fmt.Println(s)
	a := ms()
	fmt.Println("12,",a)
	fmt.Print(reflect.TypeOf(a))
	fmt.Print(a == "")
}
