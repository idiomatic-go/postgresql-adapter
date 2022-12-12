package pgxsql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fncall"
	"reflect"
)

type rowsT struct {
}

func (r *rowsT) Close()     {}
func (r *rowsT) Err() error { return nil }
func (r *rowsT) CommandTag() CommandTag {
	return CommandTag{Sql: "select *", RowsAffected: 1, Insert: false, Update: false, Delete: false, Select: true}
}
func (r *rowsT) FieldDescriptions() []FieldDescription { return nil }
func (r *rowsT) Next() bool                            { return false }
func (r *rowsT) Scan(dest ...any) error                { return nil }
func (r *rowsT) Values() ([]any, error)                { return nil, nil }
func (r *rowsT) RawValues() [][]byte                   { return nil }

func ExampleQueryStatus() {
	ctx := fncall.ContextWithContent(nil, errors.New("example error text"))

	rows, status := Query(ctx, "")
	fmt.Printf("Status : %v\n", status)
	fmt.Printf("Rows   : %v\n", rows)

	//Output:
	//Status : example error text
	//Rows   : <nil>

}

func ExampleQueryRows() {
	var i Rows = &rowsT{}

	ctx := fncall.ContextWithContent(nil, i)

	rows, status := Query(ctx, "")
	fmt.Printf("Ok         : %v\n", status.Ok())
	fmt.Printf("Rows       : %v\n", reflect.TypeOf(rows))
	fmt.Printf("CommandTag : %v\n", rows.CommandTag())

	//Output:
	//Ok         : true
	//Rows       : *pgxsql.rowsT
	//CommandTag : {select * 1 false false false true}

}
