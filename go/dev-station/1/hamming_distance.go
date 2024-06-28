package main

/*
Task:
=====
Расстояние Хэмминга — число позиций, в которых соответствующие символы двух слов одинаковой длины различны.

Напишите функцию hammingDistance(str1, str2), принимающую на вход две строки одинаковой длины.
Функция должна возвращать число - расстояние Хэмминга.

Если строки имеют разную длину, то функция должна возвращать: -1.
/*
Examples:
========
hammingDistance('Hello', 'Hallo') // 1
hammingDistance('Расстояние', 'Хэмминга') // -1
hammingDistance('1010', '1111') // 2
*/

func hammingDistance(first string, second string) int {

	firstRuned := []rune(first)
	secondRuned := []rune(second)

	if len(firstRuned) != len(secondRuned) {
		return -1
	}

	var count int

	for i := 0; i < len(firstRuned); i++ {
		if firstRuned[i] == secondRuned[i] {
			count++
		} else {
			break
		}
	}

	return count

}
