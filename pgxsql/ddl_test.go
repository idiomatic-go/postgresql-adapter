package pgxsql

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
)

func _ExampleDDL() {
	buf, err := fse.ReadFile(fsys, "resource/ddl/slo_entry_table/sql")
	fmt.Printf("Err : %v\n", err)

	fmt.Printf("SQL : %v\n", string(buf))

	//Output:
	// fail
}

func _ExampleDDLSequence() {
	buf, err := fse.ReadFile(fsys, "resource/ddl/slo_entry_sequence.sql")
	fmt.Printf("Err : %v\n", err)

	fmt.Printf("SQL : %v\n", string(buf))

	//Output:
	// fail
}

func ExampleCreateRoles() {
	tag, err := ExecDDL("resource/ddl/create_roles.sql")

	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Tag   : %v\n", tag)

	//Output:
	//fail
}

func ExecDDL(name string) (CommandTag, error) {
	buf, err := fse.ReadFile(fsys, name)
	if err != nil {
		return CommandTag{}, err
	}
	s := string(buf)
	return Exec(context.Background(), s)
}
