package ozon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckACTemp() {
	scanner := bufio.NewScanner(os.Stdin)
	outWriter := bufio.NewWriter(os.Stdout)
	defer outWriter.Flush()

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	min := 15
	max := 30
	for i := 0; i < n; i++ {
		scanner.Scan()
		nw, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		for j := 0; j < nw; j++ {
			scanner.Scan()
			numbersStr := strings.Split(scanner.Text(), " ")
			s := numbersStr[0]
			temp, ok := strconv.Atoi(numbersStr[1])
			if ok != nil {
				return
			}

			if s == "<=" && temp < max {
				max = temp
			} else if s == ">=" && temp > min {
				min = temp
			}
			if min <= max {
				fmt.Fprintf(outWriter, "%d\n", min)
			} else {
				fmt.Fprintf(outWriter, "%s\n", "-1")
			}
		}
		fmt.Fprintf(outWriter, "\n")
		min = 15
		max = 30
	}
}
