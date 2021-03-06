// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/BlueMaverick/EPAMassignment/design
// --out=$(GOPATH)/src/github.com/BlueMaverick/EPAMassignment
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v", paramaccountID)
}

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", paramaccountID, parambottleID)
}
