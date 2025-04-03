package funcs

import (
	"math/big"
)

func MultiplyString(num1 string, num2 string) string {
	if len(num1) >= 200 || len(num2) >= 200 {
		return ""
	}

	bigInt1, ok1 := new(big.Int).SetString(num1, 10)
	bigInt2, ok2 := new(big.Int).SetString(num2, 10)

	if !ok1 || !ok2 {
		return ""
	}

	bigIntMul := new(big.Int)
	mul := bigIntMul.Mul(bigInt1, bigInt2)

	return mul.String()
}
