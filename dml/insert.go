package dml

import (
	"github.com/idiomatic-go/common-lib/util"
	"strings"
)

/*
INSERT INTO table_name (column_list) VALUES
    (value_list_1),
    (value_list_2),
    ...
    (value_list_n);

*/

func WriteInsert(insert string, nextId string, attrs []util.Attr) string {
	var sb strings.Builder

	sb.WriteString(insert)
	sb.WriteString("\n")
	WriteInsertValues(sb, nextId, attrs)
	sb.WriteString(";")
	return sb.String()
}

func WriteInsertValues(sb strings.Builder, nextId string, attrs []util.Attr) {

}
