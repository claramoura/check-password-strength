package rule

import (
	"regexp"
)

// Rule 1: password is at least n characters long.
func MinSize(password string, n int) bool {

	runes := []rune(password)

	var size int = len(runes)

	if size < n {
		return false
	}

	return true
}

// Rule 2: password contains at least n uppercase letters.
func MinUppercase(password string, n int) bool {

	var re = regexp.MustCompile(`[A-Z]`)

	var countUppercase int = len(re.FindAllString(password, -1))

	if countUppercase < n {
		return false
	}

	return true
}

// Rule 3: password contains at least n lowercase letters.
func MinLowercase(password string, n int) bool {

	var re = regexp.MustCompile(`[a-z]`)

	var countLowercase int = len(re.FindAllString(password, -1))

	if countLowercase < n {
		return false
	}

	return true
}

// Rule 4: password contains at least n digits (0-9).
func MinDigit(password string, n int) bool {

	var re = regexp.MustCompile(`\d`)

	var countDigits int = len(re.FindAllString(password, -1))

	if countDigits < n {
		return false
	}

	return true
}

// Rule 5: password contains at least n special characters, i.e., neither a letter nor a digit.
func MinSpecialChars(password string, n int) bool {

	var re = regexp.MustCompile(`[!@#\$%\^&\*\(\)\-\+\\/\{\}\[\]]`)

	var countSpecialChars int = len(re.FindAllString(password, -1))

	if countSpecialChars < n {
		return false
	}

	return true
}

// Rule 6: password must not contain sequences of the same characters -- 'aba' is allowed, but 'aab' is not.
func NoRepeated(password string) bool {

	runes := []rune(password)

	for i := 0; i < len(runes); i++ {
		if (i+1) != len(runes) && runes[i] == runes[i+1] {
			return false
		}
	}

	return true
}
