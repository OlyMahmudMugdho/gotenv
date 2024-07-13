package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
)

func main() {
	var envs map[string]string = gotenv.GetFromEnv()

	for _, v := range envs {
		fmt.Println(v)
	}
}
