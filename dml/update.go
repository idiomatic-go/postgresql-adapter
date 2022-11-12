package dml

import (
	"errors"
	"github.com/idiomatic-go/common-lib/logxt"
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

func WriteUpdate(sql string, attrs []util.Attr, where []util.Attr) (string, util.StatusCode) {
	var sb strings.Builder

	sb.WriteString(sql)
	sb.WriteString("\n")
	sc := writeUpdateSet(&sb, attrs)
	if !sc.Ok() {
		return "", sc
	}
	sc = WriteUpdateWhere(&sb, where)
	return sb.String(), sc
}

func writeUpdateSet(sb *strings.Builder, attrs []util.Attr) util.StatusCode {
	max := len(attrs) - 1
	if max < 0 {
		sc := util.NewStatusInvalidArgument(errors.New("invalid insert argument, attrs slice is empty"))
		logxt.LogPrintf("%v", sc)
		return sc
	}
	return util.NewStatusOk()
}

func writeUpdateWhere(sb *strings.Builder, attrs []util.Attr) util.StatusCode {
	max := len(attrs) - 1
	if max < 0 {
		sc := util.NewStatusInvalidArgument(errors.New("invalid insert argument, attrs slice is empty"))
		logxt.LogPrintf("%v", sc)
		return sc
	}
	sb.WriteString("WHERE ")
	for i, attr := range attrs {

	}
	return util.NewStatusOk()
}

func writeSet(sb *strings.Builder, attr util.Attr) util.StatusCode {
	if attr.Name == "" {
		sc := util.NewStatusInvalidArgument(errors.New("invalid set argument, attribute name is empty"))
		logxt.LogPrintf("%v", sc)
		return sc
	}
	if attr.Val == nil {
		sc := util.NewStatusInvalidArgument(errors.New("invalid set argument, attribute value is nil"))
		logxt.LogPrintf("%v", sc)
		return sc
	}

}

