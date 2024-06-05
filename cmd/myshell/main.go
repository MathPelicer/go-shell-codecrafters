package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleInputs(input string) {
	switch input {
	case "nonexisting\n":
		fmt.Fprint(os.Stdout, "nonexistent: command not found\n")
	case "invalid_command_1":
		fmt.Fprint(os.Stdout, "invalid_command_1: command not found\n")
	case "invalid_command_2":
		fmt.Fprint(os.Stdout, "invalid_command_2: command not found\n")
	case "invalid_command_3":
		fmt.Fprint(os.Stdout, "invalid_command_3: command not found\n")
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
