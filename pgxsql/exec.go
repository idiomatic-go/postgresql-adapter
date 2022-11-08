package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
)

type execFn func(ctx context.Context, sql string, arguments ...any) (CommandTag, error)

var overrideExec execFn

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	if sql == ExecContentSql {
		return fse.ProcessContent[CommandTag](ctx)
	}
	if dbClient == nil {
		err := errors.New("error on database execution call : dbClient is nil")
		logxt.LogPrintf("%v", err)
		return CommandTag{}, err
	}
	// TODO : transaction processing.
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		logxt.LogPrintf("error on database execution call : %v", err)
		return CommandTag{}, err
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected()}, nil
}

func nilExec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	return CommandTag{}, nil
}
