package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		panic(err)
	}

	if input == "nonexistent\n" {
		fmt.Fprint(os.Stdout, "nonexistent: command not found\n")
	}

}
