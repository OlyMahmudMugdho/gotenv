package parser

import (
	"strings"
)

func ParseLine(line string) []string {
	parts := strings.SplitN(line, "=", 2)
	return parts
}
