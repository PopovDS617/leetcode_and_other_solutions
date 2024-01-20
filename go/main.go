package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numTests := GetIterationsCount()

	res := [][]string{}

	var resString string

	for i := 0; i < numTests; i++ {
		testRes := processTest()

		res = append(res, testRes)
	}

	// f, err := os.Create("data.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// _, err2 := f.WriteString("old falcon\n")

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// fmt.Println("done")

	for i := range res {
		for j := range res[i] {
			// fmt.Print(res[i][j])
			// fmt.Print("\n")

			resString += res[i][j]
			resString += "\n"
			// f.WriteString(res[i][j] + "\n")
		}
		// f.WriteString("\n")
		// fmt.Print("\n\r")
		// fmt.Println()
		resString += "\n"
	}

	resString += "\r"
	fmt.Println(resString)

}

func processTest() []string {

	numConstraints := GetIterationsCount()

	constraints := make([]string, numConstraints)

	scanner := bufio.NewReader(os.Stdin)
	for i := 0; i < numConstraints; i++ {

		line, _ := scanner.ReadString('\n')

		constraints[i] = line
	}

	result := satisfyConstraints(constraints)

	return result
}

func satisfyConstraints(constraints []string) []string {

	min := 15
	max := 30

	result := []string{}

	for _, constraint := range constraints {

		parts := strings.Fields(constraint)
		operator := parts[0]
		value, _ := strconv.Atoi(parts[1])

		switch operator {
		case ">=":
			if value < 15 || value > 30 || value > max {
				min = -1
				break
			}
			if value > min {
				min = value
			}

		case "<=":
			if value > 30 || value < 15 || value < min {
				min = -1
				break
			}
			if value < max {
				max = value
			}

		}

		result = append(result, strconv.Itoa(min))

	}

	return result
}

func GetIterationsCount() int {
	var n int
	fmt.Scanln(&n)
	return n
}
