package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCalcCoins(t *testing.T) {

	type testcase struct {
		amount int
		result map[string]int
	}

	data := []testcase{
		{amount: 1, result: map[string]int{"1": 1}},
		{amount: 3, result: map[string]int{"1": 1, "2": 1}},
		{amount: 10, result: map[string]int{"10": 1}},
		{amount: 123, result: map[string]int{"1": 1, "2": 1, "10": 12}},
		{amount: 1000, result: map[string]int{"10": 100}},
		{amount: 1006, result: map[string]int{"10": 100, "5": 1, "1": 1}},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(strconv.Itoa(tcase.amount), func(t *testing.T) {
			t.Parallel()

			if res := calcCoins(tcase.amount); !reflect.DeepEqual(tcase.result, res) {
				t.Errorf("Unexpected result. Want %v, got %v", tcase.result, res)
			}

		})

	}

}
