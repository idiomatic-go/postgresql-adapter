package dml

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

func ExampleWriteInsert() {
	stmt := WriteInsert(InsertSLOEntryStmt, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt  : %v\n", stmt)

	//Output:
	//fail
}

func ExampleWriteInsertValues() {
	sb := strings.Builder{}

	WriteInsertValues(&sb, nil)
	fmt.Printf("Stmt  : %v\n", util.NilEmpty(sb.String()))

	sb1 := strings.Builder{}
	WriteInsertValues(&sb1, []any{100})
	fmt.Printf("Stmt  : %v\n", sb1.String())

	WriteInsertValues(&sb, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt  : %v\n", sb.String())

	//Output:
	//Stmt  : <nil>
	//Stmt  : (100)
	//Stmt  : (100,'test string',false,nextval('slo_entry_Id'),now())
}
