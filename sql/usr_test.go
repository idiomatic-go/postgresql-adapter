package sql

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"time"
)

func ExampleSanitizeString() {

	sc := SanitizeString("")
	fmt.Printf("StatusCode  : %v\n", sc.Ok())

	sc = SanitizeString("adfsdfe4fc&*4")
	fmt.Printf("StatusCode  : %v\n", sc.Ok())

	sc = SanitizeString("test 1: /*")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 2: DROP   Table  ")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 3: DEL ETE FROM")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 4: - -")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 5: ;--")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 6: sa*/nitize 4 ;--")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = SanitizeString("test 7: of select  * froM customers")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	//Output:
	//StatusCode  : true
	//StatusCode  : true
	//StatusCode  : false SQL injection embedded in string [test 1: /*] : /*
	//StatusCode  : false SQL injection embedded in string [test 2: drop table] : drop table
	//StatusCode  : true
	//StatusCode  : true
	//StatusCode  : false SQL injection embedded in string [test 5: ;--] : ;
	//StatusCode  : false SQL injection embedded in string [test 6: sa*/nitize 4 ;--] : */
	//StatusCode  : false SQL injection embedded in string [test 7: of select * from customers] : select * from
}

func ExampleFmtValues() {
	var ptr *int

	v, sc := FmtValue(nil)
	fmt.Printf("Value  [nil]  : %v\n", util.NilEmpty(v))
	fmt.Printf("Status        : %v\n", sc.Ok())

	v, sc = FmtValue(ptr)
	fmt.Printf("Value  [ptr]  : %v\n", util.NilEmpty(v))
	fmt.Printf("Status        : %v\n", sc.Ok())

	v, sc = FmtValue(true)
	fmt.Printf("Value  [true] : %v\n", util.NilEmpty(v))
	fmt.Printf("Status        : %v\n", sc.Ok())

	v, sc = FmtValue(1001)
	fmt.Printf("Value  [1001] : %v\n", util.NilEmpty(v))
	fmt.Printf("Status        : %v\n", sc.Ok())

	v, sc = FmtValue("")
	fmt.Printf("Value  [\"\"]   : %v\n", util.NilEmpty(v))
	fmt.Printf("Status        : %v\n", sc.Ok())

	v, sc = FmtValue(time.Now())
	fmt.Printf("Value  [time.Now] : %v\n", util.NilEmpty(v))
	fmt.Printf("Status            : %v\n", sc.Ok())

	v, sc = FmtValue("test string")
	fmt.Printf("Value  [test string]  : %v\n", util.NilEmpty(v))
	fmt.Printf("Status                : %v\n", sc.Ok())

	v, sc = FmtValue(Function("now()"))
	fmt.Printf("Value  [now()]  : %v\n", util.NilEmpty(v))
	fmt.Printf("Status          : %v\n", sc.Ok())

	v, sc = FmtValue("drop table")
	fmt.Printf("Value  [drop table]  : %v\n", util.NilEmpty(v))
	fmt.Printf("Status               : %v\n", sc.Error())

	//Output:
	//Value  [nil]  : NULL
	//Status        : true
	//Value  [ptr]  : NULL
	//Status        : true
	//Value  [true] : true
	//Status        : true
	//Value  [1001] : 1001
	//Status        : true
	//Value  [""]   : ''
	//Status        : true
	//Value  [time.Now] : 2022-11-12 16:10:02.4825363 -0600 CST m=+0.004498101
	//Status            : true
	//Value  [test string]  : 'test string'
	//Status                : true
	//Value  [now()]  : now()
	//Status          : true
}
