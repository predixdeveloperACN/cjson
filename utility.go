package cjson

import (
	"unicode/utf8"
	"unicode"
	"strings"
	"reflect"
)

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func trimCommaLeft(s string) string {
	i := strings.IndexRune(s, ',')
	if i != -1 {
		return s[:i]
	}
	return s
}

func parseJsonField(v reflect.StructField) (string, bool) {
	var propName string

	if !isFirstRuneUpper(v.Name) {
		return "", false
	}
	jsonTagName := trimCommaLeft(v.Tag.Get("json"))
	if jsonTagName != "" {
		if jsonTagName == "-" {
			return "", false
		}
		propName = jsonTagName
	} else {
		return v.Name, true
	}
	return propName, true
}
