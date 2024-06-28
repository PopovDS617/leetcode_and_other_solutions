package main

import (
	"fmt"
	"strings"
)

/*
Task:
=====
ИМТ (индекс массы тела) — это показатель оценки физического развития, который рассчитывается по формуле:
масса тела в килограммах делится на квадрат роста в метрах.

Напишите функцию calculateBMI(weight, height), которая принимает вес weight и рост height (в сантиметрах).

Функция должна возвращать следующие категории:
- если ИМТ меньше 18.5 (не включительно) - недостаток веса
- если ИМТ находится в диапазоне от 18.5 (включительно) до 25 (не включительно) - нормальный вес
- если ИМТ находится в диапазоне от 25 (включительно) до 30 невключительно - избыточный вес
- если ИМТ находится от 30 и выше - ожирение

ИМТ должен быть округлен до двух знаков после запятой.

Формула для вычисления ИМТ: масса тела / вес (в метрах)^2

Функция должна возвращать массив. Сначала ИМТ, потом категорию.

/*
Examples:
========
calculateBMI(70, 170) // [ 24.22, 'нормальный вес' ]
calculateBMI(30, 150) // [ 13.33, 'недостаток веса' ]
calculateBMI(90, 180) // [ 27.78, 'избыточный вес' ]
calculateBMI(100, 180) // [ 30.86, 'ожирение' ]
*/

func calculateBMI(weight int, height int) string {

	var result []string
	var message string

	var heightToMeters float64 = float64(height) / 100

	bmiFloat := float64(weight) / (heightToMeters * heightToMeters)

	bmiStringify := fmt.Sprintf("%.2f", bmiFloat)

	result = append(result, bmiStringify)

	switch {
	case bmiFloat < 18.5:
		message = "недостаток веса"
	case bmiFloat >= 18.5 && bmiFloat < 25:
		message = "нормальный вес"
	case bmiFloat >= 25 && bmiFloat < 30:
		message = "избыточный вес"
	case bmiFloat >= 30:
		message = "ожирение"
	}

	result = append(result, message)

	return strings.Join(result, " ")

}
