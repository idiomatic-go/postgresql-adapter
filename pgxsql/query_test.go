package pgxsql

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
)

func ExampleQuery() {
	ctx := fse.ContextWithContent(nil, fsys, "resource/error.txt")

	rows, sc := Query(ctx, ExecContentSql)
	fmt.Printf("Error  : %v\n", sc)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Error  : example error text
	//Rows   : <nil>

}

func ExampleQueryInvalidContent() {
	ctx := fse.ContextWithContent(nil, fsys, "resource/rows.json")

	rows, sc := Query(ctx, ExecContentSql)
	fmt.Printf("Error  : %v\n", sc)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Error  : json: cannot unmarshal object into Go value of type pgxsql.Rows
	//Rows   : <nil>

}
