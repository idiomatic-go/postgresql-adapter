package pgxsql

import (
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
)

//go:embed resource/*
var fs embed.FS

func init() {
	execContentOverride = true
}

func NilEmpty(s string) string {
	if s == "" {
		return "<nil>"
	}
	return s
}

func ExampleExec() {
	ctx := fse.ContextWithContent(nil, fs, "resource/error.txt")

	cmd, sc := Exec(ctx, "")
	fmt.Printf("Error  : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	ctx = fse.ContextWithContent(nil, fs, "resource/command-tag.json")

	cmd, sc = Exec(ctx, "")
	fmt.Printf("Error  : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	//Output:
	//Error  : example error text
	//CmdTag : { 0 false false false false}
	//Error  : <nil>
	//CmdTag : {select * 1000 false false false false}

}
