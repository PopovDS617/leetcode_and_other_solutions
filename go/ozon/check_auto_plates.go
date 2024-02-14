package ozon

import (
	ozon "app/ozon/utils"
	"fmt"
	"regexp"
	"strings"
)

func CheckPlayes() {

	n := ozon.GetIterationsCount()

	inputData := make([]string, n)

	ozon.GetInput(n, inputData)

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
