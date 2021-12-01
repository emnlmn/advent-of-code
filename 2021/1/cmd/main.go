package main

import (
	"bufio"
	"day1"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	var measurements []int

	measurements, err := getMeasurementsFromInputFile()

	if err != nil {
		fmt.Printf("error obtaining input data: %s", err)
		return
	}

	slidedMeasurements := day1.SlideMeasurements(measurements)

	increments := day1.CountIncrements(slidedMeasurements)

	fmt.Println(increments)
}

func getMeasurementsFromInputFile() (measurements day1.Measurements, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return measurements, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return measurements, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		measurement, err := strconv.Atoi(line)

		if err != nil {
			return measurements, err
		}

		measurements = append(measurements, measurement)
	}

	return measurements, nil
}
