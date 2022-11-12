package dml

import (
	"fmt"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"reflect"
	"strings"
)

const (
	valueFmt  = "%v"
	stringFmt = "'%v'"
)

/*
INSERT INTO table_name (column_list) VALUES
    (value_list_1),
    (value_list_2),
    ...
    (value_list_n);

*/

func WriteInsert(sql string, values []any) string {
	sb := strings.Builder{}

	sb.WriteString(sql)
	sb.WriteString("\n")
	WriteInsertValues(&sb, values)
	sb.WriteString(";\n")
	return sb.String()
}

func WriteInsertValues(sb *strings.Builder, values []any) {
	// TODO : Check for Function in the attr value field
	max := len(values) - 1
	if max < 0 {
		return
	}
	sb.WriteString("(")
	for i, v := range values {
		t := reflect.TypeOf(v)
		if t.Kind() != reflect.String {
			sb.WriteString(fmt.Sprintf(valueFmt, v))
		} else {
			if _, function := v.(sql.Function); function {
				sb.WriteString(fmt.Sprintf(valueFmt, v))
			} else {
				safe := sql.Sanitize(v.(string))
				sb.WriteString(fmt.Sprintf(stringFmt, safe))
			}
		}
		if i < max {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")
}
