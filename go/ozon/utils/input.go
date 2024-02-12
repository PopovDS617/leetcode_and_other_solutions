package ozon

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(n int, dataToBind []string) {

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		dataToBind[i] = scanner.Text()
	}
}

func GetIterationsCount() int {
	var n int
	fmt.Scanln(&n)
	return n
}
