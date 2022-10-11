package pgxsql

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

var dbclient *pgxpool.Pool

var credentials vhost.Credentials

var clientStartup util.Func = func() {
	// Read the configuration map and database Url first
	m, url := readConfiguration()
	if m == nil {
		vhost.SendErrorResponse(Uri)
		return
	}

	// Determine if this is an override by interrogating the database url
	if strings.Contains(url, DatabaseOverride) {
		overrideExec = devExec
		overrideQuery = nilQuery
		return
	}

	// Validate credentials
	if credentials == nil {
		util.LogPrintf("%v", "pgxsql credentials function is nil")
		vhost.SendErrorResponse(Uri)
		return
	}
	
	// Create connection string, pool and acquire connection
	s := connectString(url)
	if s == "" {
		vhost.SendErrorResponse(Uri)
		return
	}
	dbclient, err := pgxpool.New(context.Background(), s)
	if err != nil {
		util.LogPrintf("unable to create connection pool : %v", err)
		vhost.SendErrorResponse(Uri)
		return
	}
	conn, err1 := dbclient.Acquire(context.Background())
	defer conn.Release()
	if err1 != nil {
		util.LogPrintf("unable to acquire connection from pool : %v", err1)
		vhost.SendErrorResponse(Uri)
		shutdown()
		return
	}
}

func clientShutdown() {
	if dbclient != nil {
		dbclient.Close()
	}
}

func connectString(url string) string {
	username, password, err := credentials()
	if err != nil {
		util.LogPrintf("error accessing credentials: %v", err)
		return ""
	}
	return fmt.Sprintf(url, username, password)
}

func readConfiguration() (map[string]string, string) {
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		util.LogPrintf("error reading configuration file from mounted file system : %v", err)
		return nil, ""
	}
	s, ok := m[DatabaseURLKey]
	if !ok || s == "" {
		util.LogPrintf("database URL does not exist in map, or value is empty : %v", DatabaseURLKey)
		return nil, ""
	}
	return m, s
}
