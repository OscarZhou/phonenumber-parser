package main

import (
	"fmt"
	"phonenumber-parser/parser"
)

func main() {
	// phoneNumber := "+640211376664"
	// phoneNumber := "+4402012341234"
	// phoneNumber := "+370860112345"
	// phoneNumber := "+8618809817047"
	// phoneNumber := "+1 (989) 341-8338"
	// phoneNumber := "+213 982402746"
	// phoneNumber := "+61 438608762"
	// phoneNumber := "+44 1463215569"
	phoneNumber := "+56 945950540"

	e164, err := parser.ParsePhoneNumber(phoneNumber)
	if err != nil {
		panic(err)
	}

	fmt.Println(e164)
}
