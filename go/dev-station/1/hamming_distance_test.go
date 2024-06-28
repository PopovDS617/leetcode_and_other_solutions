package main

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {

	type testcase struct {
		str1   string
		str2   string
		result int
	}

	data := []testcase{
		{str1: "Hello", str2: "Hallo", result: 1},
		{str1: "Расстояние", str2: "Хэмминга", result: -1},
		{str1: "1010", str2: "1111", result: 1},
		{str1: "Bobu", str2: "Bobr", result: 3},
		{str1: "Clip", str2: "Clip", result: 4},
		{str1: "Slip", str2: "Clip", result: 0},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(tcase.str1+" "+tcase.str2, func(t *testing.T) {
			t.Parallel()

			if res := hammingDistance(tcase.str1, tcase.str2); res != tcase.result {
				t.Errorf("Unexpected result. Want %d, got %d", tcase.result, res)
			}

		})

	}

}
