package queryv1

import (
	"embed"
	"fmt"
	v1 "github.com/idiomatic-go/core-types/corev1"
)

//go:embed resource/*
var content embed.FS

func ExampleSLOEntry() {
	entry := v1.SLOEntry{}
	fmt.Printf("Entry : %v\n", entry)

	//Output:
	// fail
}
