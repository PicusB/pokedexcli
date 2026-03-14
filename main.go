package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		cleanWords := cleanInput(words)
		if len(cleanWords) == 0 {
			fmt.Println("No Command Received")
		} else {
			fmt.Printf("Your command was: %s\n", cleanWords[0])
		}

	}
}
