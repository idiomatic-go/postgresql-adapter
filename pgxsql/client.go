package pgxsql

// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbClient *pgxpool.Pool

func isClientStarted() bool {
	return dbClient != nil
}

var clientStartup eventing.MessageHandler = func(msg eventing.Message) {
	if isClientStarted() {
		return
	}
	m, err := vhost.ReadMap(ConfigFileName)
	if err != nil {
		logxt.Printf("error reading configuration file from mounted file system : %v\n", err)
		vhost.SendStartupFailureResponse(Uri)
		return
	}
	credentials := vhost.AccessCredentials(&msg)
	// Validate credentials
	if credentials == nil {
		logxt.Printf("%v\n", "pgxsql credentials function is nil")
		vhost.SendStartupFailureResponse(Uri)
		return
	}
	status := ClientStartup(m, credentials)
	if status.IsError() {
		logxt.Printf("%v\n", status)
		vhost.SendStartupFailureResponse(Uri)
		return
	}
	vhost.SendStartupSuccessfulResponse(Uri)
}

func ClientStartup(config map[string]string, credentials vhost.Credentials) vhost.Status {
	if isClientStarted() {
		return vhost.NewStatusOk()
	}
	// Access database URL
	url, ok := config[DatabaseURLKey]
	if !ok || url == "" {
		return vhost.NewStatusError(errors.New(fmt.Sprintf("database URL does not exist in map, or value is empty : %v\n", DatabaseURLKey)))
	}

	// Determine if this is an override by interrogating the database url
	//if strings.Contains(url, DatabaseOverride) {
	//	overrideExec = nilExec
	//	overrideQuery = nilQuery
	//	return
	//}

	// Create connection string with credentials
	s, status := connectString(url, credentials)
	if status.IsError() {
		return status
	}
	// Create pooled client and acquire connection
	var err error
	dbClient, err = pgxpool.New(context.Background(), s)
	if err != nil {
		return vhost.NewStatusError(errors.New(fmt.Sprintf("unable to create connection pool : %v\n", err)))
	}
	conn, err1 := dbClient.Acquire(context.Background())
	if err1 != nil {
		ClientShutdown()
		return vhost.NewStatusError(errors.New(fmt.Sprintf("unable to acquire connection from pool : %v\n", err1)))
	}
	conn.Release()
	return vhost.NewStatusOk()
}

func ClientShutdown() {
	if dbClient != nil {
		dbClient.Close()
		dbClient = nil
	}
}

func connectString(url string, credentials vhost.Credentials) (string, vhost.Status) {
	if credentials == nil {
		return url, vhost.NewStatusOk()
	}
	username, password, err := credentials()
	if err != nil {
		return "", vhost.NewStatusError(errors.New(fmt.Sprintf("error accessing credentials: %v\n", err)))
	}
	return fmt.Sprintf(url, username, password), vhost.NewStatusOk()
}
