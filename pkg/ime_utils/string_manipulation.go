package ime_utils

import "strings"

func Split(sentence string, sep rune) []string {
	splitFn := func(c rune) bool {
		return c == sep
	}
	return strings.FieldsFunc(sentence, splitFn)
}
