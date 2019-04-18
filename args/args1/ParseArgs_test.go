package args

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

