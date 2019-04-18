package args

import (
	"fmt"
)

struct Args {

}


const HelpMessage  = "help message"

func GetArgs(OSArgs []string) (args []string)   {
	args = OSArgs[1:]
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

	fmt.Println("https://golang.org/pkg/regexp/")
	fmt.Println("d")

}