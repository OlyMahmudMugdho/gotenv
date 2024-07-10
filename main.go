package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OlyMahmudMugdho/gotenv/parser"
)

func main() {
	vars := make(map[string]string)

	data, err := os.ReadFile(".env")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		line := scanner.Text()
		var parts []string = parser.ParseLine(line)
		var key string = strings.TrimSpace(parts[0])
		var value string = strings.TrimSpace(parts[1])
		vars[key] = value
	}
}
