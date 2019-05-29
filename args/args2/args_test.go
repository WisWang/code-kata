package args2

import (
	"testing"
	"gotest.tools/assert"
)

func TestArgs(t *testing.T) {
	args := Args{[]string{"-l","-d","/opt","-p","8080"}}
	expectResult := [][]string{
		{"l",""},
		{"d","/opt"},
		{"p","8080"},

	}
	assert.DeepEqual(t,args.Result(),expectResult)
}

