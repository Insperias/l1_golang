package main

import (
	"fmt"
	"math/big"
)

func Div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func Add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func main() {
	a, _ := new(big.Int).SetString("1000", 10)
	b, _ := new(big.Int).SetString("1000", 10)
	c := Sub(a, b)
	fmt.Println(c)
}
