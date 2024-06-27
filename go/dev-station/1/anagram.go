package main

import (
	"strings"
)

/*
Task:
=====
Анаграмма - это слово, которое содержит все буквы другого слова в том же количестве, но ином порядке.

Нужно написать функцию isAnagram(a, b), которая проверяет, являются ли две строки анаграммами.

Регистр букв не имеет значения.

Пробелы или знаки препинания в расчет не берутся.
*/

/*
Examples:
========
isAnagram('Гора', 'РоГа') // true
isAnagram('Клоун', 'кулон') // true
isAnagram('Восемь', 'Семя') // false
isAnagram('дерево', 'самолет') // false
*/

func isAnagram(first, second string) bool {

	firstRuned, secondRuned := []rune(strings.ToLower(first)), []rune(strings.ToLower(second))

	firstLength, secondLength := len(firstRuned), len(secondRuned)

	if firstLength != secondLength {
		return false
	}

	countFirst := make(map[rune]int, firstLength)
	countSecond := make(map[rune]int, secondLength)

	for i := 0; i < firstLength; i++ {

		countFirst[firstRuned[i]]++
		countSecond[rune(secondRuned[i])]++

	}

	for k := range countFirst {

		if countFirst[k] != countSecond[k] {
			return false
		}

	}

	return true
}
