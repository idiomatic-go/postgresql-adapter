package pgxsql

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type rows struct {
	pgxRows pgx.Rows
	fd      []FieldDescription
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

func (r *rows) FieldDescriptions() []FieldDescription {
	if r == nil {
		return nil
	}
	return r.fd
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

func Query(ctx context.Context, sql string, arguments ...any) (Rows, util.StatusCode) {
	if sql == ExecContentSql {
		tag, err := fse.ProcessContent[Rows](ctx)
		return tag, util.NewStatusInvalidArgument(err)
	}
	if dbClient == nil {
		sc := util.NewStatusInvalidArgument(errors.New("error on PostgreSQL database query call: dbClient is nil"))
		logxt.LogPrintf("%v", sc)
		return nil, sc
	}
	pgxRows, err := dbClient.Query(ctx, sql, arguments)
	if err != nil {
		logxt.LogPrintf("error on PostgreSQL database query call: %v", err)
		return nil, util.NewStatusError(err)
	}
	return &rows{pgxRows: pgxRows, fd: fieldDescriptions(pgxRows.FieldDescriptions())}, util.NewStatusOk()
}

func nilQuery(ctx context.Context, sql string, arguments ...any) (Rows, error) {
	return nil, nil
}

func fieldDescriptions(fields []pgconn.FieldDescription) []FieldDescription {
	var result []FieldDescription
	for _, f := range fields {
		result = append(result, FieldDescription{Name: f.Name,
			TableOID:             f.TableOID,
			TableAttributeNumber: f.TableAttributeNumber,
			DataTypeOID:          f.DataTypeOID,
			DataTypeSize:         f.DataTypeSize,
			TypeModifier:         f.TypeModifier,
			Format:               f.Format})
	}
	return result
}
