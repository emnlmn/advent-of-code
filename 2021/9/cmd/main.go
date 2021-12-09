package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Stack struct {
	items []Point
	cnt   int
}

func (s *Stack) push(value Point) {
	s.items = append(s.items, value)
	s.cnt++
	return
}

func (s *Stack) pop() (point Point, err error) {
	if s.cnt == 0 {
		return point, errors.New("stack Underflow, No values to remove")
	}

	point = s.items[s.cnt-1]
	s.items = s.items[:s.cnt-1]
	s.cnt--

	return point, nil
}

func main() {
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	ly := len(matrix)
	lx := len(matrix[0])

	visited := make(map[Point]bool, 0)
	basins := make([][]Point, 0)

	for i := 0; i < ly; i++ {
		for j := 0; j < lx; j++ {
			p := Point{j, i}
			basin := dfs(matrix, p, visited)
			if len(basin) > 0 {
				basins = append(basins, basin)
			}
		}
	}

	lengths := []int{}
	for _, b := range basins {
		lengths = append(lengths, len(b))
	}

	sort.Ints(lengths)

	l := len(lengths)
	fmt.Println(lengths[l-1] * lengths[l-2] * lengths[l-3])
}

func dfs(matrix [][]int, sp Point, visited map[Point]bool) (basin []Point) {
	st := Stack{}
	st.push(sp)

	for st.cnt > 0 {
		p, err := st.pop()

		if err != nil {
			panic(err)
		}

		if _, ok := visited[p]; ok {
			continue
		}

		if isValid(matrix, p.x+1, p.y) && matrix[p.y][p.x+1] != 9 {
			p := Point{p.x + 1, p.y}
			st.push(p)
		}

		if isValid(matrix, p.x-1, p.y) && matrix[p.y][p.x-1] != 9 {
			p := Point{p.x - 1, p.y}
			st.push(p)
		}

		if isValid(matrix, p.x, p.y+1) && matrix[p.y+1][p.x] != 9 {
			p := Point{p.x, p.y + 1}
			st.push(p)
		}

		if isValid(matrix, p.x, p.y-1) && matrix[p.y-1][p.x] != 9 {
			p := Point{p.x, p.y - 1}
			st.push(p)
		}

		if matrix[p.y][p.x] != 9 {
			basin = append(basin, p)
		}

		visited[p] = true
	}

	return basin
}

func isValid(matrix [][]int, x int, y int) bool {
	if len(matrix) <= y || (y < 0) {
		return false
	}

	if len(matrix[y]) <= x || (x < 0) {
		return false
	}

	return true
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
