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

func main() {
	total, err := parseInput()

	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func parseInput() (total int, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return total, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return total, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		parts := strings.Split(line, " | ")

		inputs := sortS(strings.Split(parts[0], " "))
		outputs := sortS(strings.Split(parts[1], " "))

		total += decode(inputs, outputs)
	}

	return total, err
}

func decode(input []string, output []string) (total int) {
	digitLengthMap := map[int]int {
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	encoding := map[int]string{}

	//      aaaa
	//     b    c
	//     b    c
	//      dddd
	//     e    f
	//     e    f
	//      gggg

	for _, digitSegment := range input {
		if digit, ok := digitLengthMap[len(digitSegment)]; ok {
			encoding[digit] = digitSegment
		}
	}

	for _, digitSegment := range input {
		if len(digitSegment) == 5 && len(leftBisect(encoding[1], digitSegment)) == 0 {
			encoding[3] = digitSegment
			continue
		}

		if len(digitSegment) == 6 {
			switch {
			// Solve for 0, 6, and 9
			case len(leftBisect(encoding[1], digitSegment)) == 1:
				encoding[6] = digitSegment
			case len(leftBisect(encoding[4], digitSegment)) == 0:
				encoding[9] = digitSegment
			case len(leftBisect(encoding[4], digitSegment)) == 1:
				encoding[0] = digitSegment
			}
		}
	}


	rightTop := leftBisect(encoding[1], encoding[6])[0]

	for _, digitSegment := range input {
		if len(digitSegment) != 5 || digitSegment == encoding[3] {
			continue
		}

		if strings.Contains(digitSegment, string(rightTop)) {
			encoding[2] = digitSegment
		} else {
			encoding[5] = digitSegment
		}
	}

	for _, outputSegment := range output {
		total *= 10
		for k, v := range encoding {
			if v == outputSegment {
				total += k
			}
		}
	}

	return total
}

// returns chars from a not present in b
func leftBisect(a, b string) string {
	result := strings.Builder{}

search:
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				continue search
			}
		}

		result.WriteRune(aa)
	}

	return result.String()
}

func sortS(s []string) []string {
	result := make([]string, len(s))

	for i, str := range s {
		v := []rune(str)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})

		result[i] = string(v)
	}

	return result
}