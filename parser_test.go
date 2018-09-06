package parser

import "testing"

func TestParser(t *testing.T) {
	expected := E164{
		CountryCode:    "86",
		TruckPrefixes:  "",
		AreaCode:       "18",
		DailingNumber:  "8618812345678",
		DomesticNumber: "18812345678",
		Alpha2:         "CN",
		Alpha3:         "CHN",
	}

	phoneNumber := "+86 188 1234 5678"
	get, err := ParsePhoneNumber(phoneNumber)
	if err != nil {
		t.Errorf("expected:%v\n get:%v", expected, get)
	}

	expected = E164{
		CountryCode:    "64",
		TruckPrefixes:  "0",
		AreaCode:       "2",
		DailingNumber:  "64212345678",
		DomesticNumber: "0212345678",
		Alpha2:         "NZ",
		Alpha3:         "NZL",
	}

	phoneNumber = "+64 (02) 1234 5678"
	get, err = ParsePhoneNumber(phoneNumber)
	if err != nil {
		t.Errorf("expected:%v\n get:%v", expected, get)
	}

}
