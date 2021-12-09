package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func parse(file io.Reader) [][]int {
	p := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := helpers.SliceAtoi(strings.Split(scanner.Text(), ""))
		p = append(p, s)
	}
	return p
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
	answer1, answer2 = 0, 1
	lps := make([]point, 0, len(g.field))
	basins := make([]map[point]bool, 0, len(g.field))
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for k := range g.field {
		if isLowPoint(k, g, dirs) {
			lps = append(lps, k)
		}
	}
	for _, v := range lps {
		answer1 += g.field[v] + 1
		basin := make(map[point]bool)
		findBasin(basin, v, dirs, g)
		basins = append(basins, basin)
	}

	basinLens := make([]int, 0, len(basins))
	for _, basin := range basins {
		basinLens = append(basinLens, len(basin))
	}
	sort.Ints(basinLens)
	for _, v := range basinLens[len(basinLens)-3:] {
		answer2 *= v
	}
	return answer1, answer2
}

func isLowPoint(p point, g grid, ds []point) bool {
	result := true
	for _, d := range ds {
		if !g.inBound(p.plus(d)) {
			continue
		}
		result = result && g.field[p.plus(d)] > g.field[p]
	}
	return result
}

func findBasin(b map[point]bool, p point, ds []point, g grid) {
	if b[p] {
		return
	}
	b[p] = true
	for _, d := range ds {
		pnt := p.plus(d)
		if !g.inBound(pnt) || b[pnt] || g.field[pnt] == 9 {
			continue
		}
		b[pnt] = false
		findBasin(b, pnt, ds, g)
	}
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
