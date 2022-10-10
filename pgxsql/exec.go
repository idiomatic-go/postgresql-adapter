package pgxsql

import (
	"context"
	"github.com/idiomatic-go/common-lib/util"
)

type execFn func(ctx context.Context, sql string, arguments ...any) (CommandTag, error)

var overrideExec execFn

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	if overrideExec != nil {
		return overrideExec(ctx, sql, arguments)
	}
	t, err := dbclient.Exec(ctx, sql, arguments)
	if err != nil {
		util.LogPrintf("Error on execution : %v", err)
		return CommandTag{}, err
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected()}, nil
}
