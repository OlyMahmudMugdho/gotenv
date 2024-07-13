package gotenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	commentchecker "github.com/OlyMahmudMugdho/gotenv/commentChecker"
	"github.com/OlyMahmudMugdho/gotenv/parser"
)

func GetFromEnv(location ...string) (map[string]string, error) {

	vars := make(map[string]string)
	var data []byte

	if len(location) == 0 {
		d, err := os.ReadFile(".env")

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		data = d
	} else {
		d, err := os.ReadFile(location[0])

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		data = d
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

	return vars, nil
}

func SetAllVariables(vars map[string]string) {
	for k, v := range vars {
		var err error = os.Setenv(k, v)
		if err != nil {
			log.Fatal(err)
			fmt.Printf("unable set variable %v=%v", k, v)
		}
		continue
	}
}

func Load(locations ...string) error {
	vars, err := GetFromEnv(locations...)
	if err != nil {
		return err
	}
	SetAllVariables(vars)
	return nil
}
