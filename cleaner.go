package MOMO

import (
	"strings"
)

var replacer *strings.Replacer /* initialized in init.go */

func sanitizeString(str string) string {
	return replacer.Replace(str)
}
