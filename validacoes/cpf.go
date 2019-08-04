package main

import (
	"fmt"
	"regexp"
)

const (
	Regex = `^[0-9]{3}.?[0-9]{3}.?[0-9]{3}-?[0-9]{2}$`
)

func main() {
	re := regexp.MustCompile(Regex)
	tests := []string{
		"788.308.710-01",
		"78830871001",
		"788308710-01",
		"788.308.710.01",
		"788 308 710 01",
		"788-308-710-01",
	}
	for _, test := range tests {
		fmt.Println(re.MatchString(test))
	}
}
