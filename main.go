package main

import (
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
)

func main() {
	gotenv.Load(".another.env")
	gotenv.Load()
	fmt.Println(os.Getenv("AGE"))
	fmt.Println(os.Getenv("SOME"))

}
