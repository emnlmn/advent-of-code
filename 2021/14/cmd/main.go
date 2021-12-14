package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"time"
)

type elementCount struct {
	element string
	count   int
}

func main() {
	t := time.Now()
	start, transformations, err := parseInput()

	if err != nil {
		panic(err)
	}

	count := map[string]int{}

	for _, s := range start {
		countChart(string(s), count)
	}

	for i := 0; i < 10; i++ {
		polymer := start
		newStart := ""
		for sPos := 0; sPos < len(polymer)-1; sPos++ {
			firstElement := string(polymer[sPos])
			secondElement := string(polymer[sPos+1])
			pair := firstElement + secondElement

			newChar := transformations[pair]

			if sPos == 0 {
				newStart = newStart + firstElement + newChar + secondElement
			} else {
				newStart = newStart + newChar + secondElement
			}

			countChart(newChar, count)
		}

		start = newStart
	}

	sortedCount := sortCount(count)
	fmt.Println(sortedCount)

	tot := sortedCount[0].count - sortedCount[len(sortedCount)-1].count

	fmt.Println(tot)

	fmt.Println("time:", time.Since(t))
}

func sortCount(count map[string]int) []elementCount {
	var s []elementCount

	for element, count := range count {
		s = append(s, elementCount{element, count})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].count > s[j].count
	})

	return s
}

func countChart(s string, count map[string]int) {
	if _, ok := count[s]; !ok {
		count[s] = 0
	}

	count[s]++
}

func insert(toInsert string, pos int, string string) string {
	return string[:pos] + toInsert + string[pos:]
}

func parseInput() (start string, transformations map[string]string, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return start, transformations, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return start, transformations, err
	}

	defer file.Close()

	transformations = make(map[string]string)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	start = strings.Trim(firstLine, " ")
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")

		transformations[parts[0]] = parts[1]
	}

	return start, transformations, err
}
