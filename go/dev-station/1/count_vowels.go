package main

import (
	"strings"
)

/*
Task:
=====
Напишите функцию getVowelsCount(str), принимающую строку в качестве аргумента и возвращающую количество гласных, которые содержатся в этой строке.

Гласными являются А, Е, Ё, О, У, Ы, Э, И, Ю, Я.

Функция должна подсчитывать гласные любого регистра, то-есть работать независимо от регистра.
*/

/*
Examples:
========
getVowelsCount('Привет, медвед') // 4
getVowelsCount('Арифметика') // 5
getVowelsCount('Горная порода дикобраза') // 10
*/

func find(target rune, vowels []rune) bool {

	for _, v := range vowels {
		if v == target {
			return true
		}
	}

	return false
}

func countVowels(word string) int {

	vowels := []rune{'а', 'е', 'ё', 'о', 'у', 'ы', 'э', 'и', 'ю', 'я'}

	word = strings.ToLower(word)

	count := 0

	for _, v := range word {

		if res := find(v, vowels); res {
			count++
		}

	}

	return count
}
