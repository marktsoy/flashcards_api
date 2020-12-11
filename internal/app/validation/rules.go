package validation

import (
	"fmt"
	"regexp"
)

// Rule ...
type Rule func(string, string, ...interface{}) (bool, string)

const errMinLen = "%v is too short, accepted length: %v"
const errMatches = "%v does not match required format"
const errEmail = "%v must be valid Email"

//MinLen ..
func MinLen(name string, value string, args ...interface{}) (bool, string) {
	if len(args) == 0 {
		panic("MinLen Rule is lacks N param")
	}
	n := args[0].(int)

	if len(value) >= n {
		return true, ""
	}
	message := args[1:]
	if len(message) > 0 {
		return false, message[0].(string)
	}
	return false, fmt.Sprintf(errMinLen, name, n)
}

// Matches regular expression
func Matches(name string, value string, args ...interface{}) (bool, string) {
	pattern := args[0].(string)
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(value) {
		return true, ""
	}
	message := args[1:]
	if len(message) > 0 {
		return false, message[0].(string)
	}
	return false, fmt.Sprintf(errMatches, name)
}

// Email ...
func Email(name string, value string, args ...interface{}) (bool, string) {
	emailPattern := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	message := args
	var msg string
	if len(message) > 0 {
		msg = message[0].(string)
	} else {
		msg = fmt.Sprintf(errEmail, name)
	}

	return Matches(name, value, emailPattern, msg)
}
