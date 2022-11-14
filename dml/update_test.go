package dml

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
	"time"
)

func ExampleWriteUpdate() {
	where := []util.Attr{{Name: "customer_id", Val: "customer1"}, {Name: "created_ts", Val: time.Now()}}
	attrs := []util.Attr{{Name: "status_code", Val: "503"}, {Name: "minimum_code", Val: 99}, {Name: "created_ts", Val: sql.Function("now()")}}

	sql, err := WriteUpdate(UpdateSLOEntryStmt, attrs, where)
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sql))
	fmt.Printf("Error      : %v\n", err)

	//Output:
	//Stmt       : UPDATE slo_entry
	//SET status_code = '503',
	//minimum_code = 99,
	//created_ts = now()
	//WHERE customer_id = 'customer1' AND created_ts = 2022/11/14 19:17:43.703121 ;
	//Error      : <nil>
}

func ExampleWriteUpdateSet() {
	sb := strings.Builder{}

	err := WriteUpdateSet(&sb, nil)
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	sb.Reset()
	err = WriteUpdateSet(&sb, []util.Attr{{Name: "status_code", Val: "503"}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	sb.Reset()
	err = WriteUpdateSet(&sb, []util.Attr{{Name: "status_code", Val: "503"}, {Name: "minimum_code", Val: 99}, {Name: "created_ts", Val: sql.Function("now()")}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	//Output:
	//Stmt       : <nil>
	//Error      : invalid updatet argument, attrs slice is empty
	//Stmt       : SET status_code = '503'
	//Error      : <nil>
	//	Stmt       : SET status_code = '503',
	//	minimum_code = 99,
	//	created_ts = now()
	//
	//Error      : <nil>
}

func ExampleWriteUpdateWhere() {
	sb := strings.Builder{}

	err := WriteUpdateWhere(&sb, nil)
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	err = WriteUpdateWhere(&sb, []util.Attr{{Name: "", Val: nil}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(strings.Trim(sb.String(), " ")))
	fmt.Printf("Error      : %v\n", err)

	sb.Reset()
	err = WriteUpdateWhere(&sb, []util.Attr{{Name: "status_code", Val: "503"}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	sb.Reset()
	err = WriteUpdateWhere(&sb, []util.Attr{{Name: "status_code", Val: "503"}, {Name: "minimum_code", Val: 99}, {Name: "created_ts", Val: sql.Function("now()")}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", err)

	//Output:
	//Stmt       : <nil>
	//Error      : invalid insert argument, attrs slice is empty
	//Stmt       : WHERE
	//Error      : invalid attribute argument, attribute name is empty
	//Stmt       : WHERE status_code = '503'
	//Error      : <nil>
	//Stmt       : WHERE status_code = '503' AND minimum_code = 99 AND created_ts = now()
	//Error      : <nil>

}
