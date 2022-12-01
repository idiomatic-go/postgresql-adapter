package pgxsql

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/idiomatic-go/postgresql-adapter/dml"
	"github.com/idiomatic-go/postgresql-adapter/sql"
)

func ExecInsert(ctx context.Context, tag *CommandTag, sql string, values []any) (CommandTag, vhost.Status) {
	if vhost.IsContextContent(ctx) {
		return vhost.ProcessContextContent[CommandTag](ctx)
	}
	stmt, err := dml.WriteInsert(sql, values)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return ExecWithCommand(ctx, tag, stmt)
}

func ExecUpdate(ctx context.Context, tag *CommandTag, sql string, attrs []sql.Attr, where []sql.Attr) (CommandTag, vhost.Status) {
	if vhost.IsContextContent(ctx) {
		return vhost.ProcessContextContent[CommandTag](ctx)
	}
	stmt, err := dml.WriteUpdate(sql, attrs, where)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return ExecWithCommand(ctx, tag, stmt)
}

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, vhost.Status) {
	return ExecWithCommand(ctx, nil, sql, arguments)
}

func ExecWithCommand(ctx context.Context, tag *CommandTag, sql string, arguments ...any) (CommandTag, vhost.Status) {
	if vhost.IsContextContent(ctx) {
		return vhost.ProcessContextContent[CommandTag](ctx)
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
	if tag != nil && t.RowsAffected() != tag.RowsAffected {
		err0 := txn.Rollback(ctx)
		return CommandTag{}, vhost.NewStatusError(errors.New(fmt.Sprintf("error exec statement [%v] : actual RowsAffected %v != expected RowsAffected %v", t.String(), t.RowsAffected(), tag.RowsAffected)), err0)
	}
	err = txn.Commit(ctx)
	if err != nil {
		return CommandTag{}, vhost.NewStatusError(err)
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected(), Insert: t.Insert(), Update: t.Update(), Delete: t.Delete(), Select: t.Select()}, vhost.NewStatusOk()
}
