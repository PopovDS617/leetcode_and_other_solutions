package main

import (
	"strconv"
)

/*
Task:
=====
Дано натуральное число number.

Нужно написать функцию sumDigits(number), которая будет считывать все цифры числа number и возвращать их сумму.

Возвращаемое значение должно быть числом.

/*
Examples:
========
sumDigits(1246172836817263875) // 76
sumDigits(123) // 6
sumDigits(5) // 5
sumDigits(0) // 0
*/

func sumDigits(number int) int {

	result := 0

	stringed := strconv.Itoa(number)

	for _, v := range stringed {

		parsed, err := strconv.Atoi(string(v))

		if err != nil {
			parsed = 0
		}
		result += parsed

	}

	return result
}
