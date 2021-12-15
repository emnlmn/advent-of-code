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
	"time"
)

type pos struct {
	x    int
	y    int
	risk int
}

type Stack struct {
	items []pos
	cnt   int
}

func (s *Stack) push(value pos) {
	s.items = append(s.items, value)
	s.cnt++
	return
}

func (s *Stack) pop() (point pos, err error) {
	if s.cnt == 0 {
		return point, errors.New("stack Underflow, No values to remove")
	}

	point = s.items[s.cnt-1]
	s.items = s.items[:s.cnt-1]
	s.cnt--

	return point, nil
}

func main() {
	t := time.Now()
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	s := Stack{}

	s.push(pos{
		0,
		0,
		0,
	})

	maxY := len(matrix) -1
	maxX := len(matrix[len(matrix)-1]) -1

	fmt.Println(maxY)
	fmt.Println(maxX)

	for s.cnt > 0 {
		current, _ := s.pop()

		fmt.Println(current.y, current.x, s.cnt)

		if current.y == maxY && current.x == maxX {
			fmt.Println(current.risk)
			break
		}

		ry := 0
		if current.y + 1 <= maxY {
			ry = matrix[current.y+1][current.x]
		}

		rx := 0
		if current.x + 1 <= maxY {
			rx = matrix[current.y][current.x+1]
		}

		if rx == 0 {
			s.push(pos{
				current.x,
				current.y+1,
				current.risk + rx,
			})

			continue
		}

		if ry == 0 {
			s.push(pos{
				current.x+1,
				current.y,
				current.risk + rx,
			})

			continue
		}

		if rx < ry {
			s.push(pos{
				current.x+1,
				current.y,
				current.risk + rx,
			})
		} else {
			s.push(pos{
				current.x,
				current.y+1,
				current.risk + rx,
			})
		}
	}

	fmt.Println("time:", time.Since(t))
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

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line == "" {
			continue
		}

		var row []int
		for _, si := range strings.Split(line, "") {
			i, _ := strconv.Atoi(si)
			row = append(row, i)
		}

		matrix = append(matrix, row)
	}

	return matrix, err
}
