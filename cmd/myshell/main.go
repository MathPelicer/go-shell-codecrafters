package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getImplementedCommands() []string {
	return []string{"exit", "echo", "type"}
}

func handleInputs(input string) {
	inputCommand := strings.SplitN(strings.Trim(input, "\n"), " ", 2)
	switch inputCommand[0] {
	case "exit":
		converted, _ := strconv.Atoi(inputCommand[1])
		os.Exit(converted)
		return
	case "echo":
		fmt.Fprintf(os.Stdout, "%s\n", inputCommand[1])
		return
	case "type":
		implementedCommands := getImplementedCommands()
		if slices.Contains(implementedCommands, inputCommand[1]) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", inputCommand[1])
			return
		}
		fmt.Fprintf(os.Stdout, "%s not found\n", inputCommand[1])

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
