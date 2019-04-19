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

func MatchSchema(arg string) bool {
	if match,_ := regexp.MatchString(`^-[0-9]`,arg); match{
		return true
	}
	if match,_ := regexp.MatchString(`^\W`,arg); match{
		return false
	}
	return true
}

func IterateArgs( args []string) ( ArgsMap  map[string]interface{}, err error)  {
	index := 0
	for(index < len(args)){
		if MatchArg(args[index]){
			ArgsMap[args[index]] = ""
		}

		index ++
	}
	return ArgsMap, nil
}

func main()  {

	fmt.Println("https://golang.org/pkg/regexp/")
	fmt.Println("d")

}