package main

import (
	"bufio"
	"day2"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	submarinePosition := day2.SubmarinePosition{}

	lines, err := readInput()

	if err != nil {
		panic(err)
	}

	submarinePosition = executeMoves(lines, submarinePosition)

	fmt.Println(submarinePosition.Value())
}

func executeMoves(lines []string, position day2.SubmarinePosition) day2.SubmarinePosition {
	for _, commandLine := range lines {
		commandParts := strings.Split(commandLine, " ")

		command := commandParts[0]
		amount, err := strconv.Atoi(commandParts[1])

		if err != nil {
			panic(err)
		}

		switch command {
		case "forward":
			position = position.GoForward(amount)
		case "down":
			position = position.GoDown(amount)
		case "up":
			position = position.GoUp(amount)
		}
	}

	return position
}

func readInput() (lines []string, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return lines, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		lines = append(lines, line)
	}

	return lines, nil
}
