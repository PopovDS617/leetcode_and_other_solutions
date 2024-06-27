package main

import (
	"strings"
)

/*
Task:
=====
Палиндром — слово, предложение или последовательность символов,
которая абсолютно одинаково читается как в
привычном направлении, так и в обратном.
К примеру, “Anna” — это палиндром, а “table” и “John” — нет.

Дана строка str. Нужно написать функцию isPalindrome(str), которая позволяет вернуть значение true,
если строка является палиндромом, и false — если нет.

Алгоритм не должен быть зависимым от регистра символов.
*/

/*
Examples:
=========
isPalindrome('Мадам') // true
isPalindrome('Шалаш') // true
isPalindrome('Стол') // false
isPalindrome('Кореной') // false
*/

func isPalindrome(str string) bool {

	runed := []rune(str)

	for i, j := 0, len(runed)-1; i < len(runed); i, j = i+1, j-1 {
		if !strings.EqualFold(string(runed[i]), string(runed[j])) {
			return false
		}
	}
	return true
}
