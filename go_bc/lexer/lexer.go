package lexer

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	Number TokenType = iota
	Plus
	Minus
	Multiply
	Divide
	LeftParen
	RightParen
	EOF
	Invalid
)

var TokenValues = map[rune]TokenType{
	'+': Plus,
	'-': Minus,
	'*': Multiply,
	'/': Divide,
	'(': LeftParen,
	')': RightParen,
	';': EOF,
}

type Token struct {
	Type  TokenType
	Value string
}

func LexInput(input string) ([]Token, error) {
	output := []Token{}
	parenthesisCount := 0
	prevValNumber := false

	for i := 0; i < len(input); {
		element := rune(input[i])

		if unicode.IsSpace(element) {
			i++
			continue
		}

		if tokenType, ok := TokenValues[element]; ok {

			if tokenType == LeftParen {
				parenthesisCount++
			}
			if tokenType == RightParen {
				parenthesisCount--
				if parenthesisCount < 0 {
					return nil, fmt.Errorf("unexpected ')' at position %d", i)
				}
			}

			if tokenType == EOF {
				if parenthesisCount != 0 {
					return nil, fmt.Errorf("unclosed parenthesis")
				}
			}

			if prevValNumber == false && tokenType != LeftParen {
				return nil, fmt.Errorf("unexpected %q symbol at position %d", element, i)
			}
			prevValNumber = false
			output = append(output, Token{Type: tokenType, Value: string(element)})
			i++
			continue
		}

		if unicode.IsNumber(element) {
			prevValNumber = true
			start := i
			for i < len(input) && unicode.IsNumber(rune(input[i])) {
				i++
			}
			output = append(output, Token{Type: Number, Value: input[start:i]})
			continue
		} else {
			prevValNumber = false
		}

		return nil, fmt.Errorf("invalid character %q at position %d", element, i)
	}

	if parenthesisCount != 0 {
		return nil, fmt.Errorf("unclosed parenthesis")
	}
	return output, nil
}
