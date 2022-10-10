package pgxsql

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbclient *pgxpool.Pool

var credentials vhost.Credentials

var clientStartup util.Func = func() {
	if credentials == nil {
		util.LogPrintf("%v", "pgxsql Credentials function is nil")
		vhost.SendErrorResponse(Uri)
		return
	}
	s := connectString()
	if s == "" {
		vhost.SendErrorResponse(Uri)
		return
	}
	dbclient, err := pgxpool.New(context.Background(), s)
	if err != nil {
		util.LogPrintf("Unable to create connection pool : %v", err)
		vhost.SendErrorResponse(Uri)
		return
	}
	conn, err1 := dbclient.Acquire(context.Background())
	if err1 != nil {
		util.LogPrintf("Unable to acquire connection from pool : %v", err1)
		vhost.SendErrorResponse(Uri)
		shutdown()
		return
	}
	conn.Release()
}

func clientShutdown() {
	if dbclient != nil {
		dbclient.Close()
	}
}

func connectString() string {
	t := connectStringTemplate()
	username, password, err := credentials()
	if err != nil {
		util.LogPrintf("error on accessing credentials: %v", err)
	}
	return fmt.Sprintf(t, username, password)
}

func connectStringTemplate() string {
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		util.LogPrintf("Connection creation file mount access error : %v", err)
		return ""
	}
	s, ok := m[DatabaseURLKey]
	if !ok || s == "" {
		util.LogPrintf("Database URL does not exist in map, or value is empty : %v", DatabaseURLKey)
		return ""
	}
	return s
}
