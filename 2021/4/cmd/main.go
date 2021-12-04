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

type row = []int
type board = []row
type boards = []board

type boardMatch = [5][5]bool
type boardMatches = []boardMatch

func main() {
	numbers, boards, err := readInput()

	length := len(boards)
	var boardMatches = make(boardMatches, length)

	if err != nil {
		panic(err)
	}

	for _, drownedNumber := range numbers {
		for boardId, board := range boards {
			for rowId, row := range board {
				for numberId, number := range row {
					if number == drownedNumber {
						boardMatches[boardId][rowId][numberId] = true

						score := getBoardScore(board, boardMatches[boardId], rowId, numberId)

						if score > 0 {
							fmt.Println(score, number)
							fmt.Println(score * number)
							os.Exit(1)
						}
					}
				}
			}
		}
	}
}

func getBoardScore(board board, boardMatch boardMatch, rowId int, columnId int) (score int) {
	winRow := false
	for _, matchRow := range boardMatch[rowId] {
		if !matchRow {
			winRow = false
			break
		}

		winRow = true
	}

	if winRow {
		return calculateBoardScore(board, boardMatch)
	}

	winCol := false
	for _, row := range boardMatch {
		if !row[columnId] {
			winCol = false
			break
		}

		winCol = true;
	}

	if winCol {
		return calculateBoardScore(board, boardMatch)
	}

	return 0
}

func calculateBoardScore(b board, match boardMatch) (score int) {
	for rowId, row := range match {
		for columnId, matched := range row {
			if !matched {
				score += b[rowId][columnId]
			}
		}
	}

	return score
}

func readInput() (numbers []int, boards boards, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return numbers, boards, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return numbers, boards, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbersLine := scanner.Text()
	numbersLine = strings.Trim(numbersLine, " ")

	for _, number := range strings.Split(numbersLine, ",") {
		numberInt, _ := strconv.Atoi(number)
		numbers = append(numbers, numberInt)
	}

	scanner.Scan()

	board := make(board, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		row := make(row, 0)
		for _, number := range strings.Split(line, " ") {
			number = strings.Trim(number, " ")
			if len(number) == 0 {
				continue
			}

			numberInt, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			row = append(row, numberInt)
		}

		board = append(board, row)

		if len(line) == 0 {
			boards = append(boards, board)
			board = [][]int{}
			continue
		}
	}

	boards = append(boards, board)

	return numbers, boards, nil
}
