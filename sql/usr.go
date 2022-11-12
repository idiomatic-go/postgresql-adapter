package sql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"reflect"
	"strings"
)

const (
	valueFmt  = "%v"
	stringFmt = "'%v'"
)

type Function string

// TODO : create a way to detect and remove SQL inject attacks
// DROP TABLE, DELETE FROM, SELECT * FROM, a double-dashed sequence ‘--’, or a semicolon ;
// quotes /*

var tokens = []string{"drop table", "delete from", "--", ";", "/*", "*/", "select * from"}

func SanitizeString(s string) util.StatusCode {
	sc := util.NewStatusOk()
	if s == "" {
		return sc
	}
	trimmed := trimAllSpace(strings.ToLower(s))
	for _, t := range tokens {
		index := strings.Index(trimmed, t)
		if index != -1 {
			sc = util.NewStatusError(errors.New(fmt.Sprintf("SQL injection embedded in string [%v] : %v", trimmed, t)))
		}
	}
	return sc
}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func FmtValue(v any) (string, util.StatusCode) {
	if util.IsNil(v) {
		return "NULL", util.NewStatusOk()
	}
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.String {
		return fmt.Sprintf(valueFmt, v), util.NewStatusOk()
	}
	if _, function := v.(Function); function {
		return fmt.Sprintf(valueFmt, v), util.NewStatusOk()
	}
	sc := SanitizeString(v.(string))
	if !sc.Ok() {
		logxt.LogPrintf("%v", sc)
		return "", sc
	}
	return fmt.Sprintf(stringFmt, v.(string)), util.NewStatusOk()
}
