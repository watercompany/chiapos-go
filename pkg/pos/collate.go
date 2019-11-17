package pos

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/kargakis/gochia/pkg/utils"
)

// CollaSize returns the collation size for t.
func CollaSize(t int) (size *int, err error) {
	size = new(int)
	switch t {
	case 2:
		*size = 1
	case 3, 7:
		*size = 2
	case 4, 5:
		*size = 4
	case 6:
		*size = 3
	default:
		return nil, fmt.Errorf("collation size for t=%d is undefined", t)
	}
	return
}

// Ct is a collation function for t.
func Ct(t int, k uint64, x ...uint64) (*big.Int, error) {
	if t < 2 || t > 7 {
		return nil, fmt.Errorf("collation function for t=%d is undefined", t)
	}
	twoToTMinusTwo := int(math.Pow(2, float64(t-2)))
	if len(x) != twoToTMinusTwo {
		return nil, fmt.Errorf("invalid x count: %d, expected %d", len(x), twoToTMinusTwo)
	}

	switch t {
	case 2:
		return new(big.Int).SetUint64(x[0]), nil

	case 3:
		return utils.Concat(k, x[0], x[1]), nil

	case 4:
		return utils.Concat(k, x[0], x[1], x[2], x[3]), nil

	case 5:
		left := utils.Concat(k, x[0], x[1], x[2], x[3])
		right := utils.Concat(k, x[4], x[5], x[6], x[7])
		return left.Xor(left, right), nil

	case 6:
		first := utils.Concat(k, x[0], x[1], x[2])
		second := utils.Concat(k, x[4], x[5], x[6])
		third := utils.Concat(k, x[8], x[9], x[10])
		fourth := utils.Concat(k, x[12], x[13], x[14])
		return first.Xor(first, second).Xor(first, third).Xor(first, fourth), nil

	case 7:
		first := utils.Concat(k, x[0], x[1])
		second := utils.Concat(k, x[4], x[5])
		third := utils.Concat(k, x[8], x[9])
		fourth := utils.Concat(k, x[12], x[13])
		fifth := utils.Concat(k, x[16], x[17])
		sixth := utils.Concat(k, x[20], x[21])
		seventh := utils.Concat(k, x[24], x[25])
		eighth := utils.Concat(k, x[28], x[29])
		return first.Xor(first, second).
			Xor(first, third).
			Xor(first, fourth).
			Xor(first, fifth).
			Xor(first, sixth).
			Xor(first, seventh).
			Xor(first, eighth), nil
	}
	return nil, errors.New("should never reach here")
}
