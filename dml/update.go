package dml

import (
	"errors"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"strings"
)

const (
	where = "WHERE "
	and   = " AND "
	set   = "SET "
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
	err := WriteUpdateSet(&sb, attrs)
	if err != nil {
		return "", err
	}
	err = WriteUpdateWhere(&sb, where)
	return sb.String(), err
}

func WriteUpdateSet(sb *strings.Builder, attrs []util.Attr) error {
	max := len(attrs) - 1
	if max < 0 {
		return errors.New("invalid update argument, attrs slice is empty")
	}
	sb.WriteString(set)
	for i, attr := range attrs {
		s, err := sql.FmtAttr(attr)
		if err != nil {
			return err
		}
		sb.WriteString(s)
		if i < max {
			sb.WriteString(",\n")
		}
	}
	sb.WriteString("\n")
	return nil
}

func WriteUpdateWhere(sb *strings.Builder, attrs []util.Attr) error {
	max := len(attrs) - 1
	if max < 0 {
		return errors.New("invalid insert argument, attrs slice is empty")
	}
	sb.WriteString(where)
	for i, attr := range attrs {
		s, err := sql.FmtAttr(attr)
		if err != nil {
			return err
		}
		sb.WriteString(s)
		if i < max {
			sb.WriteString(and)
		}
	}
	sb.WriteString(";")
	return nil
}
