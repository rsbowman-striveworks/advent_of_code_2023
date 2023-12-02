package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)

var digitStrings = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func scanLine(line string) (value int, remainingText string) {
	if len(line) <= 0 {
		return -1, ""
	}

	if (unicode.IsDigit(rune(line[0]))) {
		return int(line[0] - '0'), line[1:]
	}

	for i, digitString := range(digitStrings) {
		if strings.HasPrefix(line, digitString) {
			// handle "eightwo" and "sevenine" with the -1; by examinataion
			// these are the only overlapping cases
			return i, line[len(digitString) - 1:]
		}
	}

	return -1, line[1:]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	sum := 0
	var line string

	for scanner.Scan() {
		var lastValue *int = nil
		// firstValue := 0
		line = scanner.Text()
		// origLine := line


		for {
			v, newLine := scanLine(line)
			line = newLine

			if v >= 0 {
				if lastValue == nil {
					// first number on this line
					sum += v * 10
					// firstValue = v
				}
				lastValue = &v
			}

			if len(line) <= 0 {
				if lastValue == nil {
					fmt.Fprintf(os.Stderr, "no rune on line %d\n", lineNum);
					os.Exit(1)
				}

				// fmt.Printf("%-*s - %d + %d\n", 50, origLine, firstValue, *lastValue)
				sum += *lastValue
				break
			}
		}
		lineNum += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading lines: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", sum)
}
