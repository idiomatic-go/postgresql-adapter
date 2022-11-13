package dml

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

/*
UPDATE table_name
SET column1 = value1,
    column2 = value2,
    ...
WHERE condition;

*/

func WriteUpdate(sql string, attrs []util.Attr, where []util.Attr) (string, error) {
	var sb strings.Builder

	sb.WriteString(sql)
	sb.WriteString("\n")
	err := WriteUpdateSet(&sb, attrs[0])
	if err != nil {
		return "", err
	}
	err = WriteUpdateWhere(&sb, where)
	return sb.String(), err
}

func WriteUpdateWhere(sb *strings.Builder, attrs []util.Attr) error {
	max := len(attrs) - 1
	if max < 0 {
		return errors.New("invalid insert argument, attrs slice is empty")
	}
	sb.WriteString("WHERE ")
	for i, attr := range attrs {
		s, err := sql.FmtAttr(attr)
		if err != nil {
			return err
		}
		sb.WriteString(s)
		if i < max {
			sb.WriteString(" AND ")
		}
	}
	return nil
}

func WriteUpdateSet(sb *strings.Builder, attr util.Attr) error {
	if attr.Name == "" {
		return errors.New("invalid set argument, attribute name is empty")
	}
	s, err := sql.FmtValue(attr.Val)
	if err != nil {
		return err
	}
	sb.WriteString(fmt.Sprintf("%v=%v", attr.Name, s))
	return nil
}
