package main

import (
	"fmt"
	"testing"
)

func TestIsMatrixHasUniqueValues(t *testing.T) {

	type testcase struct {
		matrix [][]int
		result bool
	}

	data := []testcase{
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, result: true},
		{matrix: [][]int{
			{0, 2, 3},
			{4, 0, 6},
			{7, 8, 0},
		}, result: false},
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 9},
		}, result: false},
		{matrix: [][]int{
			{},
			{},
			{},
		}, result: false},
	}

	for i, tcase := range data {
		tcase := tcase

		t.Run(fmt.Sprintf("case - %d", i), func(t *testing.T) {
			t.Parallel()

			if res := isMatrixHasUniqueValues(tcase.matrix); res != tcase.result {
				t.Errorf("Unexpected result. Want %v, got %v", tcase.result, res)
			}

		})

	}

}
