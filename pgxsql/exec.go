package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
)

var execContentOverride = false

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, util.StatusCode) {
	if execContentOverride {
		tag, err := fse.ProcessContent[CommandTag](ctx)
		return tag, util.NewStatusInvalidArgument(err)
	}
	if dbClient == nil {
		sc := util.NewStatusInvalidArgument(errors.New("error on PostgreSQL exec call : dbClient is nil"))
		logxt.LogPrintf("%v", sc)
		return CommandTag{}, sc
	}
	// Transaction processing.
	txn, err0 := dbClient.Begin(ctx)
	if err0 != nil {
		logxt.LogPrintf("error on PostgreSQL begin transaction : %v", err0)
		return CommandTag{}, util.NewStatusError(err0)
	}
	t, err := dbClient.Exec(ctx, sql, arguments)
	if err != nil {
		logxt.LogPrintf("error on PostgreSQL exec call : %v", err)
		err0 := txn.Rollback(ctx)
		if err0 != nil {
			logxt.LogPrintf("error on PostgreSQL rollback transaction call : %v", err0)
		}
		return CommandTag{}, util.NewStatusError(err, err0)
	}
	err = txn.Commit(ctx)
	if err != nil {
		logxt.LogPrintf("error on PostgreSQL commit transaction call : %v", err)
		return CommandTag{}, util.NewStatusError(err)
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected()}, util.NewStatusOk()
}
