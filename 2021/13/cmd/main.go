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

type fold struct {
	dir string
	pos int
}

func main() {
	t := time.Now()
	sheet, instructions, err := parseInput()

	if err != nil {
		panic(err)
	}

	for _, instruction := range instructions {
		cnt := 0
		switch instruction.dir {
		case "x":
			sheet = foldX(sheet, instruction.pos)
		case "y":
			sheet = foldY(sheet, instruction.pos)
		}

		for _, x := range sheet {
			for _, n := range x {
				if n > 0 {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
				cnt += n
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
		fmt.Println(cnt)
		fmt.Printf("\n")
	}

	cnt := 0
	for _, x := range sheet {
		for _, n := range x {
			cnt += n
		}
	}

	fmt.Println(cnt)
	fmt.Println("time:", time.Since(t))
}

func foldY(sheet [][]int, pos int) [][]int {
	for i, j := pos+1, pos-1; i < len(sheet); i, j = i+1, j-1 {
		for x := 0; x < len(sheet[i]); x++ {
			if sheet[i][x] == 1 {
				sheet[j][x] = sheet[i][x]
			}
		}
	}

	sheet = sheet[:pos]

	return sheet
}

func foldX(sheet [][]int, pos int) [][]int {
	for i, j := pos+1, pos-1; i < len(sheet[0]); i, j = i+1, j-1 {
		for y := range sheet {
			if sheet[y][i] == 1 {
				sheet[y][j] = sheet[y][i]
			}
		}
	}

	var s [][]int
	for _, l := range sheet {
		s = append(s, l[:pos])
	}

	return s
}

func parseInput() (sheet [][]int, instructions []fold, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return sheet, instructions, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return sheet, instructions, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sheet1 := [9999][9999]int{}

	maxY := 0
	maxX := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line == "" {
			break
		}

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		sheet1[y][x] = 1
	}

	sheet2 := sheet1[:maxY+1]
	var sheet3 [][]int
	for _, l := range sheet2 {
		nl := []int(nil);
		for i, n := range l {
			if i <= maxX {
				nl = append(nl, n)
			}
		}

		sheet3 = append(sheet3, nl)
	}

	sheet = sheet3

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "fold along ")

		parts := strings.Split(line, "=")
		c, _ := strconv.Atoi(parts[1])

		f := fold{
			parts[0],
			c,
		}

		instructions = append(instructions, f)
	}

	return sheet, instructions, err
}
