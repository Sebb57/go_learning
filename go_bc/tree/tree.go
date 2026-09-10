package tree

import (
	"go_bc/lexer"
	"go_bc/utils"
)

type BinaryTreeNode struct {
	Token lexer.Token
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func buildExpressionTree(tokens []lexer.Token) *BinaryTreeNode {
	stack := make([]*BinaryTreeNode, 0)

	for _, token := range tokens {
		node := &BinaryTreeNode{Token: token}
		if token.Type == lexer.Number {
			stack = append(stack, node)
			continue
		}
		node.Right = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, node)
	}
	return stack[0]
}

func evaluate(root *BinaryTreeNode) string {
	switch root.Token.Type {
		case lexer.Number:
			return root.Token.Value

		case lexer.Plus:
			return utils.Add(evaluate(root.Left), evaluate(root.Right))

		case lexer.Minus:
			return utils.Minus(evaluate(root.Left), evaluate(root.Right))

		case lexer.Divide:
			return utils.Divide(evaluate(root.Left), evaluate(root.Right))

		case lexer.Multiply:
			return utils.Multiply(evaluate(root.Left), evaluate(root.Right))
	}

	return "0"
}

func getPrecedence(token lexer.Token) int {
	switch token.Type {
		case lexer.Divide:
			return 2
		case lexer.Multiply:
			return 2
		default:
			return 1
	}
}

func Solve(input []lexer.Token) string {
	if len(input) == 0 {
		return ""
	}

	operators := make([]lexer.Token, 0)
	output := make([]lexer.Token, 0)

	for _, token := range input {
		switch token.Type {
		case lexer.Number:
			output = append(output, token)

		case lexer.LeftParen:
			operators = append(operators, token)

		case lexer.RightParen:
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				if top.Type == lexer.LeftParen {
					break
				}
				output = append(output, top)
			}

		default:
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				if getPrecedence(top) < getPrecedence(token) {
					break
				}
				output = append(output, top)
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}
	for len(operators) > 0 {
		top := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		output = append(output, top)
	}

	var root *BinaryTreeNode = buildExpressionTree(output)
	return evaluate(root)
}
