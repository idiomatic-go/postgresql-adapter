package pgxsql

import (
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/util"
)

//go:embed resource/*
var fs embed.FS

func init() {
	execContentOverride = true
}

func ExampleExec() {
	ctx := fse.ContextWithContent(nil, fs, "resource/error.txt")

	cmd, sc := Exec(ctx, "")
	fmt.Printf("Error  : %v\n", util.NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	ctx = fse.ContextWithContent(nil, fs, "resource/command-tag.json")

	cmd, sc = Exec(ctx, "")
	fmt.Printf("Error  : %v\n", util.NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	//Output:
	//Error  : example error text
	//CmdTag : { 0 <nil>}
	//Error  : <nil>
	//CmdTag : {select * 1000 <nil>}

}
