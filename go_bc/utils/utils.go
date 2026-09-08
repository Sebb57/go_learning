package utils

import "math/big"

func Add(a, b string) string {
	x, _ := new(big.Int).SetString(a, 10)
	y, _ := new(big.Int).SetString(b, 10)

	result := new(big.Int).Add(x, y)
	return result.String()
}

func Minus(a, b string) string {
	x, _ := new(big.Int).SetString(a, 10)
	y, _ := new(big.Int).SetString(b, 10)

	result := new(big.Int).Sub(x, y)
	return result.String()
}

func Divide(a, b string) string {
	x, _ := new(big.Int).SetString(a, 10)
	y, _ := new(big.Int).SetString(b, 10)

	result := new(big.Int).Quo(x, y)
	return result.String()
}

func Multiply(a, b string) string {
	x, _ := new(big.Int).SetString(a, 10)
	y, _ := new(big.Int).SetString(b, 10)

	result := new(big.Int).Mul(x, y)
	return result.String()
}
