package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()

	firstGeneration := map[int]int{}
	for _, d := range input {
		if _, ok := firstGeneration[d]; ok {
			firstGeneration[d] += 1
		} else {
			firstGeneration[d] = 1
		}
	}

	if err != nil {
		panic(err)
	}

	for gen := 0; gen < 256; gen++ {
		newGeneration := map[int]int{}

		fmt.Println(gen)

		for day, count := range firstGeneration {
			if day == 0 {
				newGeneration[6] += count
				newGeneration[8] += count
			} else {
				newGeneration[day-1] += count
			}
		}

		firstGeneration = newGeneration
	}
	tot := 0
	for _, cnt := range firstGeneration {
		tot += cnt
	}

	fmt.Println(tot)
}

func readInput() (lanternLights []int, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return lanternLights, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return lanternLights, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, " ")

	for _, element := range strings.Split(line, ",") {
		intElement, _ := strconv.Atoi(element)
		lanternLights = append(lanternLights, intElement)
	}

	return lanternLights, err
}
