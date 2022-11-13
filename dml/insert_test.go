package dml

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

func ExampleWriteInsert() {
	stmt, err := WriteInsert(InsertSLOEntryStmt, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt    : %v\n", stmt)
	fmt.Printf("Error   : %v\n", err)

	//Output:
	//Stmt    : INSERT INTO slo_entry (id,customer_id,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,rps_low_comparison,rps_high_comparison,locality_scope,disable_processing,disable_triage,name,application,route_name,filter_status_codes,status_codes) VALUES
	//(100,'test string',false,nextval('slo_entry_Id'),now());
	//
	//
	//Error   : <nil>
}

func ExampleWriteInsertValues() {
	sb := strings.Builder{}

	err := WriteInsertValues(&sb, nil)
	fmt.Printf("Stmt    : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error   : %v\n", err)

	sb1 := strings.Builder{}
	err = WriteInsertValues(&sb1, []any{100})
	fmt.Printf("Stmt    : %v\n", sb1.String())
	fmt.Printf("Error   : %v\n", err)

	err = WriteInsertValues(&sb, []any{100, "test string", false, sql.Function(SLOEntryNextValFn), sql.Function(ChangedTimestampFn)})
	fmt.Printf("Stmt    : %v\n", sb.String())
	fmt.Printf("Error   : %v\n", err)

	//Output:
	//Stmt    : <nil>
	//Error   : invalid insert argument, values slice is empty
	//Stmt    : (100)
	//Error   : <nil>
	//Stmt    : (100,'test string',false,nextval('slo_entry_Id'),now())
	//Error   : <nil>
}
