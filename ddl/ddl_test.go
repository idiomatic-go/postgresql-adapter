package ddl

import (
	"context"
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fncall"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/postgresql-adapter/pgxsql"
)

//go:embed resource/*
var fs embed.FS

const (
	databaseUrl = "postgresql://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
)

func createConfig(url string) map[string]string {
	return map[string]string{pgxsql.DatabaseURLKey: url}
}

func ExampleCreateDatabase() {
	fmt.Printf("%v\n", "create database")

	status := pgxsql.ClientStartup(createConfig(databaseUrl), nil)
	if status.IsError() {
		fmt.Printf("%v\n", status)
		return
	}
	defer pgxsql.ClientShutdown()

	status = createRoles()
	if status.IsError() {
		fmt.Printf("%v\n", status)
		return
	}

	//Output:
	//fail
}

func createRoles() fncall.Status {
	tag, status := execDDL("resource/create_roles.sql")
	if status.IsError() {
		return status
	}
	fmt.Printf("%v\n", tag)
	return fncall.NewStatusOk()
}

func execDDL(name string) (pgxsql.CommandTag, fncall.Status) {
	buf, err := fse.ReadFile(fs, name)
	if err != nil {
		return pgxsql.CommandTag{}, fncall.NewStatusError(err)
	}
	s := string(buf)
	return pgxsql.Exec(context.Background(), s)
}

func _ExampleDropDatabase() {
	fmt.Printf("%v\n", "drop database")

	status := pgxsql.ClientStartup(createConfig(databaseUrl), nil)
	if status.IsError() {
		fmt.Printf("%v\n", status)
		return
	}
	defer pgxsql.ClientShutdown()

	status = execScripts(nil)
	if status.IsError() {
		fmt.Printf("%v\n", status)
		return
	}
	//Output:
	//fail
}

func execScripts(scripts []string) fncall.Status {
	var tag pgxsql.CommandTag
	var status fncall.Status

	if scripts == nil || len(scripts) == 0 {
		return fncall.NewStatusOk()
	}
	for _, cmd := range scripts {
		tag, status = pgxsql.Exec(context.Background(), cmd)
		if status.IsError() {
			return status
		}
		fmt.Printf("%v\n", tag)
	}
	return fncall.NewStatusOk()
}
