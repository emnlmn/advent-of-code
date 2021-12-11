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

type Stack struct {
	items []interface{}
	cnt   int
}

func (s *Stack) push(value interface{}) {
	s.items = append(s.items, value)
	s.cnt++
	return
}

func (s *Stack) pop() (value interface{}, err error) {
	if s.cnt == 0 {
		return value, errors.New("stack Underflow, No values to remove")
	}

	value = s.items[s.cnt-1]
	s.items = s.items[:s.cnt-1]
	s.cnt--

	return value, nil
}

var gTot int = 0

func main() {
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	found := false
	for n := 1; !found; n++ {
		flashed := map[string]bool{}
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix); j++ {
				matrix[i][j]++
			}
		}

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix); j++ {
				if matrix[i][j] > 9 {
					flash(matrix, i, j, flashed)
				}
			}
		}

		if isAll0(matrix) {
			found = true
			fmt.Println("step:", n)
		}
	}

	fmt.Println("flashes:", gTot)

	for _, k := range matrix {
		for _, o := range k {
			if o == 0 {
				fmt.Printf("\033[1m%d\033[0m", o)
			} else {
				fmt.Printf("%d", o)
			}
		}
		fmt.Printf("\n")
	}
}

func isAll0(matrix [][]int) bool {
	for _, k := range matrix {
		for _, o := range k {
			if o != 0 {
				return false
			}
		}
	}
	return true
}

func flash(matrix [][]int, i int, j int, flashed map[string]bool) {
	if (i > 9) || (j > 9) || (i < 0) || (j < 0) {
		return
	}

	key := fmt.Sprintf("%d%d", i, j)

	if _, ok := flashed[key]; ok {
		return
	}

	matrix[i][j]++

	if matrix[i][j] > 9 {
		matrix[i][j] = 0
		gTot++
		flashed[key] = true

		flash(matrix, i+1, j, flashed)
		flash(matrix, i+1, j+1, flashed)
		flash(matrix, i+1, j-1, flashed)
		flash(matrix, i-1, j, flashed)
		flash(matrix, i-1, j+1, flashed)
		flash(matrix, i-1, j-1, flashed)
		flash(matrix, i, j+1, flashed)
		flash(matrix, i, j-1, flashed)
	}
}

func parseInput() (matrix [][]int, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return matrix, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return matrix, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var row []int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		parts := strings.Split(line, "")

		row = make([]int, 0)

		for _, v := range parts {
			vi, _ := strconv.Atoi(v)
			row = append(row, vi)
		}

		matrix = append(matrix, row)
	}

	return matrix, err
}
