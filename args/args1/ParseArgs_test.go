package args1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetArgs(t *testing.T) {
	OsArgs := []string{"par.go", "-d", "80", "-m", "helo"}
	expectArgs := []string{"-d", "80", "-m", "helo"}
	realArgs := GetArgs(OsArgs)
	if !reflect.DeepEqual(expectArgs, realArgs) {
		t.Error("get args failed")
	}

}

func TestMatchArg(t *testing.T) {
	ShouldMatchArg := []string{"-h", "-H"}
	ShouldNotMatchArg := []string{"-5", "A", "-h3"}
	var arg string
	for _, arg = range ShouldMatchArg {
		if !MatchArg(arg) {
			t.Errorf("%s should match\n", arg)
		}
	}
	for _, arg = range ShouldNotMatchArg {
		if MatchArg(arg) {
			t.Errorf("%s should not match\n", arg)
		}
	}
}

func TestMatchSchema(t *testing.T) {
	ShouldMatchSchemas := []string{"-5,-3", "-4", "false,node"}
	ShouldNotMatchSchemas := []string{"-d", "-h", "-node"}
	var arg string
	for _, arg = range ShouldMatchSchemas {
		if !matchSchema(arg) {
			t.Errorf("%s should match\n", arg)
		}
	}
	for _, arg = range ShouldNotMatchSchemas {
		if matchSchema(arg) {
			t.Errorf("%s should not match\n", arg)
		}
	}
}

func TestIterateArgs(t *testing.T) {
	RealArgs := []string{"-l", "-p", "8080", "-d", "/opt/logs"}
	expectMap := map[string]string{
		"l": "",
		"p": "8080",
		"d": "/opt/logs",
	}
	if ArgsMap, err := IterateArgs(RealArgs); err == "" {
		if !reflect.DeepEqual(ArgsMap, expectMap) {
			t.Errorf("ArgsMap: %s is not equal expectMap: %s\n", ArgsMap, expectMap)
		}
	}
}

func TestCommandP(t *testing.T) {
	errorSchemas := []string{"-12", "nd", "a2"}
	rightSchemas := []string{"80", "8080"}
	for _, schema := range errorSchemas {
		if err := CommandP(schema); err == "" {
			t.Errorf("schema: %s is not a valid schema for function p", schema)
		}
	}
	for _, schema := range rightSchemas {
		if err := CommandP(schema); err != "" {
			t.Errorf("schema: %s is a valid schema for function p", schema)
		}
	}
}

func TestCommandL(t *testing.T) {
	errorSchemas := []string{"-12", "nd", "a2"}
	rightSchemas := []string{"true", "false", ""}
	for _, schema := range errorSchemas {
		if err := CommandL(schema); err == "" {
			t.Errorf("schema: %s is not a valid schema for function p", schema)
		}
	}
	for _, schema := range rightSchemas {
		if err := CommandL(schema); err != "" {
			t.Errorf("schema: %s is a valid schema for function p", schema)
		}
	}
}

func TestProcessMap(t *testing.T) {
	testCases := []map[string]string{
		{
			"l": "",
			"p": "8080",
			"d": "/opt/logs",
		},
	}
	for _, testCase := range testCases {
		if err := ProcessMap(testCase); err != "" {
			fmt.Println(err)
			t.Errorf("map: %s ProcessMap error", testCase)
		}
	}
}
