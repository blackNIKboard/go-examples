package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type name struct {
	num int
	deg int
}

func main() {
	test := []name{
		{64, 2},
		{25, 2},
		{125, 3},
		{216, 3},
	}

	for _, n := range test {
		res, err := NewRoot(new(big.Int).SetInt64(int64(n.num)), n.deg)
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	}
}

func NewRoot(num *big.Int, deg int) (*big.Int, error) {
	res := &big.Int{}

	res, err := SumOfDigits(num)
	if err != nil {
		return res, err
	}

	res.Sub(res, new(big.Int).SetInt64(int64(deg)))

	return res, nil
}

func SumOfDigits(num *big.Int) (*big.Int, error) {
	sum := &big.Int{}

	text, err := num.MarshalText()
	if err != nil {
		return sum, err
	}

	for _, b := range text {
		inInt, err := strconv.Atoi(string(b))
		if err != nil {
			return &big.Int{}, err
		}

		sum.Add(sum, new(big.Int).SetInt64(int64(inInt)))
	}

	return sum, nil
}
