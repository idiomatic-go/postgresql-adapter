package pgxsql

// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbClient *pgxpool.Pool

func isClientStarted() bool {
	return dbClient != nil
}

var clientStartup vhost.MessageHandler = func(msg vhost.Message) {
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		logxt.LogPrintf("error reading configuration file from mounted file system : %v", err)
		vhost.SendErrorResponse(Uri)
		return
	}
	credentials := vhost.AccessCredentials(&msg)
	// Validate credentials
	if credentials == nil {
		logxt.LogPrintf("%v", "pgxsql credentials function is nil")
		vhost.SendErrorResponse(Uri)
		return
	}
	if !StartupDirect(credentials, m) {
		vhost.SendErrorResponse(Uri)
	}
}

func StartupDirect(credentials vhost.Credentials, config map[string]string) bool {
	url, ok := config[DatabaseURLKey]
	if !ok || url == "" {
		logxt.LogPrintf("database URL does not exist in map, or value is empty : %v", DatabaseURLKey)
		return false
	}

	// Determine if this is an override by interrogating the database url
	//if strings.Contains(url, DatabaseOverride) {
	//	overrideExec = nilExec
	//	overrideQuery = nilQuery
	//	return
	//}

	// Create connection string, pool and acquire connection
	s := connectString(url, credentials)
	if s == "" {
		return false
	}
	var err error
	dbClient, err = pgxpool.New(context.Background(), s)
	if err != nil {
		logxt.LogPrintf("unable to create connection pool : %v", err)
		return false
	}
	conn, err1 := dbClient.Acquire(context.Background())
	defer conn.Release()
	if err1 != nil {
		logxt.LogPrintf("unable to acquire connection from pool : %v", err1)
		ClientShutdown()
		return false
	}
	return true
}

func ClientShutdown() {
	if dbClient != nil {
		dbClient.Close()
		dbClient = nil
	}
}

func connectString(url string, credentials vhost.Credentials) string {
	username, password, err := credentials()
	if err != nil {
		logxt.LogPrintf("error accessing credentials: %v", err)
		return ""
	}
	return fmt.Sprintf(url, username, password)
}


