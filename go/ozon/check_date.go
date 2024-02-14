package ozon

import (
	ozon "app/ozon/utils"
	"fmt"
	"strconv"
	"strings"
)

func CheckDate() {

	n := ozon.GetIterationsCount()

	inputData := make([]string, n)

	ozon.GetInput(n, inputData)

	for i := range inputData {

		stringSlice := strings.Split(inputData[i], " ")

		intSlice := make([]int, len(stringSlice))

		for i := range stringSlice {
			v, _ := strconv.Atoi(stringSlice[i])
			intSlice[i] = v
		}

		d := intSlice[0]
		m := intSlice[1]
		y := intSlice[2]

		result := isDateCorrect(d, m, y)

		if result {

			fmt.Println(strings.Trim("YES", "\r"))
		} else {
			fmt.Println(strings.Trim("NO", "\r"))
		}

	}
}

func isDateCorrect(d, m, y int) bool {
	res := true

	isLeapYear := y%4 == 0 && y%100 != 0 || y%400 == 0

	if m == 2 {

		if isLeapYear && d > 29 {
			res = false
		}

		if !isLeapYear && d > 28 {
			res = false
		}
	}

	if (m == 4 || m == 6 || m == 9 || m == 11) && d > 30 {
		res = false
	}

	if d <= 0 || d > 31 {
		res = false
	}

	if m <= 0 || m > 12 {
		res = false
	}

	if y < 1950 || y > 2300 {
		res = false
	}

	return res
}
