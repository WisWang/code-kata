package args1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Args struct {
	name rune
}

const HelpMessage = "help message"

func GetArgs(OSArgs []string) (args []string) {
	args = OSArgs[1:]
	return
}

func MatchArg(arg string) bool {
	if match, _ := regexp.MatchString(`-([a-zA-Z]$)`, arg); match {
		return true
	}
	return false
}

func MatchSchema(arg string) bool {
	if match, _ := regexp.MatchString(`^-[0-9]`, arg); match {
		return true
	}
	if match, _ := regexp.MatchString(`^/`, arg); match {
		return true
	}
	if match, _ := regexp.MatchString(`^\W`, arg); match {
		return false
	}
	return true
}

func IterateArgs(args []string) (ArgsMap map[string]string, err string) {
	index := 0
	ArgsMap = make(map[string]string)
	for index < len(args) {
		if MatchArg(args[index]) {
			if index+1 == len(args) {
				ArgsMap[strings.Replace(args[index], "-", "", -1)] = ""
				break
			}
			if MatchSchema(args[index+1]) {
				ArgsMap[strings.Replace(args[index], "-", "", -1)] = args[index+1]
				index++
			} else {
				ArgsMap[strings.Replace(args[index], "-", "", -1)] = ""
			}
		}
		index++
	}
	return ArgsMap, ""
}

func ProcessMap(ArgsMap map[string]string) (err string) {
	err = ""
	for arg, schema := range ArgsMap {
		switch arg {
		case "d":
			err = CommandD(schema)
		case "p":
			err = CommandP(schema)
		case "l":
			err = CommandL(schema)
		default:
			err = fmt.Sprintf("not support flags: %s", arg)
		}
		if err != "" {
			return
		}
	}
	return
}

func CommandD(schema string) (commandErr string) {
	fmt.Printf("d function process %s\n", schema)
	return ""
}
func CommandP(schema string) (commandErr string) {
	portnumber, err := strconv.Atoi(schema)
	if err != nil || portnumber < 0 {
		commandErr = "p function only support positive int"
		return
	}
	fmt.Printf("CommandP function process %s\n", schema)
	return ""
}

func CommandL(schema string) (commandErr string) {
	if schema == "" {
		schema = "true"
	}
	logging, err := strconv.ParseBool(schema)
	if err != nil {
		commandErr = "CommandL function only support bool"
		return
	}
	fmt.Printf("l function process %t\n", logging)
	return ""
}
