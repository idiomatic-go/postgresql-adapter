package dml

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
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

func WriteInsert(sql string, values []any) (string, util.StatusCode) {
	sb := strings.Builder{}

	sb.WriteString(sql)
	sb.WriteString("\n")
	sc := WriteInsertValues(&sb, values)
	sb.WriteString(";\n")
	return sb.String(), sc
}

func WriteInsertValues(sb *strings.Builder, values []any) util.StatusCode {
	max := len(values) - 1
	if max < 0 {
		sc := util.NewStatusInvalidArgument(errors.New("invalid insert argument, values slice is empty"))
		logxt.LogPrintf("%v", sc)
		return sc
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
				sc := sql.SanitizeString(v.(string))
				if !sc.Ok() {
					logxt.LogPrintf("%v", sc)
					return sc
				}
				sb.WriteString(fmt.Sprintf(stringFmt, v.(string)))
			}
		}
		if i < max {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")
	return util.NewStatusOk()
}
