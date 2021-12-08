package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()

	if err != nil {
		panic(err)
	}

	fuels := make(map[int]int, len(input))

	max := 0
	min := 0
	for _, e := range input {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}

	for pos := min; pos <= max; pos++ {
		for _, crab := range input {
			n := int(math.Abs(float64(pos - crab)))

			n = n*(n+1)/2

			if _, ok := fuels[pos]; ok {
				fuels[pos] += n
				continue
			}
			fuels[pos] = n
		}
	}

	f := make([]int, 0)
	for _, v := range fuels {
		f = append(f, v)
	}

	sort.Ints(f)

	fmt.Println(f[0])
}


func readInput() (positions []int, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return positions, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return positions, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, " ")

	for _, element := range strings.Split(line, ",") {
		intElement, _ := strconv.Atoi(element)
		positions = append(positions, intElement)
	}

	return positions, err
}
