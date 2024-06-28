package main

import "fmt"

/*
Task:
=====
Напишите функцию calcCoins(amount), принимающую на вход сумму денег amount.

Функция должна возвращать минимальное количество монет разных номиналов для набора определенной суммы денег.

Имеются монеты номиналом: 1, 2, 5, 10 рублей.
/*
Examples:
========
calcCoins(123) // {'1': 1, '2': 1, '10': 12}
calcCoins(184) // {'2': 2, '10': 18}
calcCoins(111) // {'1': 1, '10': 11}
*/

func calcCoins(amount int) map[string]int {

	result := make(map[string]int)

	for amount != 0 {

		switch {
		case amount >= 10:
			amount -= 10
			result["10"]++
		case amount >= 5:
			amount -= 5
			result["5"]++
		case amount >= 2:
			amount -= 2
			result["2"]++
		case amount >= 1:
			amount -= 1
			result["1"]++
		}

	}

	return result

}

func main() {
	fmt.Println(calcCoins(123))
}
