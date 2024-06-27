package main

import (
	"testing"
)

func TestCountVowels(t *testing.T) {

	type testcase struct {
		word   string
		result int
	}

	data := []testcase{
		{word: "Гора", result: 2},
		{word: "ААА", result: 3},
		{word: "Ббб", result: 0},
		{word: "Горная порода дикобраза", result: 10},
		{word: "Меч", result: 1},
		{word: "Мечта", result: 2},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(tcase.word, func(t *testing.T) {
			t.Parallel()

			if res := countVowels(tcase.word); res != tcase.result {
				t.Errorf("Unexpected result. Want %v, got %v", tcase.result, res)
			}

		})

	}

}
