package ddl

import (
	"context"
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/postgresql-adapter/pgxsql"
)

//go:embed resource/*
var fs embed.FS

func ExampleCreateRoles() {
	tag, err := ExecDDL("resource/create_roles.sql")

	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Tag   : %v\n", tag)

	//Output:
	//fail
}

func ExecDDL(name string) (pgxsql.CommandTag, error) {
	buf, err := fse.ReadFile(fs, name)
	if err != nil {
		return pgxsql.CommandTag{}, err
	}
	s := string(buf)
	return pgxsql.Exec(context.Background(), s)
}
