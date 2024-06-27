package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {

	type testcase struct {
		first, second string
		result        bool
	}

	data := []testcase{
		{first: "Гора", second: "ГоРа", result: true},
		{first: "Клоун", second: "КулОн", result: true},
		{first: "Восемь", second: "Семь", result: false},
		{first: "Роза", second: "Озар", result: true},
		{first: "Горб", second: "Брог", result: true},
		{first: "ITT", second: "TTI", result: true},
		{first: "WORK", second: "KROW", result: true},
		{first: "WORK", second: "CROW", result: false},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(fmt.Sprintf("%s - %s", tcase.first, tcase.second), func(t *testing.T) {
			t.Parallel()

			if res := isAnagram(tcase.first, tcase.second); res != tcase.result {
				t.Errorf("Unexpected result. Want %v, got %v", tcase.result, res)
			}

		})

	}

}
