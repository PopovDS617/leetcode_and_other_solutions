package ozon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckSystemUnderPressure() {
	scanner := bufio.NewScanner(os.Stdin)
	outWriter := bufio.NewWriter(os.Stdout)
	defer outWriter.Flush()

	var n int

	_ = scanner.Scan()

	line := scanner.Text()
	line = strings.Trim(line, "\n\r\t")

	n, _ = strconv.Atoi(line)

	for i := 0; i < n; i++ {

		scanner.Scan()
		line = scanner.Text()
		line = strings.Trim(line, "\n\r\t")

		res := resolver(line)

		if !res {
			fmt.Fprint(outWriter, "NO\n")
		} else {
			fmt.Fprint(outWriter, "YES\n")
		}

	}

}

func resolver(line string) bool {

	started := false
	canceled := false
	restarted := false
	ended := false

	for _, c := range line {

		if c == 'M' {
			ended = false
			canceled = false
			restarted = false
			if !started {
				started = true

			} else {
				return false
			}
		} else if c == 'R' {
			if restarted {
				return false
			}
			if !started || canceled {
				return false
			}
			restarted = true
			started = true
		} else if c == 'C' {
			if !started {
				return false
			}

			started = false
			ended = false
		} else if c == 'D' {
			if !started || canceled || ended {
				return false
			}
			canceled = false
			started = false
			ended = true
			restarted = false
		}
	}

	return ended

}
