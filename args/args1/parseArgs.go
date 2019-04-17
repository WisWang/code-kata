package main

import (
	"fmt"
	"os"
	"reflect"
)

const HelpMessage  = "help message"

func GetArgs() (args []string)   {
	args = os.Args[1:]
	return
}

func PrintHelpMessage()  {
	fmt.Println(HelpMessage)
}

func IterateArgs( args []string) ( m2  map[string]interface{})  {
	Index := 0
	for(Index < len(args)){

		Index ++
	}
	return
}

func main()  {

	fmt.Println(reflect.TypeOf(GetArgs()[0]))
	fmt.Println(reflect.TypeOf(GetArgs()[1]))

}