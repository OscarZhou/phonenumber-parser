package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	Original    string
	Formatted   string
	CountryCode string
	AreaCode    string
}

// func ParsePhoneNumberWithCountryCode(phoneNumber, countryCode string) (formmated string, err error) {

// }

// func ParsePhoneNumberWithAreaCode(phoneNumber, countryCode string) (formatted string, err error) {

// }

// func ParsePhoneNumber2CountryCode(phoneNumber string) (countryCode string, err error) {

// }

func ParsePhoneNumber(phoneNumber string) (*PhoneNumber, error) {
	var entity PhoneNumber
	if phoneNumber == "" {
		return nil, errors.New("phone number is required")
	}

	entity.Original = phoneNumber
	iso3166 := GetISO3166()
	for _, v := range iso3166 {
		pattern := strings.Replace(phoneNumber, "+", "", -1)
		if regexp.MustCompile(`\A` + v.CountryCode + `[\d*]`).MatchString(pattern) {
			fmt.Printf("pattern is %s\n", pattern)
			pattern = strings.Replace(pattern, v.CountryCode, "", 1)
			fmt.Printf("pattern is %s\n", pattern)
			for _, mobileCode := range v.MobileBeginWith {
				if regexp.MustCompile(`\A0?` + mobileCode).MatchString(pattern) {
					entity.CountryCode = v.Alpha2
					entity.AreaCode = v.CountryCode
					if pattern[0] == '0' {
						entity.Formatted = "+" + v.CountryCode + pattern[1:]
					} else {
						entity.Formatted = "+" + v.CountryCode + pattern
					}
					fmt.Printf("formatted is %s\n", entity.Formatted)
					break
				}
			}
		}
	}

	return &entity, nil
}
