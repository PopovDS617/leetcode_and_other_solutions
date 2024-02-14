package ozon

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CheckFloatRound() {
	scanner := bufio.NewScanner(os.Stdin)
	outWriter := bufio.NewWriter(os.Stdout)
	defer outWriter.Flush()

	var n int

	_ = scanner.Scan()

	line := scanner.Text()
	line = strings.Trim(line, "\n\r\t")

	n, _ = strconv.Atoi(line)

	for i := 0; i < n; i++ {
		var goodsAmount, comission int
		_ = scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &goodsAmount, &comission)

		var result float64
		for j := 0; j < goodsAmount; j++ {
			scanner.Scan()
			var a int
			fmt.Sscanf(scanner.Text(), "%d", &a)

			resFloat := float64(a) * float64(comission) / 100
			result += getFractional(resFloat)
		}

		fmt.Fprintf(outWriter, "%.2f\n", result)
	}
}

func getFractional(num float64) float64 {
	return num - math.Floor(num)
}
