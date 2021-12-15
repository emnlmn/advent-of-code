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
	"time"
)

type pos struct {
	x    int
	y    int
	risk int
}

type Queue struct {
	items []pos
	cnt   int
}

func (s *Queue) push(value pos) {
	s.items = append(s.items, value)
	s.cnt++
	return
}

func (s *Queue) popLeft() (value pos, err error) {
	if s.cnt == 0 {
		return value, errors.New("stack Underflow, No values to remove")
	}

	value = s.items[0]
	s.items = s.items[1:]
	s.cnt--

	return value, nil
}
func main() {
	t := time.Now()
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	q := Queue{}

	q.push(pos{
		0,
		0,
		0,
	})

	var riskLevels []int

	maxY := len(matrix) -1
	maxX := len(matrix[len(matrix)-1]) -1

	fmt.Println(maxY)
	fmt.Println(maxX)

	for q.cnt > 0 {
		current, _ := q.popLeft()

		if current.y == maxY && current.x == maxX {
			riskLevels = append(riskLevels, current.risk)
			continue
		}

		if current.y + 1 <= maxY {
			q.push(pos{
				current.x,
				current.y+1,
				current.risk + matrix[current.y+1][current.x],
			})
		}

		if current.x + 1 <= maxY {
			q.push(pos{
				current.x+1,
				current.y,
				current.risk + matrix[current.y][current.x+1],
			})
		}
	}

	sort.Ints(riskLevels)

	fmt.Println(riskLevels[0])
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
