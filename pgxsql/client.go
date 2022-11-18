package pgxsql

// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbClient *pgxpool.Pool

func isClientStarted() bool {
	return dbClient != nil
}

var clientStartup eventing.MessageHandler = func(msg eventing.Message) {
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		vhost.LogPrintf("error reading configuration file from mounted file system : %v", err)
		eventing.SendErrorResponse(Uri, vhost.StatusInternal)
		return
	}
	credentials := vhost.AccessCredentials(&msg)
	// Validate credentials
	if credentials == nil {
		vhost.LogPrintf("%v", "pgxsql credentials function is nil")
		eventing.SendErrorResponse(Uri, vhost.StatusInternal)
		return
	}
	if !StartupDirect(m, credentials) {
		eventing.SendErrorResponse(Uri, vhost.StatusInternal)
	}
}

func StartupDirect(config map[string]string, credentials vhost.Credentials) bool {
	if isClientStarted() {
		vhost.LogPrintf("%v", "database client is already running")
		return false
	}
	url, ok := config[DatabaseURLKey]
	if !ok || url == "" {
		vhost.LogPrintf("database URL does not exist in map, or value is empty : %v", DatabaseURLKey)
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
		vhost.LogPrintf("unable to create connection pool : %v", err)
		return false
	}
	conn, err1 := dbClient.Acquire(context.Background())
	defer conn.Release()
	if err1 != nil {
		vhost.LogPrintf("unable to acquire connection from pool : %v", err1)
		clientShutdown()
		return false
	}
	return true
}

func clientShutdown() {
	if dbClient != nil {
		dbClient.Close()
		dbClient = nil
	}
}

func connectString(url string, credentials vhost.Credentials) string {
	username, password, err := credentials()
	if err != nil {
		vhost.LogPrintf("error accessing credentials: %v", err)
		return ""
	}
	return fmt.Sprintf(url, username, password)
}
