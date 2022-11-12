package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/postgresql-adapter/dml"
)

var execContentOverride = false

// TODO : verify string data to prevent SQL injection attacks
//        can/is this being done automatically via PostgreSQL?

func ExecInsert(ctx context.Context, sql string, values []any) (CommandTag, util.StatusCode) {
	if len(values) == 0 {
		return CommandTag{}, util.NewStatusInvalidArgument(errors.New("invalid argument: insert attributes list is empty"))
	}
	return Exec(ctx, dml.WriteInsert(sql, values))
}

func ExecUpdate(ctx context.Context, sql string, attrs ...util.Attr) (CommandTag, util.StatusCode) {
	if len(attrs) == 0 {
		return CommandTag{}, util.NewStatusInvalidArgument(errors.New("invalid argument: update attributes list is empty"))
	}
	return Exec(ctx, dml.WriteUpdate(sql, attrs))
}

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
