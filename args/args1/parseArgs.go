package args1

import (
	"fmt"
	"regexp"
)

type Args struct {
	name rune

}


const HelpMessage  = "help message"

func GetArgs(OSArgs []string) (args []string)   {
	args = OSArgs[1:]
	return
}

func PrintHelpMessage()  {
	fmt.Println(HelpMessage)
}

func MatchArg(arg string) bool {
	if match,_ := regexp.MatchString(`-([a-zA-Z]$)`,arg); match{
		return true
	}
	return false
}

func IterateArgs( args []string) ( ArgsMap  map[string]string, err error)  {
	index := 0
	for(index < len(args)){

		index ++
	}
	return ArgsMap, nil
}

func main()  {

	fmt.Println("https://golang.org/pkg/regexp/")
	fmt.Println("d")

}