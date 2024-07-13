package main

import (
	"fmt"
	"os"

	"github.com/OlyMahmudMugdho/gotenv/gotenv"
)

func main() {
	gotenv.Load(".another.env")
	fmt.Println(os.Getenv("AGE"))
	fmt.Println(os.Getenv("SOME"))

}
