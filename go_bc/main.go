package main

import (
	"fmt"
	"bufio"
	"os"
	"go_bc/lexer"
	"go_bc/tree"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lexedLine, err := lexer.LexInput(line)

		if  err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(tree.Solve(lexedLine))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
