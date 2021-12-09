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
	lowPoints := make(map[Point]int, 0)

	for i := 0; i < ly; i++ {
		for j := 0; j < lx; j++ {
			p := Point{j, i}
			lowPoint, found := dfs(matrix, p, visited)
			if found {
				lowPoints[lowPoint] = matrix[lowPoint.y][lowPoint.x]
			}
		}
	}

	tot := 0
	for _, k := range lowPoints {
		tot += k + 1
	}

	fmt.Println(tot)
	fmt.Println(lowPoints)
}

func dfs(matrix [][]int, sp Point, visited map[Point]bool) (lower Point, found bool) {
	st := Stack{}
	st.push(sp)

	for st.cnt > 0 {
		p, err := st.pop()

		if err != nil {
			panic(err)
		}

		if _, ok := visited[p]; ok {
			break
		}

		height := matrix[p.y][p.x]

		hasLower := false

		if isValid(matrix, p.x+1, p.y) && height >= matrix[p.y][p.x+1] {
			hasLower = true

			p := Point{p.x + 1, p.y}
			st.push(p)
		}

		if isValid(matrix, p.x-1, p.y) && height >= matrix[p.y][p.x-1] {
			hasLower = true

			p := Point{p.x - 1, p.y}
			st.push(p)
		}

		if isValid(matrix, p.x, p.y+1) && height >= matrix[p.y+1][p.x] {
			hasLower = true

			p := Point{p.x, p.y + 1}
			st.push(p)
		}

		if isValid(matrix, p.x-1, p.y-1) && height >= matrix[p.y-1][p.x] {
			hasLower = true

			p := Point{p.x, p.y - 1}
			st.push(p)
		}

		visited[p] = true
		if !hasLower {
			lower = p
		}
	}

	height := matrix[lower.y][lower.x]

	if isValid(matrix, lower.x+1, lower.y) && height > matrix[lower.y][lower.x+1] {
		return lower, false
	}

	if isValid(matrix, lower.x-1, lower.y) && height > matrix[lower.y][lower.x-1] {
		return lower, false
	}

	if isValid(matrix, lower.x, lower.y+1) && height > matrix[lower.y+1][lower.x] {
		return lower, false
	}

	if isValid(matrix, lower.x, lower.y-1) && height > matrix[lower.y-1][lower.x] {
		return lower, false
	}

	return lower, true
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
