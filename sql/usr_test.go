package sql

import "fmt"

func ExampleSanitize() {

	sc := Sanitize("")
	fmt.Printf("StatusCode  : %v\n", sc.Ok())

	sc = Sanitize("adfsdfe4fc&*4")
	fmt.Printf("StatusCode  : %v\n", sc.Ok())

	sc = Sanitize("test 1: /*")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 2: DROP   Table  ")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 3: DEL ETE FROM")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 4: - -")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 5: ;--")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 6: sa*/nitize 4 ;--")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	sc = Sanitize("test 7: of select  * froM customers")
	fmt.Printf("StatusCode  : %v %v\n", sc.Ok(), sc.String())

	//Output:
	//fail
}
