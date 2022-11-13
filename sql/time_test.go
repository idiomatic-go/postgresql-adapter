package sql

import (
	"fmt"
	"time"
)

func _ExampleFmtTimestamp() {
	s := FmtTimestamp(time.Now())

	fmt.Printf("Timestamp : %v\n", s)

	//Output:
	//fail
}
