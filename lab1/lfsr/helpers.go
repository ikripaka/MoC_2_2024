package lfsr

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

var ErrUnsupportedSymbol = errors.New("unsupported symbol")

func scaleFromLeft[T any](slices ...[]T) [][]T {
	maxLen := len(lo.MaxBy(slices, func(a []T, b []T) bool {
		return len(a) > len(b)
	}))

	for i := 0; i < len(slices); i++ {
		toAdd := maxLen - len(slices[i])

		sl := make([]T, maxLen)

		for j := 0; j < len(slices[i]); j++ {
			sl[j+toAdd] = slices[i][j]
		}

		slices[i] = sl
	}

	return slices
}

func bitStringToUint8Slice(s string) ([]uint8, error) {
	res := make([]uint8, 0, len(s))

	for _, sym := range strings.Split(s, "") {
		switch sym {
		case "0":
			res = append(res, 0)
		case "1":
			res = append(res, 1)
		default:
			return nil, fmt.Errorf("%w: %v", ErrUnsupportedSymbol, s)
		}
	}

	return res, nil
}
