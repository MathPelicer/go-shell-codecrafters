package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

type cmd struct {
	Command string
	Args    string
}

func parseInput(input string) *cmd {
	inputCommand := strings.Fields(input)
	args := strings.Join(inputCommand[1:], " ")
	return &cmd{inputCommand[0], args}
}

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

func runExecutable(executable string, param string) (string, error) {
	filePath, err := seachExecutable(executable)

	if err != nil {
		return "", err
	}

	cmd := exec.Command(filePath, param)
	output, err := cmd.Output()

	return string(output), nil
}

func handleInputCommands(input string) {
	inputCommand := parseInput(input)
	switch inputCommand.Command {
	case "exit":
		converted, _ := strconv.Atoi(inputCommand.Args)
		os.Exit(converted)
	case "echo":
		fmt.Fprintf(os.Stdout, "%s\n", inputCommand.Args)
	case "type":
		implementedCommands := getImplementedCommands()

		if slices.Contains(implementedCommands, inputCommand.Args) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", inputCommand.Args)
			return
		}

		executablePath, err := seachExecutable(inputCommand.Args)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s not found\n", inputCommand.Args)
			return
		}

		fmt.Fprintf(os.Stdout, "%s is %s\n", inputCommand.Args, executablePath)
	default:
		output, err := runExecutable(inputCommand.Command, inputCommand.Args)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", inputCommand.Command)

			return
		}
		fmt.Fprintf(os.Stdout, "%s", output)
		return
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
		time.Sleep(300 * time.Millisecond)
	}
}
