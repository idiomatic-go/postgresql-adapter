package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fncall"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Stat() (*pgxpool.Stat, fncall.Status) {
	if dbClient == nil {
		return nil, fncall.NewStatusInvalidArgument(errors.New("error on PostgreSQL stat call : dbClient is nil"))
	}
	return dbClient.Stat(), nil
}

func Ping(ctx context.Context) fncall.Status {
	if dbClient == nil {
		return fncall.NewStatusInvalidArgument(errors.New("error on PostgreSQL pingc call : dbClient is nil"))
	}
	err := dbClient.Ping(ctx)
	if err != nil {
		return fncall.NewStatusError(err)
	}
	return fncall.NewStatusOk()
}
