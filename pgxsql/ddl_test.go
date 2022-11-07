package pgxsql

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
)

func ExampleDDL() {
	buf, err := fse.ReadFile(fsys, "resource/ddl/slo_entry_table/sql")
	fmt.Printf("Err : %v\n", err)

	fmt.Printf("SQL : %v\n", string(buf))

	//Output:
	// fail
}

func ExampleDDLSequence() {
	buf, err := fse.ReadFile(fsys, "resource/ddl/slo_entry_sequence.sql")
	fmt.Printf("Err : %v\n", err)

	fmt.Printf("SQL : %v\n", string(buf))

	//Output:
	// fail
}
