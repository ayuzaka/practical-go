package chapter03

import "math/big"

type ImmutableMoney struct {
	currency int
	amount   *big.Int
}

func (im ImmutableMoney) Currency() int {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c int) ImmutableMoney {
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}
