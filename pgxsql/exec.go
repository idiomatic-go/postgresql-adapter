package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/idiomatic-go/postgresql-adapter/dml"
	"github.com/idiomatic-go/postgresql-adapter/sql"
)

var execContentOverride = false

func ExecInsert(ctx context.Context, sql string, values []any) (CommandTag, vhost.Status) {
	stmt, err := dml.WriteInsert(sql, values)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return Exec(ctx, stmt)
}

func ExecUpdate(ctx context.Context, sql string, attrs []sql.Attr, where []sql.Attr) (CommandTag, vhost.Status) {
	stmt, err := dml.WriteUpdate(sql, attrs, where)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return Exec(ctx, stmt)
}

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, vhost.Status) {
	if execContentOverride {
		tag, err := fse.ProcessContent[CommandTag](ctx)
		return tag, vhost.NewStatusInvalidArgument(err)
	}
	if dbClient == nil {
		return CommandTag{}, vhost.NewStatusInvalidArgument(errors.New("error on PostgreSQL exec call : dbClient is nil"))
	}
	// Transaction processing.
	txn, err0 := dbClient.Begin(ctx)
	if err0 != nil {
		return CommandTag{}, vhost.NewStatusError(err0)
	}
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		err0 := txn.Rollback(ctx)
		return CommandTag{}, vhost.NewStatusError(err, err0)
	}
	err = txn.Commit(ctx)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected(), Insert: t.Insert(), Update: t.Update(), Delete: t.Delete(), Select: t.Select()}, vhost.NewStatusOk()
}
