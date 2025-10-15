package smallhelpers

import (
	"strings"
)

func Normalize(unnormal string) string {
	return strings.ToLower(strings.TrimSpace(unnormal))
}