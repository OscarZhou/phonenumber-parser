package main

import (
	"fmt"
	parser "phonenumber-parser"
)

func main() {
	phoneNumber := "+56 945950540"
	e164, err := parser.ParsePhoneNumber(phoneNumber)
	if err != nil {
		panic(err)
	}

	fmt.Println(e164)
}
