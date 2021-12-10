package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
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

func main() {
	matrix, err := parseInput()

	if err != nil {
		panic(err)
	}

	pMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	pRMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	scoreBoard := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	var errors []string

	var missings [][]string

	for _, row := range matrix {
		e, m := parseRow(row, pMap, pRMap)
		errors = append(errors, e)
		if len(m) > 0 {
			missings = append(missings, m)
		}
	}

	var scores []int
	for _, p := range missings {
		score := 0
		for _, m := range p {
			if s, ok := scoreBoard[m]; ok {
				score = score * 5
				score += s
			}
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	l := len(scores)

	fmt.Println(scores[l/2])
}

func parseRow(row []string, pMap map[string]string, pRMap map[string]string) (error string, missing []string) {
	st := Stack{}

	for _, parenthesis := range row {
		if _, ok := pMap[parenthesis]; ok {
			st.push(parenthesis)
			continue
		}

		if op, ok := pRMap[parenthesis]; ok {
			toMatch, err := st.pop()
			toMatchS := fmt.Sprintf("%v", toMatch)

			if err != nil {
				panic(err)
			}

			if op != toMatchS {
				return parenthesis, missing
			}

			continue
		}
	}

	for st.cnt > 0 {
		mp, err := st.pop()

		if err != nil {
			panic(err)
		}

		mpS := fmt.Sprintf("%v", mp)

		if cp, ok := pMap[mpS]; ok {
			missing = append(missing, cp)
			continue
		}
	}

	return error, missing
}

func parseInput() (matrix [][]string, err error) {
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
	var row []string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		parts := strings.Split(line, "")

		row = make([]string, 0)

		for _, v := range parts {
			row = append(row, v)
		}

		matrix = append(matrix, row)
	}

	return matrix, err
}
