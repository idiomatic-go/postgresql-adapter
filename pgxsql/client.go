package pgxsql

// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/common-lib/vhost"
	start "github.com/idiomatic-go/common-lib/vhost/startup"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

var dbclient *pgxpool.Pool

var clientStartup util.Func = func() {
	// Read the configuration map and database Url first
	m, url := readConfiguration()
	if m == nil {
		start.SendErrorResponse(Uri)
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
		logxt.LogPrintf("%v", "pgxsql credentials function is nil")
		start.SendErrorResponse(Uri)
		return
	}

	// Create connection string, pool and acquire connection
	s := connectString(url)
	if s == "" {
		start.SendErrorResponse(Uri)
		return
	}
	dbclient, err := pgxpool.New(context.Background(), s)
	if err != nil {
		logxt.LogPrintf("unable to create connection pool : %v", err)
		start.SendErrorResponse(Uri)
		return
	}
	conn, err1 := dbclient.Acquire(context.Background())
	defer conn.Release()
	if err1 != nil {
		logxt.LogPrintf("unable to acquire connection from pool : %v", err1)
		start.SendErrorResponse(Uri)
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
		logxt.LogPrintf("error accessing credentials: %v", err)
		return ""
	}
	return fmt.Sprintf(url, username, password)
}

func readConfiguration() (map[string]string, string) {
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		logxt.LogPrintf("error reading configuration file from mounted file system : %v", err)
		return nil, ""
	}
	s, ok := m[DatabaseURLKey]
	if !ok || s == "" {
		logxt.LogPrintf("database URL does not exist in map, or value is empty : %v", DatabaseURLKey)
		return nil, ""
	}
	return m, s
}
