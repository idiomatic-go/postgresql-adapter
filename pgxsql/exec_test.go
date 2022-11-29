package pgxsql

import (
	"embed"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
)

//go:embed resource/*
var fs embed.FS

func NilEmpty(s string) string {
	if s == "" {
		return "<nil>"
	}
	return s
}

func ExampleExec() {
	ctx := vhost.ContextWithAnyContent(nil, errors.New("example error text"))

	cmd, sc := Exec(ctx, "")
	fmt.Printf("Error  : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	ctx = vhost.ContextWithAnyContent(nil, CommandTag{Sql: "select *", RowsAffected: 1000, Insert: false, Update: false, Delete: false, Select: true})

	cmd, sc = Exec(ctx, "")
	fmt.Printf("Error  : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("CmdTag : %v\n", cmd)

	//Output:
	//Error  : example error text
	//CmdTag : { 0 false false false false}
	//Error  : <nil>
	//CmdTag : {select * 1000 false false false true}

}
