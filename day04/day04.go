package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CountsString(c map[int]int) string {
	var strs []string
	keys := make([]int, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		strs = append(strs, fmt.Sprintf("%d: %d", k, c[k]))
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ", "))
}

func GetDefault(c map[int]int, key int, defaultVal int) int {
	v, found := c[key]
	if found {
		return v
	}
	return defaultVal
}

func main() {
	scratchpads := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line[strings.Index(line, ":")+1:], " | ")
		winningStr := parts[0]
		numbersStr := parts[1]

		winningSet := make(map[int]struct{})
		for _, nStr := range strings.Fields(winningStr) {
			n, _ := strconv.Atoi(nStr)
			winningSet[n] = struct{}{}
		}

		numbers := make([]int, 0)
		for _, nStr := range strings.Fields(numbersStr) {
			n, _ := strconv.Atoi(nStr)
			numbers = append(numbers, n)
		}

		nMatches := 0
		for _, n := range numbers {
			if _, found := winningSet[n]; found {
				nMatches++
			}
		}

		scratchpads[i] = GetDefault(scratchpads, i, 1)
		nExtraPads := scratchpads[i]

		for j := i + 1; j < i + nMatches + 1; j++ {
			scratchpads[j] = GetDefault(scratchpads, j, 1) + nExtraPads
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	sum := 0
	for _, v := range scratchpads {
		sum += v
	}
	fmt.Println(sum)
}
