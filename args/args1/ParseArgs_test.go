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

func TestIterateArgs(t *testing.T)  {
	RealArgs := []string{"-", "-p", "8080", "-d", "/opt/logs"}
	if ArgsMap ,err:= IterateArgs(RealArgs);err == nil{

	}
}

