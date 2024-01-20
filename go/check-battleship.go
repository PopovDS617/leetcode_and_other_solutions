package main

import (
	"app/utils"
	"fmt"
	"strconv"
	"strings"
)

func CheckBattleship() {

	n := utils.GetIterationsCount()

	inputData := make([]string, n)

	utils.GetInput(n, inputData)

	for i := range inputData {

		stringSlice := strings.Split(inputData[i], " ")

		intSlice := make([]int, len(stringSlice))

		for i := range stringSlice {
			v, _ := strconv.Atoi(stringSlice[i])
			intSlice[i] = v
		}

		getResult(intSlice)

	}

}

func getResult(dataSlice []int) {
	result := make(map[string]int, 3)

	for _, v := range dataSlice {

		switch v {
		case 1:
			result["1"] += 1

		case 2:
			result["2"] += 1

		case 3:
			result["3"] += 1

		case 4:
			result["4"] += 1

		}

	}

	if result["1"] == 4 && result["2"] == 3 && result["3"] == 2 && result["4"] == 1 {

		fmt.Println(strings.Trim("YES", "\r"))
	} else {
		fmt.Println(strings.Trim("NO", "\r"))
	}
}
