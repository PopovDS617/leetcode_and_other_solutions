package main

import (
	"app/utils"
	"fmt"
	"regexp"
	"strings"
)

func CheckPlayes() {

	n := utils.GetIterationsCount()

	inputData := make([]string, n)

	utils.GetInput(n, inputData)

	for _, v := range inputData {

		rule := regexp.MustCompile(`([a-zA-Z]\d\d[a-zA-Z][a-zA-Z])|([a-zA-Z]\d[a-zA-Z][a-zA-Z])`)

		res := rule.FindAllString(v, 999)

		if pass := rule.ReplaceAllString(v, ""); len(pass) == 0 && len(res) != 0 {
			fmt.Println(strings.Join(res, " "))
		} else {
			fmt.Println("-")
		}

	}

}
