# yon
[![Go Reference](https://pkg.go.dev/badge/github.com/MeztliRA/yon.svg)](https://pkg.go.dev/github.com/MeztliRA/yon)

Go module to prompt user with a yes or no question
## example
```
package main

import (
	"fmt"

	"github.com/MeztliRA/yon"
)

func main() {
	answer1 := yon.Prompt("Do you think Go is a good language")
	if answer1 == yon.Yes {
		fmt.Println("ahh")
	} else {
		fmt.Println("umm")
	}
	answer2 := yon.Promptln("Do you think Go is a bad language")
	if answer2 == yon.No {
		fmt.Println("ahh")
	} else {
		fmt.Println("umm")
	}
}
```