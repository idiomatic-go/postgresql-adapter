package dml

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

func ExampleWriteInsert() {
	stmt, sc := WriteInsert(InsertSLOEntryStmt, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt       : %v\n", stmt)
	fmt.Printf("StatusCode : %v\n", sc.Ok())

	//Output:
	//fail
}

func ExampleWriteInsertValues() {
	sb := strings.Builder{}

	sc := WriteInsertValues(&sb, nil)
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("StatusCode : %v\n", sc.Ok())

	sb1 := strings.Builder{}
	sc = WriteInsertValues(&sb1, []any{100})
	fmt.Printf("Stmt       : %v\n", sb1.String())
	fmt.Printf("StatusCode : %v\n", sc.Ok())

	sc = WriteInsertValues(&sb, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt       : %v\n", sb.String())
	fmt.Printf("StatusCode : %v\n", sc.Ok())

	//Output:
	//Stmt       : <nil>
	//StatusCode : false
	//Stmt       : (100)
	//StatusCode : true
	//Stmt       : (100,'test string',false,nextval('slo_entry_Id'),now())
	//StatusCode : true
}
