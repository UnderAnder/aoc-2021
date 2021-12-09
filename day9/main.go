package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type points [][]int

type point struct {
	x, y int
}

func parse(file io.Reader) points {
	p := make(points, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := sliceAtoi(strings.Split(scanner.Text(), ""))
		p = append(p, s)
	}
	return p
}

func sliceAtoi(s []string) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		result = append(result, mustAtoi(v))
	}
	return result
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int from input")
	}
	return i
}

func solve(ps points) (answer1, answer2 int) {
	answer1, answer2 = 0, 1
	lps := make([]point, 0, len(ps))
	basins := make([]map[point]bool, 0, len(ps))

	for i := range ps {
		for j := range ps[i] {
			if isLowPoint(i, j, ps) {
				lps = append(lps, point{i, j})
			}
		}
	}
	for _, v := range lps {
		answer1 += ps[v.x][v.y] + 1
		basin := make(map[point]bool)
		allowMove(v, ps, basin)
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

func isLowPoint(i, j int, ps points) bool {
	result := true
	pnt := ps[i][j]
	if j+1 < len(ps[i]) {
		result = result && ps[i][j+1] > pnt
	}
	if j > 0 {
		result = result && ps[i][j-1] > pnt
	}
	if i+1 < len(ps) {
		result = result && ps[i+1][j] > pnt
	}
	if i > 0 {
		result = result && ps[i-1][j] > pnt
	}
	return result
}

func allowMove(p point, ps points, b map[point]bool) {
	if b[p] {
		return
	}
	b[p] = true

	pnt := point{p.x, p.y + 1}
	if p.y+1 < len(ps[p.x]) && ps[p.x][p.y+1] != 9 && !b[pnt] {
		b[pnt] = false
		allowMove(pnt, ps, b)
	}
	pnt = point{p.x, p.y - 1}
	if p.y > 0 && ps[p.x][p.y-1] != 9 && !b[pnt] {
		b[pnt] = false
		allowMove(pnt, ps, b)
	}
	pnt = point{p.x + 1, p.y}
	if p.x+1 < len(ps) && ps[p.x+1][p.y] != 9 && !b[pnt] {
		b[pnt] = false
		allowMove(pnt, ps, b)
	}
	pnt = point{p.x - 1, p.y}
	if p.x > 0 && ps[p.x-1][p.y] != 9 && !b[pnt] {
		b[pnt] = false
		allowMove(pnt, ps, b)
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
	p := parse(file)
	a1, a2 := solve(p)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
