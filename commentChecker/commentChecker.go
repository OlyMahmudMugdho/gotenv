package commentchecker

import "strings"

func IsComment(line string) bool {
	return strings.HasPrefix(line, "#")
}
