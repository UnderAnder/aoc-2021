package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"aoc/pkg/helpers"
)

type queue struct {
	items []point
}

func (q *queue) add(p point, g grid) {
	if len(q.items) == 0 {
		q.items = append(q.items, p)
		return
	}
	var insertFlag bool
	for i, v := range q.items {
		if g.field[p] < g.field[v] {
			if i > 0 {
				q.items = append(q.items[:i+1], q.items[i:]...)
				q.items[i] = p
				insertFlag = true
			} else {
				q.items = append([]point{p}, q.items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		q.items = append(q.items, p)
	}
}

func (q *queue) lpop() point {
	p := q.items[0]
	q.items = q.items[1:len(q.items)]
	return p
}

type grid struct {
	field map[point]int
	h, w  int
}

type point struct {
	x, y int
}

func (p point) left() point {
	return point{p.x + 1, p.y}
}

func (p point) right() point {
	return point{p.x - 1, p.y}
}

func (p point) down() point {
	return point{p.x, p.y + 1}
}

func (p point) up() point {
	return point{p.x, p.y - 1}
}

func (p point) neighbors() []point {
	return []point{p.left(), p.up(), p.right(), p.down()}
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

func (g grid) findPath(start, end point) ([]point, int) {
	visited := make(map[point]struct{})
	dist := make(map[point]int)
	prev := make(map[point]point)

	q := queue{items: make([]point, 0, 100)}

	for k := range g.field {
		dist[k] = math.MaxInt64
	}
	dist[start] = 0
	q.add(start, g)
	for len(q.items) != 0 {
		v := q.lpop()

		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		near := v.neighbors()

		for _, val := range near {
			if !g.inBound(val) {
				continue
			}
			if _, ok := visited[val]; ok {
				continue
			}
			if dist[v]+g.field[val] < dist[val] {
				store := val

				dist[val] = dist[v] + g.field[val]
				g.field[val] += dist[v]
				prev[val] = v
				q.add(store, g)
			}
		}
	}
	pathval := prev[end]
	var final []point
	final = append(final, end)
	for pathval != start {
		final = append(final, pathval)
		pathval = prev[pathval]
	}
	final = append(final, pathval)
	for i, j := 0, len(final)-1; i < j; i, j = i+1, j-1 {
		final[i], final[j] = final[j], final[i]
	}
	return final, dist[end]
}

func solve(file io.Reader) int {
	g := newGrid(parse(file))
	start := point{0, 0}
	end := point{g.h - 1, g.w - 1}
	_, cost := g.findPath(start, end)
	return cost
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
	a1 := solve(file)
	fmt.Println("First answer: ", a1)
}
