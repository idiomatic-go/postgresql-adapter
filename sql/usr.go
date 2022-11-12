package sql

type Function string

// TODO : create a way to detect and remove SQL inject attacks
// DROP TABLE, DELETE FROM, SELECT * FROM, a double-dashed sequence ‘--’, or a semicolon ;
// quotes /*
func Sanitize(s string) string {
	return s
}
