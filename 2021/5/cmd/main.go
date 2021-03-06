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
	firstGeneration, err := readInput()
	fmt.Println(firstGeneration)


	if err != nil {
		panic(err)
	}

	var newElements int
	var newGeneration []int

	for _, day := range firstGeneration {
		newElements = 0
		newGeneration = make([]int, len(firstGeneration))
		day--
		if day == 0 {
			day = 6
			newElements++
		}
		newGeneration = append(newGeneration, day)

		for i := 0; i < newElements; i++ {
			newGeneration = append(newGeneration, 8)
		}

	}
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
