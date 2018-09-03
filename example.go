package main

import "phonenumber-parser/parser"

func main() {
	// phoneNumber := "+640211376664"
	// phoneNumber := "+4402012341234"
	// phoneNumber := "+370860112345"
	phoneNumber := "+8618809817047"
	_, err := parser.ParsePhoneNumber(phoneNumber)
	if err != nil {
		panic(err)
	}
}
