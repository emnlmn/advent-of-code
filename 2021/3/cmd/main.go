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
	binaries, err := readInput()

	if err != nil {
		panic(err)
	}

	leng := len(binaries[0])

	var oneCount = make([]int, leng)
	var zeroCount = make([]int, leng)
	for _, binaryString := range binaries {
		for key, char := range binaryString {
			digit := char - '0'
			switch digit {
			case 1:
				oneCount[key] = oneCount[key] + 1
			case 0:
				zeroCount[key] = zeroCount[key] + 1
			}
		}
	}

	gamma, epsilon := calculateRates(oneCount, zeroCount)


	g, err := strconv.ParseInt(gamma, 2, 64);
	if err != nil {
		fmt.Println(err)
	}

	e, err := strconv.ParseInt(epsilon, 2, 64);

	if err != nil {
		fmt.Println(err)
	}

	tot := g * e

	fmt.Println(tot)
}

func calculateRates(oneCount []int, zeroCount []int) (gamma string, epsilon string) {
	leng := len(oneCount)
	gammaArr := make([]string, leng)
	epsilonArr := make([]string, leng)
	for key, one := range oneCount {
		zero := zeroCount[key]
		if one > zero {
			gammaArr[key] = "1"
			epsilonArr[key] = "0"
		} else {
			gammaArr[key] = "0"
			epsilonArr[key] = "1"
		}
	}

	gamma = strings.Join(gammaArr, "")
	epsilon = strings.Join(epsilonArr, "")

	return gamma, epsilon
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
