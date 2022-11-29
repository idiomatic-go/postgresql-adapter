package pgxsql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
)

func ExampleQuery() {
	ctx := vhost.ContextWithAnyContent(nil, errors.New("example error text"))

	rows, sc := Query(ctx, "")
	fmt.Printf("Error  : %v\n", sc)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Error  : example error text
	//Rows   : <nil>

}

func ExampleQueryInvalidContent() {
	ctx := vhost.ContextWithAnyContent(nil, errors.New("json: cannot unmarshal object into Go value of type pgxsql.Rows"))

	rows, sc := Query(ctx, "")
	fmt.Printf("Error  : %v\n", sc)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Error  : json: cannot unmarshal object into Go value of type pgxsql.Rows
	//Rows   : <nil>

}
