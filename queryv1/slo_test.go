package queryv1

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	v1 "github.com/idiomatic-go/core-types/corev1"
)

//go:embed resource/*
var fs embed.FS

func ExampleSLOEntryUnmarshal() {
	buf, err := fse.ReadFile(fs, "resource/code-5xx.json")
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Buf   : %v\n", string(buf))

	entry := v1.SLOEntry{}
	err = json.Unmarshal(buf, &entry)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Entry : %v\n", entry)

	//Output:
	// fail
}

func ExampleSLOQuery() {

}
