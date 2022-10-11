package pgxsql

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
)

type lookup struct {
	Name  string
	Email string
}

var lookupEqual util.IsEqual = func(key, val any) bool {
	if key == nil || val == nil {
		return false
	}
	k, ok := key.(string)
	if !ok {
		return false
	}
	v, ok1 := val.(lookup)
	if !ok1 {
		return false
	}
	return k == v.Name
}

func _ExampleAnyTableLookup() {
	data := []lookup{lookup{"dave", "test@google.com"}, lookup{Name: "bill", Email: "billg@msn.com"}, lookup{Name: "zuck", Email: "king@facebook.com"}}

	tag, err := devExec(nil, "POST email", lookupEqual)
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "PUT email", data[0])
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "GET email/dave")
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "DELETE email/dave")
	fmt.Printf("Result : %v %v\n\n", tag, err)

	// Next table
	tag, err = devExec(nil, "POST email2", lookupEqual)
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "PUT email2", data[1])
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "GET email2/bill")
	fmt.Printf("Result : %v %v\n", tag, err)

	tag, err = devExec(nil, "DELETE email2/bill")
	fmt.Printf("Result : %v %v\n", tag, err)

	//Output:
	// Count : 0
}
