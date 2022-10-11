package pgxsql

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"strings"
)

type execFn func(ctx context.Context, sql string, arguments ...any) (CommandTag, error)

var overrideExec execFn

var database = make(map[string]*util.AnyTable, 1)

func Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	if overrideExec != nil {
		return overrideExec(ctx, sql, arguments)
	}
	t, err := dbclient.Exec(ctx, sql, arguments)
	if err != nil {
		util.LogPrintf("error on database execution call : %v", err)
		return CommandTag{}, err
	}
	return CommandTag{Sql: t.String(), RowsAffected: t.RowsAffected()}, nil
}

func devExec(ctx context.Context, sql string, arguments ...any) (CommandTag, error) {
	if sql == "" {
		return CommandTag{}, nil
	}
	// The command text is of the form : [http-method]' '[tablename/resource-id]
	method, tablename, resource, err := parseCommand(sql)
	if err != nil {
		util.LogPrintf("error on database execution command parse : %v", err)
		return CommandTag{}, err
	}
	table := database[tablename]
	switch method {
	case "GET":
		result, ok := table.Get(resource)
		count := 0
		if ok {
			count = 1
		}
		return CommandTag{Sql: "GET", RowsAffected: int64(count), Result: result}, nil
	case "PUT":
		ok := table.Put(arguments[0])
		count := 0
		if ok {
			count = 1
		}
		return CommandTag{Sql: "PUT", RowsAffected: int64(count), Result: nil}, nil
	case "DELETE":
		ok := table.Delete(resource)
		count := 0
		if ok {
			count = 1
		}
		return CommandTag{Sql: "DELETE", RowsAffected: int64(count), Result: nil}, nil
	case "POST":
		if table == nil {
			eq, ok := arguments[0].(util.IsEqual)
			if !ok {
				util.LogPrintf("%v", "error on table creation: IsEqual function is not available")
				return CommandTag{}, err
			}
			table = util.CreateAnyTable(eq)
			database[tablename] = table
		}
		return CommandTag{Sql: "POST", RowsAffected: 0, Result: nil}, nil
	}
	return CommandTag{}, errors.New(fmt.Sprintf("invalid method: %v", method))
}

func parseCommand(sql string) (method string, tablename string, resource string, err error) {
	tokens := strings.Split(sql, " ")
	if len(tokens) < 2 {
		return "", "", "", errors.New("invalid command syntax : missing component")
	}
	method = tokens[0]
	if method == "" {
		return "", "", "", errors.New("invalid method: empty")
	}
	parts := strings.Split(tokens[1], "/")
	tablename = parts[0]
	if tablename == "" {
		return "", "", "", errors.New("invalid tablename: empty")
	}
	//if resource == "" {
	//	return "", "", "", errors.New("invalid resource: empty")
	//}
	if len(parts) > 1 {
		resource = parts[1]
	}
	return method, tablename, resource, nil
}
