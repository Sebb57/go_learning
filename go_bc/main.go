package main

import (
	"bufio"
	"fmt"
	"go_bc/lexer"
	"os"
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
		for _, token := range lexedLine {
			fmt.Println("type:", token.Type, "value:", token.Value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
