package dml

import (
	"github.com/idiomatic-go/common-lib/util"
	"strings"
)

/*
UPDATE table_name
SET column1 = value1,
    column2 = value2,
    ...
WHERE condition;

*/

func WriteUpdate(sql string, attrs []util.Attr) (string, util.StatusCode) {
	var sb strings.Builder
	where := sql

	sb.WriteString(sql)
	sb.WriteString("\n")
	cond := WriteUpdateSet(sb, attrs)
	WriteUpdateWhere(sb, where, cond)
	return sb.String(), util.NewStatusOk()
}

func WriteUpdateSet(sb strings.Builder, attrs []util.Attr) []util.Attr {

	return nil
}

func WriteUpdateWhere(sb strings.Builder, where string, attrs []util.Attr) {

}
