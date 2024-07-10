package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
)

func main() {
	var envs map[string]string = gotenv.GetFromEnv()
	fmt.Println(envs["NAME"])
}
