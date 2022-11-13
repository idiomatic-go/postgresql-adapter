package dml

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

func ExampleWriteUpdateWhere() {
	sb := strings.Builder{}

	sc := WriteUpdateWhere(&sb, nil)
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", util.NilEmpty(sc.Error()))

	sc = WriteUpdateWhere(&sb, []util.Attr{{Name: "", Val: nil}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", util.NilEmpty(sc.Error()))

	sb.Reset()
	sc = WriteUpdateWhere(&sb, []util.Attr{{Name: "status_code", Val: "503"}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", util.NilEmpty(sc.Error()))

	sb.Reset()
	sc = WriteUpdateWhere(&sb, []util.Attr{{Name: "status_code", Val: "503"}, {Name: "minimum_code", Val: 99}, {Name: "created_ts", Val: sql.Function("now()")}})
	fmt.Printf("Stmt       : %v\n", util.NilEmpty(sb.String()))
	fmt.Printf("Error      : %v\n", util.NilEmpty(sc.Error()))

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
