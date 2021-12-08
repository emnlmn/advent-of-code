package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	_, digits, err := readInput()

	digitLengthMap := map[int]int {
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}

	if err != nil {
		panic(err)
	}



	digitCount := 0

	for _, digitSegment := range digits {
		if _, ok := digitLengthMap[len(digitSegment)]; ok {
			digitCount += 1
		}
	}

	fmt.Println(digitCount)
}

func readInput() (segments []string, digits []string, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return segments, digits, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return segments, digits, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		parts := strings.Split(line, " | ")

		for _, seg := range strings.Split(parts[0], " ") {
			segments = append(segments, seg)
		}

		for _, seg := range strings.Split(parts[1], " ") {
			digits = append(digits, seg)
		}
	}

	return segments, digits, err
}
