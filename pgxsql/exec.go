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
		err := errors.New("error on PostgreSQL exec call : dbClient is nil")
		logxt.LogPrintf("%v", err)
		return CommandTag{}, err
	}
	// Transaction processing.
	txn, err0 := dbClient.Begin(ctx)
	if err0 != nil {
		logxt.LogPrintf("error on PostgreSQL begin transaction : %v", err0)
		return CommandTag{}, err0
	}
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		logxt.LogPrintf("error on PostgreSQL exec call : %v", err)
		err0 := txn.Rollback(ctx)
		if err0 != nil {
			logxt.LogPrintf("error on PostgreSQL rollback transaction call : %v", err0)
		}
		return CommandTag{}, err
	}
	err = txn.Commit(ctx)
	if err != nil {
		logxt.LogPrintf("error on PostgreSQL commit transaction call : %v", err)
		return CommandTag{}, err
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected()}, nil
}

func nilExec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	return CommandTag{}, nil
}
