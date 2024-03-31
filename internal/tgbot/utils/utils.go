package utils

import "strings"

func EscapeHTML(s string) string {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
	return replacer.Replace(s)
}
