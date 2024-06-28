package main

import (
	"fmt"
	"testing"
)

func TestCalculateBMI(t *testing.T) {

	type testcase struct {
		weight int
		height int
		result string
	}

	data := []testcase{
		{weight: 70, height: 170, result: "24.22 нормальный вес"},
		{weight: 70, height: 200, result: "17.50 недостаток веса"},
		{weight: 90, height: 180, result: "27.78 избыточный вес"},
		{weight: 100, height: 180, result: "30.86 ожирение"},
		{weight: 220, height: 170, result: "76.12 ожирение"},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(fmt.Sprintf("w: %d, h: %d", tcase.weight, tcase.height), func(t *testing.T) {
			t.Parallel()

			if res := calculateBMI(tcase.weight, tcase.height); res != tcase.result {
				t.Errorf("Unexpected result. Want %s, got %s", tcase.result, res)
			}

		})

	}

}
