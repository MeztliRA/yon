// yon is a go module to prompt user with a yes or no question
package yon

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// type for user response
type Response bool

// constant for yes or no response from user
const (
	Yes = Response(true)
	No  = Response(false)
)

// prompt the user with a yes or no question using the passed argument as the prompt, return yon.Yes or yon.No
func Prompt(prompt string) Response {
	return promptCore(prompt+"(yes/no) ", false)
}

// similar to Prompt(), but a new line is appended to the prompt, so the user is inputting their response on a new line, return yon.Yes or yon.No
func Promptln(prompt string) Response {
	return promptCore(prompt+"(yes/no)", true)
}

// Similar to Prompt(), but only prompt the user for a responce once, return yon.Yes or yon.No and a error, error is not nil if user didnt response with yes or no
func PromptOnce(prompt string) (Response, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + "(yes/no) ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answer = strings.Trim(answer, "\n")

	switch answer {
	case "y", "Y", "yes", "Yes", "YES":
		return Yes, nil
	case "n", "N", "no", "No", "NO":
		return No, nil
	default:
		return No, errors.New("prompt: user didnt input yes or no")
	}
}

func promptCore(prompt string, ln bool) Response {
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
