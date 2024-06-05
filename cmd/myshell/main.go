package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleInputs(input string) {
	inputCommand := strings.SplitN(strings.Trim(input, "\n"), " ", 2)
	switch inputCommand[0] {
	case "exit 0":
		os.Exit(0)
	case "echo":
		fmt.Fprintf(os.Stdout, "%s", inputCommand[1])
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
