package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/dml"
)

var execContentOverride = false

func ExecInsert(ctx context.Context, sql string, values []any) (CommandTag, util.Status) {
	stmt, err := dml.WriteInsert(sql, values)
	if err != nil {
		return CommandTag{}, util.NewStatusError(err)
	}
	return Exec(ctx, stmt)
}

func ExecUpdate(ctx context.Context, sql string, attrs []util.Attr, where []util.Attr) (CommandTag, util.Status) {
	stmt, err := dml.WriteUpdate(sql, attrs, where)
	if err != nil {
		return CommandTag{}, util.NewStatusError(err)
	}
	return Exec(ctx, stmt)
}

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, util.Status) {
	if execContentOverride {
		tag, err := fse.ProcessContent[CommandTag](ctx)
		return tag, util.NewStatusInvalidArgument(err)
	}
	if dbClient == nil {
		return CommandTag{}, util.NewStatusInvalidArgument(errors.New("error on PostgreSQL exec call : dbClient is nil"))
	}
	// Transaction processing.
	txn, err0 := dbClient.Begin(ctx)
	if err0 != nil {
		return CommandTag{}, util.NewStatusError(err0)
	}
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		err0 := txn.Rollback(ctx)
		return CommandTag{}, util.NewStatusError(err, err0)
	}
	err = txn.Commit(ctx)
	if err != nil {
		return CommandTag{}, util.NewStatusError(err)
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected(), Insert: t.Insert(), Update: t.Update(), Delete: t.Delete(), Select: t.Select()}, util.NewStatusOk()
}
