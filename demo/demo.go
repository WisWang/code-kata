package main

import (
	"fmt"
	"strings"
)

func GenMap() (M map[string]interface{}) {
	M= make(map[string]interface{})
	M["nar"] = "sheme"
	return
}

func main()  {
	fmt.Println("w")
	l := []string{
		"12", "123",
	}
	s := "id"
	fmt.Println(strings.Replace(s,"i","",-1))
	fmt.Println(GenMap())
	fmt.Println(len(l))
	fmt.Println(l[1])
}