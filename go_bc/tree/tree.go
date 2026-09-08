package tree

import (
	"go_bc/utils"
	"go_bc/lexer"
)

type BinaryTreeNode struct {
	Token lexer.Token
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
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

func createNode(token lexer.Token) *BinaryTreeNode {
	return &BinaryTreeNode{
		Token: token,
	}
}

func Solve(input []lexer.Token) string {
	root := createNode(input[1])
	root.Left = createNode(input[0])
	root.Right = createNode(input[2])

	return evaluate(root)
}
