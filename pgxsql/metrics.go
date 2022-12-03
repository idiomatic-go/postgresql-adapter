package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Stat() (*pgxpool.Stat, vhost.Status) {
	if dbClient == nil {
		return nil, vhost.NewStatusInvalidArgument(errors.New("error on PostgreSQL stat call : dbClient is nil"))
	}
	return dbClient.Stat(), nil
}

func Ping(ctx context.Context) vhost.Status {
	if dbClient == nil {
		return vhost.NewStatusInvalidArgument(errors.New("error on PostgreSQL pingc call : dbClient is nil"))
	}
	err := dbClient.Ping(ctx)
	if err != nil {
		return vhost.NewStatusError(err)
	}
	return vhost.NewStatusOk()
}
