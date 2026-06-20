package commands

import "fmt"

func Echo(message []string) {
	for _, m := range message {
		fmt.Printf("%s ", m)
	}
	fmt.Println()
}
