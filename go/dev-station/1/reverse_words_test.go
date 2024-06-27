package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {

	type testcase struct {
		original string
		result   string
	}

	data := []testcase{
		{original: "Гречка может быть вкусной!", result: "акчерГ тежом ьтыб !йонсукв"},
		{original: "Апельсины - профилактика от авитоминоза", result: "ынисьлепА - акиткалифорп то азонимотива"},
		{original: "Не всё должно быть наделено смыслом", result: "еН ёсв онжлод ьтыб онеледан молсымс"},
		{original: "Абба", result: "аббА"},
	}

	for _, tcase := range data {
		tcase := tcase

		t.Run(tcase.original, func(t *testing.T) {
			t.Parallel()

			if res := reverseWords(tcase.original); res != tcase.result {
				t.Errorf("Unexpected result. Want %s, got %s", tcase.result, res)
			}

		})

	}

}
