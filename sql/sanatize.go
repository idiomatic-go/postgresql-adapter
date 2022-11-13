package sql

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"strings"
)

// TODO : create a way to detect and remove SQL inject attacks
// DROP TABLE, DELETE FROM, SELECT * FROM, a double-dashed sequence ‘--’, or a semicolon ;
// quotes /*

func SanitizeString(s string) error {
	trimmed := util.TrimDoubleSpace(strings.ToLower(s))
	for _, t := range tokens {
		index := strings.Index(trimmed, t)
		if index != -1 {
			return errors.New(fmt.Sprintf("SQL injection embedded in string [%v] : %v", trimmed, t))
		}
	}
	return nil
}
