package apitp

import (
	"fmt"
	"os"
	"strings"
)

var retry bool

func RunCLI() {
quit:
	for {
		if !retry {
			fmt.Println("Welcome to the Merriam-WebsterCLI made by Thorgan!")
		} else {
			retry = false
		}
		fmt.Println()
		fmt.Println("Type a word to search in the dictionary:")
		var word string
		fmt.Scanln(&word)
		if word == "" {
			retry = true
			fmt.Println("You must type a word!")
			continue
		}
		fmt.Println()
		data := apiFetch(word)
		if data == nil {
			fmt.Println("Did you mean any of those words?")
			fmt.Println()
			fmt.Println(strings.Join(propositions, "\n"))
			propositions = append(propositions[0:0])
			retry = true
			continue
		} else if len(data) == 0 {
			fmt.Println(word, ": not found!")
			fmt.Println()
			retry = true
			continue
		} else {
			apiDisplay(data)
		}
		for {
			fmt.Println()
			fmt.Print("Do you want to continue? [y/N] ")
			var answer string
			fmt.Scanln(&answer)
			if strings.ToLower(answer) == "y" {
				fmt.Println()
				fmt.Println()
				retry = true
				break
			} else if strings.ToLower(answer) == "n" || answer == "" {
				break quit
			} else {
				fmt.Println("Invalid answer!")
			}
		}
	}
	fmt.Println()
	fmt.Println("See you later, and thanks for using our program!")
	os.Exit(0)
}
