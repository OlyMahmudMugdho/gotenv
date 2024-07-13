package main

import (
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
)

func main() {
	//var envs map[string]string = gotenv.GetFromEnv()

	/* for _, v := range envs {
		fmt.Println(v)
	} */

	gotenv.Load()
	fmt.Println(os.Getenv("AGE"))

}
