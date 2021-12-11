package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"aoc/pkg/helpers"
)

type grid struct {
	field map[point]int
	h, w  int
}

type point struct {
	x, y int
}

func (p point) plus(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

func (g grid) inBound(p point) bool {
	if p.x >= g.w || p.x < 0 || p.y >= g.h || p.y < 0 {
		return false
	}
	return true
}

func newGrid(ps [][]int) grid {
	g := grid{field: make(map[point]int)}
	for i := range ps {
		for j := range ps[i] {
			g.field[point{j, i}] = ps[i][j]
		}
	}
	g.h, g.w = len(ps), len(ps[0])
	return g
}

func solve(g grid) (answer1, answer2 int) {
	flashes := make(map[point]bool)
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	for i := 0; i < 10000; i++ {
		for k := range g.field {
			g.field[k]++
			if g.field[k] > 9 {
				flashes[k] = true
			}
		}
	chainReaction:
		flashesPerStep := 0
		chain := false
		for k := range flashes {
			if flashes[k] {
				flashes[k] = false
				flashesPerStep++
				for _, dir := range dirs {
					key := k.plus(dir)
					if !g.inBound(key) {
						continue
					}
					g.field[key]++
					if g.field[key] == 10 {
						flashes[key] = true
						chain = true
					}
				}
			}
		}
		if chain {
			goto chainReaction
		}
		if flashesPerStep == g.h*g.w {
			// checks for case if the answer to part 2 comes before the answer to part 1
			if answer2 == 0 {
				answer2 = i + 1
			}
			if i > 99 {
				break
			}
		}
		for key := range flashes {
			g.field[key] = 0
			if i < 100 {
				answer1++
			}
		}
		flashes = make(map[point]bool)
	}

	return answer1, answer2
}

func parse(file io.Reader) [][]int {
	p := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := helpers.SliceAtoi(strings.Split(scanner.Text(), ""))
		p = append(p, s)
	}
	return p
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	ps := parse(file)
	g := newGrid(ps)
	a1, a2 := solve(g)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
