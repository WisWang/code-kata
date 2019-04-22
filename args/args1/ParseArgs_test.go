package args1

import (
	"testing"
	"reflect"

)

func TestGetArgs(t *testing.T) {
	OsArgs := []string{"par.go","-d", "80", "-m","helo"}
	expectArgs := []string{"-d", "80", "-m","helo"}
	realArgs := GetArgs(OsArgs)
	if !reflect.DeepEqual(expectArgs, realArgs) {
		t.Error("get args failed")
	}

}

func TestMatchArg(t *testing.T){
	ShouldMatchArg := []string{"-h", "-H"}
	ShouldNotMatchArg := []string{"-5","A","-h3"}
	var arg string
	for _,arg = range ShouldMatchArg{
		if !MatchArg(arg){
			t.Errorf("%s should match\n", arg)
		}
	}
	for _,arg = range ShouldNotMatchArg{
		if MatchArg(arg){
			t.Errorf("%s should not match\n", arg)
		}
	}
}

func TestMatchSchema(t *testing.T) {
	ShouldMatchSchemas := []string{"-5,-3", "-4","false,node"}
	ShouldNotMatchSchemas := []string{"-d","-h", "-node"}
	var arg string
	for _,arg = range ShouldMatchSchemas{
		if !MatchSchema(arg){
			t.Errorf("%s should match\n", arg)
		}
	}
	for _,arg = range ShouldNotMatchSchemas{
		if MatchSchema(arg){
			t.Errorf("%s should not match\n", arg)
		}
	}
}

func TestIterateArgs(t *testing.T)  {
	RealArgs := []string{"-l", "-p", "8080", "-d", "/opt/logs"}
	expectMap := map[string]interface{}{
		"l" : "",
		"p" : "8080",
		"d" : "/opt/logs",
	}
	if ArgsMap ,err:= IterateArgs(RealArgs);err == ""{
		if !reflect.DeepEqual(ArgsMap, expectMap){
			t.Errorf("ArgsMap: %s is not equal expectMap: %s", ArgsMap, expectMap)
		}
	}
}



