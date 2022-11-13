package dml

import (
	"errors"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

/*
INSERT INTO table_name (column_list) VALUES
    (value_list_1),
    (value_list_2),
    ...
    (value_list_n);

*/

func WriteInsert(sql string, values []any) (string, error) {
	sb := strings.Builder{}

	sb.WriteString(sql)
	sb.WriteString("\n")
	err := WriteInsertValues(&sb, values)
	sb.WriteString(";\n")
	return sb.String(), err
}

func WriteInsertValues(sb *strings.Builder, values []any) error {
	max := len(values) - 1
	if max < 0 {
		return errors.New("invalid insert argument, values slice is empty")
	}
	sb.WriteString("(")
	for i, v := range values {
		s, err := sql.FmtValue(v)
		if err != nil {
			return err
		}
		sb.WriteString(s)
		if i < max {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")
	return nil
}
