package sql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"strings"
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
			sc = util.NewStatusError(errors.New(fmt.Sprintf("invalid content embedded in string %v : %v", trimmed, t)))
		}
	}
	return sc
}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
