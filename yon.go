// yon is a go module to prompt user with a yes or no question
package yon

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// type for user response
type Response string

// constant for yes or no response from user
const (
	Yes = Response("YES")
	No  = Response("NO")
)

// prompt the user with a yes or no question using the passed argument as the prompt, return yon.Yes or yon.No
func Prompt(prompt string) Response {
	answer := promptCore(prompt+"(yes/no) ", false)
	return answer
}

// similar to Prompt(), but a new line is appended to the prompt, so the user is inputting their response on a new line, return yon.Yes or yon.No
func Promptln(prompt string) Response {
	answer := promptCore(prompt+"(yes/no)", true)
	return answer
}

func promptCore(prompt string, ln bool) Response {
	log.SetPrefix("prompt: ")
	log.SetFlags(0)

	reader := bufio.NewReader(os.Stdin)

	for {
		if ln {
			fmt.Println(prompt)
		} else {
			fmt.Print(prompt)
		}
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answer = strings.Trim(answer, "\n")

		switch answer {
		case "y", "Y", "yes", "Yes", "YES":
			return Yes
		case "n", "N", "no", "No", "NO":
			return No
		default:
			continue
		}
	}
}
