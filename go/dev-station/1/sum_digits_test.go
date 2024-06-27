package main

import (
	"fmt"
	"testing"
)

func TestSumDigits(t *testing.T) {

	type testcase struct {
		original int
		result   int
	}

	data := []testcase{
		{original: 123, result: 6},
		{original: 222, result: 6},
		{original: 1111, result: 4},
		{original: 36817263875, result: 56},
		{original: 12461728, result: 31},
		{original: 11223344, result: 20},
		{original: 1246172836817263875, result: 87},
		{original: 0, result: 0},
		{original: -1, result: 1},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(fmt.Sprintf("%d", tcase.original), func(t *testing.T) {
			t.Parallel()

			if res := sumDigits(tcase.original); res != tcase.result {
				t.Errorf("Unexpected result. Want %d, got %d", tcase.result, res)
			}

		})

	}

}
