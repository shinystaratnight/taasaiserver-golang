package utils

import "strings"

func Capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}
