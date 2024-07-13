package gotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commentchecker "github.com/OlyMahmudMugdho/gotenv/commentChecker"
	"github.com/OlyMahmudMugdho/gotenv/parser"
)

func GetFromEnv() map[string]string {
	vars := make(map[string]string)

	data, err := os.ReadFile(".env")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		line := scanner.Text()

		if commentchecker.IsComment(line) {
			continue
		}

		var parts []string = parser.ParseLine(line)

		if len(parts) < 2 || len(parts) > 2 {
			continue
		}

		var key string = strings.TrimSpace(parts[0])
		var value string = strings.TrimSpace(parts[1])
		vars[key] = value
	}

	return vars
}
