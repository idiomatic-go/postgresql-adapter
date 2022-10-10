package pgxsql_test

import (
	"fmt"
	"github.com/idiomatic-go/postgresql-adapter/pgxsql"
)

func ExampleQuery() {
	rows, err := pgxsql.Query(nil, "select * from $1", "tablename")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	if rows != nil {

	}
}
