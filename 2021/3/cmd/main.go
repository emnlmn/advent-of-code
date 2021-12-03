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

	oxigen, co2 := calculateRates2(binaries)

	g, err := strconv.ParseInt(oxigen, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	e, err := strconv.ParseInt(co2, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	tot := g * e

	fmt.Println(tot)
}

func calculateCounts(binaries []string, significantDigit int) (oneCount int, zeroCount int) {
	oneCount = 0
	zeroCount = 0

	for _, binaryString := range binaries {
		digit := binaryString[significantDigit] - '0'

		switch digit {
		case 1:
			oneCount ++
		case 0:
			zeroCount ++
		}
	}
	return oneCount, zeroCount
}

func calculateRates2(binaries []string) (oxigen string, co2 string) {
	oxBinaries := binaries

	significantDigit := 0
	for len(oxBinaries) > 1 {
		oneCount, zeroCount := calculateCounts(oxBinaries, significantDigit)
		isOne := oneCount > zeroCount
		if oneCount == zeroCount {
			isOne = true
		}

		var b []string
		for _, binary := range oxBinaries {
			digit := binary[significantDigit] - '0'
			switch digit {
			case 1:
				if isOne {
					b = append(b, binary)
				}
			case 0:
				if !isOne {
					b = append(b, binary)
				}
			}
		}
		significantDigit++
		oxBinaries = b
	}

	co2Binaries := binaries
	significantDigit = 0
	for len(co2Binaries) > 1 {
		oneCount, zeroCount := calculateCounts(co2Binaries, significantDigit)
		isOne := oneCount > zeroCount
		if oneCount == zeroCount {
			isOne = true
		}

		var b []string
		for _, binary := range co2Binaries {
			digit := binary[significantDigit] - '0'
			switch digit {
			case 1:
				if !isOne {
					b = append(b, binary)
				}
			case 0:
				if isOne {
					b = append(b, binary)
				}
			}
		}

		significantDigit++
		co2Binaries = b
	}

	oxigen = oxBinaries[0]
	co2 = co2Binaries[0]

	return oxigen, co2
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
