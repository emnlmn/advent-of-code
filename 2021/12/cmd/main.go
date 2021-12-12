package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"unicode"
)

type caveName string

type cave struct {
	name  caveName
	paths map[caveName]*cave
}

type pos struct {
	cave              cave
	visitedSmallCaves map[caveName]bool
	twice             bool
}

func (c *cave) connect(to cave) {
	if _, ok := c.paths[to.name]; ok {
		return
	}

	to.paths[c.name] = c
	c.paths[to.name] = &to
}

func (c *cave) isSmall() bool {
	for _, r := range c.name {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func newCave(name caveName) cave {
	return cave{
		name,
		make(map[caveName]*cave),
	}
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
	cavePaths, err := parseInput()

	if err != nil {
		panic(err)
	}

	caves := buildCaves(cavePaths)

	q := Queue{}
	q.push(pos{
		caves["start"],
		map[caveName]bool{"start": true},
		false,
	})

	cnt := 0

	for q.cnt > 0 {
		cpos, _ := q.popLeft()

		if cpos.cave.name == "end" {
			cnt++
			continue
		}

		for _, ccave := range cpos.cave.paths {
			visited := false

			if _, v := cpos.visitedSmallCaves[ccave.name]; v {
				visited = true
			}

			nvm := cloneVisitedSmallCaves(cpos.visitedSmallCaves)
			if !visited {
				if ccave.isSmall() {
					nvm[ccave.name] = true
				}

				q.push(pos{
					*ccave,
					nvm,
					cpos.twice,
				})
			} else if visited && !cpos.twice && ccave.name != "end" && ccave.name != "start" {
				q.push(pos{
					*ccave,
					nvm,
					true,
				})
			}
		}
	}

	fmt.Println(cnt)
}

func cloneVisitedSmallCaves(m map[caveName]bool) map[caveName]bool {
	nvm := map[caveName]bool{}
	for n, t := range m {
		nvm[n] = t
	}

	return nvm
}

func buildCaves(paths []string) (caves map[caveName]cave) {
	caves = map[caveName]cave{}

	for _, p := range paths {
		startEnd := strings.Split(p, "-")
		caveNameFrom := caveName(startEnd[0])
		caveNameTo := caveName(startEnd[1])

		caveTo, ok := caves[caveNameTo]

		if !ok {
			caveTo = newCave(caveNameTo)
			caves[caveTo.name] = caveTo
		}

		caveFrom, ok := caves[caveNameFrom]
		if !ok {
			caveFrom = newCave(caveNameFrom)
			caves[caveFrom.name] = caveFrom
		}

		caveFrom.connect(caveTo)
	}

	return caves
}

func parseInput() (cavePaths []string, err error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return cavePaths, errors.New("could not get runtime path")
	}

	filepath := path.Join(path.Dir(filename), "../input.txt")

	file, err := os.Open(filepath)

	if err != nil {
		return cavePaths, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		cavePaths = append(cavePaths, line)
	}

	return cavePaths, err
}
