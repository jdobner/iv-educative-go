package main

import "strings"

func runFunctions() {
	upper, lower := toUpperAndLower("Hello", "World")
	println(upper, lower)
	println(toLowerAndUpper("Hello", "World"))
}

func toUpperAndLower(s1, s2 string) (upper, lower string) {
	upper = strings.ToUpper(s1)
	lower = strings.ToLower(s2)
	return
}

func toLowerAndUpper(s1, s2 string) (upper, lower string) {
	return strings.ToLower(s1), strings.ToUpper(s2)
}
