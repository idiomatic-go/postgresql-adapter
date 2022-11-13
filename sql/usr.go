package sql

const (
	valueFmt  = "%v"
	stringFmt = "'%v'"
	attrFmt   = "%v = %v"
)

type Function string

var tokens = []string{"drop table", "delete from", "--", ";", "/*", "*/", "select * from"}
