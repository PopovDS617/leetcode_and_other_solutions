package main

import (
	"strings"
)

/*
Task:
=====
Напишите функцию reverseWords(str), которая принимает на вход предложение и возвращает его же,
но уже с развернутыми словами, разделенными пробелами.

/*
Examples:
========
reverseWords('Гречка может быть вкусной!')
// ”акчерГ тежом ьтыб !йонсукв”

reverseWords('Апельсины - профилактика от онкологии')
// 'ынисьлепА - акиткалифорп то ииголокно'

reverseWords('Не всё должно быть наделено смыслом')
// 'еН ёсв онжлод ьтыб онеледан молсымс'
*/

func reverse(word string) string {

	runed := []rune(word)

	for i, j := 0, len(runed)-1; i < j; i, j = i+1, j-1 {
		runed[i], runed[j] = runed[j], runed[i]
	}

	return string(runed)
}

func reverseWords(text string) string {

	reversedSlice := strings.Split(text, " ")

	for i := 0; i < len(reversedSlice); i++ {
		reversedSlice[i] = reverse(reversedSlice[i])
	}

	return strings.Join(reversedSlice, " ")
}
