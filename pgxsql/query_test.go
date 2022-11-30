package pgxsql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
)

func ExampleQueryStatus() {
	ctx := vhost.ContextWithContent(nil, errors.New("example error text"))

	rows, status := Query(ctx, "")
	fmt.Printf("Status : %v\n", status)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Status : example error text
	//Rows   : <nil>

}

func ExampleQueryRows() {
	ctx := vhost.ContextWithContent(nil, errors.New("json: cannot unmarshal object into Go value of type pgxsql.Rows"))

	rows, status := Query(ctx, "")
	fmt.Printf("Status : %v\n", status)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Status : json: cannot unmarshal object into Go value of type pgxsql.Rows
	//Rows   : <nil>

}
