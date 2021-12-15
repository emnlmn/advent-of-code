package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	x        int
	y        int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	pi := pq[i]
	pj := pq[j]

	if pi != nil && pj != nil {
		return pq[i].priority < pq[j].priority
	}

	return true
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, x int, y int, priority int) {
	item.x = x
	item.y = y
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	t := time.Now()
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	q := make(PriorityQueue, 1)
	q[0] = &Item{
		0,
		0,
		0,
		0,
	}
	heap.Init(&q)

	maxY := len(matrix) - 1
	maxX := len(matrix[len(matrix)-1]) - 1

	visited := map[int]map[int]bool{}

	type dir struct {
		x int
		y int
	}

	dirs := []dir{
		dir{0, 1},
		dir{0, -1},
		dir{1, 0},
		dir{-1, 0},
	}

	for q.Len() > 0 {
		current := heap.Pop(&q).(*Item)

		if _, ok := visited[current.y]; !ok {
			visited[current.y] = map[int]bool{}
		}

		if cy, ok := visited[current.y]; ok {
			if _, ok2 := cy[current.x]; ok2 {
				continue
			}
		}
		visited[current.y][current.x] = true

		if current.y == maxY && current.x == maxX {
			fmt.Println(current.priority)
			break
		}

		for _, dir := range dirs {
			ry := dir.y + current.y
			rx := dir.x + current.x

			if ry > maxY || rx > maxX || ry < 0 || rx < 0 {
				continue
			}

			is := fmt.Sprintf("%d%d", ry, rx)
			i, _ := strconv.Atoi(is)
			heap.Push(&q, &Item{
				rx,
				ry,
				current.priority + matrix[ry][rx],
				i,
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
