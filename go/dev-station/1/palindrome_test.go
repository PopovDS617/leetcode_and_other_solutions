package main

import (
	"testing"
)

type testcase struct {
	word   string
	result bool
}

func TestIsPalindrome(t *testing.T) {

	data := []testcase{
		{word: "Мадам", result: true},
		{word: "Шалаш", result: true},
		{word: "Стол", result: false},
		{word: "Кореной", result: false},
		{word: "Rover", result: false},
		{word: "Bob", result: true},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(tcase.word, func(t *testing.T) {
			t.Parallel()

			if res := isPalindrome(tcase.word); res != tcase.result {
				t.Errorf("Unexpected result. Want %v, got %v", tcase.result, res)
			}

		})

	}

}
