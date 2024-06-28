package main

/*
Task:
=====
Двумерный массив - массив, элементами которого являются другие массивы.
Он формирует таблицу с рядами и столбцами, где каждый элемент имеет два индекса: для строки и столбца.

Напишите функцию isMatrixHasUniqueValues(square), которая принимает двумерный массив
и проверяет уникальность каждого элемента массива - числа.

Функция возвращает true, если каждое число является уникальным, и false - в противном случае.

/*
Examples:
========
isMatrixHasUniqueValues([
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
]) // true
isMatrixHasUniqueValues([
  [1, 2, 3],
  [1, 2, 3],
  [1, 2, 3]
]) // false
isMatrixHasUniqueValues([
  [9, 2, 4],
  [3, 5, 8],
  [1, 6, 7]
]) // true
isMatrixHasUniqueValues([
  [0, 0, 0],
  [0, 0, 0],
  [0, 0, 0]
]) // false
*/

func isMatrixHasUniqueValues(matrix [][]int) bool {

	count := make(map[int]struct{}, 9)

	for _, i := range matrix {
		for _, j := range i {
			count[j] = struct{}{}
		}
	}

	return len(count) == 9
}
