package util

import (
	"regexp"
	"strings"
)

func StartsWithCountryCode(phoneNumber string, countryCode string) bool {
	return strings.HasPrefix(phoneNumber, countryCode)
}

func HasCharacters(input string) bool {
	// Regular expressions for uppercase letter, number, and special character
	uppercaseRegex := `[A-Z]`
	numberRegex := `[0-9]`
	specialCharRegex := `[^a-zA-Z0-9]`

	// Check if the input string matches all three regex patterns
	return regexp.MustCompile(uppercaseRegex).MatchString(input) &&
		regexp.MustCompile(numberRegex).MatchString(input) &&
		regexp.MustCompile(specialCharRegex).MatchString(input)
}
