package main

import (
	"fmt"
	"strings"
)

var bunchaEmails = []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
var bunchaEmails2 = []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}

func numUniqueEmails(emails []string) int {
	var emailCount int
	var uniqueEmails = map[string]map[string]struct{}{}

	for _, email := range emails {
		fmt.Println("starting with: ", email)
		// separate email address into localname and domain strings
		splitEmail := strings.Split(email, "@")

		name := splitEmail[0]
		domain := splitEmail[1]

		// count how many keys are in each of our uniqueEmails child maps
		if _, ok := uniqueEmails[domain]; !ok {
			uniqueEmails[domain] = map[string]struct{}{}
		}

		// remove . characters
		// ignore everything to the right of the +
		//fmt.Println("trimtest", strings.(name, ".")[1])
		noDots := strings.Replace(name, ".", "", -1)
		fmt.Println("noDots: ", noDots)
		baseName := strings.Split(noDots, "+")[0]

		fmt.Println("baseName plus domain:", baseName, domain)

		if _, ok := uniqueEmails[domain][baseName]; !ok {
			uniqueEmails[domain][baseName] = struct{}{}
		}
	}

	for _, uniqueDomain := range uniqueEmails {
		emailCount += len(uniqueDomain)
	}

	return emailCount
}

func main() {
	fmt.Println(numUniqueEmails(bunchaEmails))
	fmt.Println(numUniqueEmails(bunchaEmails2))
}
