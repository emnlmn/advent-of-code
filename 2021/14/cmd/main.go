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

	for i := 0; i < len(start) -1; i++ {
		firstElement := string(start[i])
		secondElement := string(start[i+1])
		pair := firstElement + secondElement

		countPair(pair, count)
	}

	for i := 0; i < 41; i++ {
		if i == 10 || i == 40 {
			tempCount := map[string]int{}
			for s := range count {
				tempCount[string(s[0])] += count[s]
			}
			tempCount[string(start[len(start)-1])] += 1

			sortedCount := sortCount(tempCount)
			tot := sortedCount[0].count - sortedCount[len(sortedCount)-1].count
			fmt.Println(tot)
		}

		count2 := map[string]int{}
		for pair, c := range count {
			np1 := string(pair[0]) + transformations[pair]
			np2 := transformations[pair] + string(pair[1])
			count2[np1] += c
			count2[np2] += c
		}
		count = count2
	}

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

func countPair(s string, count map[string]int) {
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
