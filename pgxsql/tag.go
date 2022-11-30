package pgxsql

import (
	"github.com/idiomatic-go/postgresql-adapter/sql"
	"github.com/jackc/pgx/v5/pgconn"
)

func StatementToString(tag pgconn.CommandTag) string {
	if tag.Insert() {
		return sql.Insert
	}
	if tag.Update() {
		return sql.Update
	}
	if tag.Delete() {
		return sql.Delete
	}
	return sql.Select
}
