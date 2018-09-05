package parser

import (
	"errors"
	"regexp"
	"strings"
)

// E164 follows the E.164 standard
// When calling a number, country code should be prefixed
//
type E164 struct {
	CountryCode    string // 64
	TruckPrefixes  string // 0
	AreaCode       string // 21
	PhoneNumber    string // 1376664
	DailingNumber  string // CountryCode + AreaCode + PhoneNumber
	DomesticNumber string // TruckPrefixes + PhoneNumber
	Alpha2         string // NZ
	Alpha3         string // NZL
}

func ParsePhoneNumber(phoneNumber string) (E164, error) {
	var entity E164
	if phoneNumber == "" {
		return E164{}, errors.New("phone number is required")
	}

	tmpPhoneNumber := strings.Replace(phoneNumber, " ", "", -1)
	tmpPhoneNumber = strings.Replace(tmpPhoneNumber, "(", "", -1)
	tmpPhoneNumber = strings.Replace(tmpPhoneNumber, ")", "", -1)
	tmpPhoneNumber = strings.Replace(tmpPhoneNumber, "-", "", -1)
	tmpPhoneNumber = strings.Replace(tmpPhoneNumber, "+", "", -1)
	iso3166 := GetISO3166()
	current := ISO3166{}
	for _, v := range iso3166 {
		if regexp.MustCompile(`\A` + v.CountryCode + `[\d*]`).MatchString(tmpPhoneNumber) {
			entity.CountryCode = v.CountryCode
			entity.Alpha2 = v.Alpha2
			entity.Alpha3 = v.Alpha3
			tmpPhoneNumber = strings.Replace(tmpPhoneNumber, v.CountryCode, "", 1)
			truckCode := filterTruckCode(entity.Alpha2)
			entity.TruckPrefixes = truckCode
			tmpPhoneNumber = tmpPhoneNumber[len(truckCode):]
			current = v
			break
		}
	}

	if current.CountryCode != "" {
		for _, areaCode := range current.MobileBeginWith {
			if strings.HasPrefix(tmpPhoneNumber, areaCode) {
				entity.AreaCode = areaCode
				tmpPhoneNumber = strings.Replace(tmpPhoneNumber, areaCode, "", 1)
				entity.PhoneNumber = tmpPhoneNumber
				entity.DailingNumber = entity.CountryCode + entity.AreaCode + entity.PhoneNumber
				entity.DomesticNumber = entity.TruckPrefixes + entity.PhoneNumber
				break
			}
		}
	}
	return entity, nil
}

func filterTruckCode(country string) string {
	truckCode := ""
	switch country {
	case "MX":
		truckCode = "01"
	case "AZ", "RU", "KZ", "TC", "UZ", "BY":
		truckCode = "8"
	}
	return truckCode
}
