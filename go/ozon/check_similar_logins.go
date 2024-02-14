package ozon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckSimilarLogins() {
	scanner := bufio.NewScanner(os.Stdin)
	outWriter := bufio.NewWriter(os.Stdout)
	defer outWriter.Flush()

	var n int

	_ = scanner.Scan()

	line := scanner.Text()
	line = strings.Trim(line, "\n\r\t")

	n, _ = strconv.Atoi(line)

	logins := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&logins[i])
	}

	var newEmplCount int

	_ = scanner.Scan()

	line = scanner.Text()
	line = strings.Trim(line, "\n\r\t")

	newEmplCount, _ = strconv.Atoi(line)

	for j := 0; j < newEmplCount; j++ {

		_ = scanner.Scan()

		newLogin := scanner.Text()
		newLogin = strings.Trim(newLogin, "\n\r\t")

		found := false
		for _, login := range logins {
			if login == newLogin {
				found = true
				break
			}

			if areLoginsSimilar(login, newLogin) {
				found = true
				break
			}
		}

		_ = found

		if found {
			fmt.Fprintln(outWriter, 1)
		} else {
			fmt.Fprintln(outWriter, 0)
		}
	}
}

func areLoginsSimilar(login, newLogin string) bool {
	if login == newLogin {
		return true
	}

	if len(login) != len(newLogin) {
		return false
	}

	for i := 0; i < len(login)-1; i++ {
		if login[i] == newLogin[i+1] && login[i+1] == newLogin[i] {
			if strings.HasPrefix(login[i+2:], newLogin[i+2:]) || strings.HasPrefix(newLogin[i+2:], login[i+2:]) {
				return true
			}
		}
	}

	return false
}
