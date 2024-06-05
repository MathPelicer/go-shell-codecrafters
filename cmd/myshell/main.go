package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleInputs(input string) {
	input = strings.Trim(input, "\n")
	switch input {
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
	}
}

func main() {
	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			panic(err)
		}

		handleInputs(input)
	}
}
