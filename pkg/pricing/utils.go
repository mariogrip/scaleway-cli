package pricing

import "math/big"

var (
	intZero = big.NewInt(0)
	intOne  = big.NewInt(1)
	ratZero = big.NewRat(0, 1)
	ratOne  = big.NewRat(1, 1)
)

// Returns a new big.Int set to the ceiling of x.
func ratCeil(x *big.Rat) *big.Int {
	z := new(big.Int)
	m := new(big.Int)
	z.DivMod(x.Num(), x.Denom(), m)
	if m.Cmp(intZero) == 1 {
		z.Add(z, intOne)
	}
	return z
}

// Returns a new big.Rat set to maximum of x and y
func ratMax(x, y *big.Rat) *big.Rat {
	if x.Cmp(y) < 1 {
		return y
	}
	return x
}

// Returns a new big.Rat set to minimum of x and y
func ratMin(x, y *big.Rat) *big.Rat {
	if x.Cmp(y) > 0 {
		return y
	}
	return x
}
