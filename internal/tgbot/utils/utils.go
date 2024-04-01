package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func EscapeHTML(s string) string {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
	return replacer.Replace(s)
}

func Format(format string, args ...interface{}) string {
	re := regexp.MustCompile(`\{(\d+)\}`)
	return re.ReplaceAllStringFunc(format, func(s string) string {
		i, _ := strconv.Atoi(s[1 : len(s)-1])
		if i < len(args) {
			return fmt.Sprint(args[i])
		}
		return s
	})
}
