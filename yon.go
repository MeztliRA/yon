package yon

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Yes = "YES"
	No  = "NO"
)

func Prompt(prompt string) string {
	answer := promptCore(prompt+"(yes/no) ", false)
	return answer
}

func Promptln(prompt string) string {
	answer := promptCore(prompt+"(yes/no)", true)
	return answer
}

func promptCore(prompt string, ln bool) string {
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
