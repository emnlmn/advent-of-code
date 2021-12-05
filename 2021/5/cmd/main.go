package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type point = struct {
	x int
	y int
}

type lineCoords = struct {
	begin point
	end   point
}

type terrain = map[int]map[int]int

func main() {
	coords, err := readInput()

	if err != nil {
		panic(err)
	}

	terrain := make(terrain, len(coords))
	for _, coord := range coords {
		if coord.begin.x == coord.end.x {
			x := coord.begin.x
			y1 := int(math.Max(float64(coord.begin.y), float64(coord.end.y)))
			y2 := int(math.Min(float64(coord.begin.y), float64(coord.end.y)))

			for i := y2; i <= y1; i++ {
				if _, ok := terrain[x]; ok {
					if _, ok := terrain[x][i]; ok {
						terrain[x][i] = terrain[x][i] + 1
					} else {
						terrain[x][i] = 1
					}
				} else {
					tmp := map[int]int{i: 1}
					terrain[x] = tmp
				}
			}

			continue
		}

		if coord.begin.y == coord.end.y {
			y := coord.begin.y
			x1 := int(math.Max(float64(coord.begin.x), float64(coord.end.x)))
			x2 := int(math.Min(float64(coord.begin.x), float64(coord.end.x)))

			for i := x2; i <= x1; i++ {
				if _, ok := terrain[i]; ok {
					if _, ok := terrain[i][y]; ok {
						terrain[i][y] = terrain[i][y] + 1
					} else {
						terrain[i][y] = 1
					}
				} else {
					tmp := map[int]int{y: 1}
					terrain[i] = tmp
				}
			}

			continue
		}

		// diagonal lines
		// {{8 0} {0 8}}
		// {{6 4} {2 0}}
		// {{0 0} {8 8}}
		// {{5 5} {8 2}}
		// . . 1 . . . . . .
		// . . . 1 . . . . .
		// . . . . 1 . . . 1
		// . . . . . 1 . 1 .
		// . . . . . . 2 . .
		// . . . . . 1 . . .
		// . . . . . . . . .
		// . . . . . . . . .
		var begin point
		var end point
		if coord.begin.x > coord.end.x {
			begin = coord.end
			end = coord.begin
		} else {
			begin = coord.begin
			end = coord.end
		}

		step := 0
		y := begin.y
		for i := begin.x; i <= end.x; i++ {
			if begin.y < end.y {
				y = begin.y + step
			} else {
				y = begin.y - step
			}

			if _, ok := terrain[i]; ok {
				if _, ok := terrain[i][y]; ok {
					terrain[i][y] = terrain[i][y] + 1
				} else {
					terrain[i][y] = 1
				}
			} else {
				tmp := map[int]int{y: 1}
				terrain[i] = tmp
			}

			step++
		}
	}

	count := 0
	for _, x := range terrain {
		for _, y := range x {
			if y > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func readInput() (coords []lineCoords, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return coords, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return coords, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		r, _ := regexp.Compile(`(\d+),(\d+) -> (\d+),(\d+)`)

		matches := r.FindStringSubmatch(line)

		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		coord := lineCoords{
			point{
				x: x1,
				y: y1,
			},
			point{
				x: x2,
				y: y2,
			},
		}

		coords = append(coords, coord)
	}

	return coords, err
}
