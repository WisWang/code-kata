package args1

import (
	"fmt"
	"regexp"
	"strings"
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
	if match,_ := regexp.MatchString(`^/`,arg); match{
		return true
	}
	if match,_ := regexp.MatchString(`^\W`,arg); match{
		return false
	}
	return true
}

func IterateArgs( args []string) ( ArgsMap  map[string]interface{}, err string)  {
	index := 0
	ArgsMap = make(map[string]interface{})
	for(index < len(args)){
		if MatchArg(args[index]) {
			//if index+1 >= len(args){
			//	ArgsMap[strings.Replace(args[index],"-","",-1)] = ""
			//	index ++
			//	continue
			//}
			if MatchSchema(args[index+1])  {
				ArgsMap[strings.Replace(args[index],"-","",-1)] = args[index+1]
				index ++
			}else {
				ArgsMap[strings.Replace(args[index],"-","",-1)] = ""
			}
		}
		index ++
	}
	return ArgsMap, ""
}

func main()  {

	fmt.Println("https://golang.org/pkg/regexp/")
	fmt.Println("d")

}