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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt + "(yes/no) ")
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

func Promptln(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(prompt + "(yes/no)")
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
