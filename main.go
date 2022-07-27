package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

/*
1 . Write a function with this signature:

	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		// your code goes here
	}

	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
	Aspiration.com
	You hand me back
	asPirAtiOn.cOm
	Please note:
	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].
*/

func isSpecialCharacter(s string) bool {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(s, " ") == " "
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// change entire string to lower case
	s = strings.ToLower(s)

	count := 1
	for i := range s {
		if isSpecialCharacter(string(s[i])) {
			continue
		} else if count > len(s) {
			break
		}

		if count%3 == 0 {
			s = s[:i] + strings.ToUpper(string(s[i])) + s[i+1:]
		}

		count++
	}

	return s
}

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}
