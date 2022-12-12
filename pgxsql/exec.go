package pgxsql

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fncall"
	"github.com/idiomatic-go/postgresql-adapter/dml"
	"github.com/idiomatic-go/postgresql-adapter/sql"
)

func ExecInsert(ctx context.Context, tag *CommandTag, sql string, values []any) (CommandTag, fncall.Status) {
	if fncall.IsContextContent(ctx) {
		return fncall.ProcessContextContent[CommandTag](ctx)
	}
	stmt, err := dml.WriteInsert(sql, values)
	if err != nil {
		return CommandTag{}, fncall.NewStatusError(err)
	}
	return ExecWithCommand(ctx, tag, stmt)
}

func ExecUpdate(ctx context.Context, tag *CommandTag, sql string, attrs []sql.Attr, where []sql.Attr) (CommandTag, fncall.Status) {
	if fncall.IsContextContent(ctx) {
		return fncall.ProcessContextContent[CommandTag](ctx)
	}
	stmt, err := dml.WriteUpdate(sql, attrs, where)
	if err != nil {
		return CommandTag{}, fncall.NewStatusError(err)
	}
	return ExecWithCommand(ctx, tag, stmt)
}

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, fncall.Status) {
	return ExecWithCommand(ctx, nil, sql, arguments)
}

func ExecWithCommand(ctx context.Context, tag *CommandTag, sql string, arguments ...any) (CommandTag, fncall.Status) {
	if fncall.IsContextContent(ctx) {
		return fncall.ProcessContextContent[CommandTag](ctx)
	}
	if dbClient == nil {
		return CommandTag{}, fncall.NewStatusInvalidArgument(errors.New("error on PostgreSQL exec call : dbClient is nil"))
	}
	// Transaction processing.
	txn, err0 := dbClient.Begin(ctx)
	if err0 != nil {
		return CommandTag{}, fncall.NewStatusError(err0)
	}
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		err0 := txn.Rollback(ctx)
		return CommandTag{}, fncall.NewStatusError(err, err0)
	}
	if tag != nil && t.RowsAffected() != tag.RowsAffected {
		err0 := txn.Rollback(ctx)
		return CommandTag{}, fncall.NewStatusError(errors.New(fmt.Sprintf("error exec statement [%v] : actual RowsAffected %v != expected RowsAffected %v", t.String(), t.RowsAffected(), tag.RowsAffected)), err0)
	}
	err = txn.Commit(ctx)
	if err != nil {
		return CommandTag{}, fncall.NewStatusError(err)
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected(), Insert: t.Insert(), Update: t.Update(), Delete: t.Delete(), Select: t.Select()}, fncall.NewStatusOk()
}
