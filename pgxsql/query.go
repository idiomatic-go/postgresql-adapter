package pgxsql

import (
	"context"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/jackc/pgx/v5"
)

type rows struct {
	pgxRows pgx.Rows
}

func (r *rows) Close() {
	if r != nil {
		r.pgxRows.Close()
	}
}

func (r *rows) Err() error {
	if r == nil {
		return nil
	}
	return r.pgxRows.Err()
}

func (r *rows) CommandTag() CommandTag {
	if r == nil {
		return CommandTag{}
	}
	t := r.pgxRows.CommandTag()
	return CommandTag{RowsAffected: t.RowsAffected(), Sql: t.String()}
}

func (r *rows) Next() bool {
	if r == nil {
		return false
	}
	return r.pgxRows.Next()
}

func (r *rows) Scan(dest ...any) error {
	if r == nil {
		return nil
	}
	return r.pgxRows.Scan(dest)
}

func (r *rows) Values() ([]any, error) {
	if r == nil {
		return nil, nil
	}
	return r.pgxRows.Values()
}

func (r *rows) RawValues() [][]byte {
	if r == nil {
		return nil
	}
	return r.pgxRows.RawValues()
}

type queryFn func(ctx context.Context, sql string, arguments ...any) (Rows, error)

var overrideQuery queryFn

func Query(ctx context.Context, sql string, arguments ...any) (Rows, error) {
	if overrideQuery != nil {
		return overrideQuery(ctx, sql, arguments)
	}
	pgxRows, err := dbclient.Query(ctx, sql, arguments)
	if err != nil {
		util.LogPrintf("Error on query : %v", err)
		return nil, err
	}
	return &rows{pgxRows: pgxRows}, nil
}

func nilQuery(ctx context.Context, sql string, arguments ...any) (Rows, error) {
	return nil, nil
}
