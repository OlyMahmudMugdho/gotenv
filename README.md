
# gotenv

`gotenv` is a Go package designed to load environment variables from a `.env` file into your application.

## Installation

To install the package, use the following command:

```
go get github.com/OlyMahmudMugdho/gotenv
```
## Usage
1. Create a `.env` file in the root of your project with the following content:
	```
	KEY=value
	ANOTHER_KEY=another_value
    ```
2.	In your Go application, import the package and load the environment variables:
	```
	package main
	import (
	    "fmt"
	    "github.com/OlyMahmudMugdho/gotenv"
	    "os"
	)

	func main() {
	    gotenv.Load()

	    value := os.Getenv("KEY")
	    fmt.Println("KEY:", value)
	}
	```
3. Run your application:
		`go run main.go`  
		
	You should see the output:
	`KEY: value
`
