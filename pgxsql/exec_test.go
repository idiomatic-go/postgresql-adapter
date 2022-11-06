package pgxsql

import (
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
)

//go:embed resource/*
var fsys embed.FS

func ExampleExec() {
	ctx := fse.ContextWithContent(nil, fsys, "resource/error.txt")

	cmd, err := Exec(ctx, ExecContentSql)
	fmt.Printf("Error  : %v\n", err)
	fmt.Printf("CmdTag : %v\n", cmd)

	ctx = fse.ContextWithContent(nil, fsys, "resource/command-tag.json")

	cmd, err = Exec(ctx, ExecContentSql)
	fmt.Printf("Error  : %v\n", err)
	fmt.Printf("CmdTag : %v\n", cmd)

	//Output:
	//Error  : example error text
	//CmdTag : { 0 <nil>}
	//Error  : <nil>
	//CmdTag : {select * 1000 <nil>}

}
