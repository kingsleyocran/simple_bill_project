package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Printf(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
