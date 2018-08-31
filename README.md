# phonenumber-parser
Normalize any forms of phone number all around the world



## Reference Documents

1. [E.164](https://en.wikipedia.org/wiki/E.164)  
2. [ISO3166](https://en.wikipedia.org/wiki/ISO_3166)  
3. [List of country calling codes](https://en.wikipedia.org/wiki/List_of_country_calling_codes)  



## International Phone Number Format  

**E.164 phone number format contains**

`[+][country code][area code][local phone number]`  
*Description:*

* *+* - plus sign
* country code - internaltional country code
* area code/ national destination code - code without leading 0
* phone number - local phone number


*Example*

| Country   | Local phone number | E.164 formatted number |
|-----------|--------------------|------------------------|
| USA       | 415 123 1234       | +14151231234           |
| UK        | 020 1234 1234      | +442012341234          |
| Lithuania | 8 601 12345        | +37060112345           |
| NZ        | 021 123 1234       | +64211231234           |
