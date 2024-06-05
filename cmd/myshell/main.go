package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleInputs(input string) {
	inputCommand := strings.SplitN(strings.Trim(input, "\n"), " ", 2)
	switch inputCommand[0] {
	case "exit":
		converted, _ := strconv.Atoi(inputCommand[1])
		os.Exit(converted)
	case "echo":
		fmt.Fprintf(os.Stdout, "%s\n", inputCommand[1])
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", inputCommand[0])
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
