package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func getImplementedCommands() []string {
	return []string{"exit", "echo", "type"}
}

func seachExecutable(executable string) (string, error) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		filePath := filepath.Join(path, executable)
		_, err := os.Stat(filePath)
		if err == nil {
			return filePath, nil
		}
	}

	newError := errors.New("file not found")

	return "", newError
}

func handleInputCommands(input string) {
	inputCommand := strings.SplitN(strings.Trim(input, "\n"), " ", 2)
	switch inputCommand[0] {
	case "exit":
		converted, _ := strconv.Atoi(inputCommand[1])
		os.Exit(converted)
	case "echo":
		fmt.Fprintf(os.Stdout, "%s\n", inputCommand[1])
	case "type":
		implementedCommands := getImplementedCommands()

		if slices.Contains(implementedCommands, inputCommand[1]) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", inputCommand[1])
			return
		}

		executablePath, err := seachExecutable(inputCommand[1])

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s not found\n", inputCommand[1])
			return
		}

		fmt.Fprintf(os.Stdout, "%s is %s\n", inputCommand[1], executablePath)
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

		handleInputCommands(input)
	}
}
