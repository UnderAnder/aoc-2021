package main

import (
	"fmt"
	"strings"

	"aoc/pkg/helpers"
)

const (
	empty = iota
	east
	south
)

type point struct {
	x, y int
}

type grid struct {
	field map[point]int
	h, w  int
}

func (g grid) move() int {
	buf := make([]point, 0, 100)
	buf2 := make([]point, 0, 100)
	for p, v := range g.field {
		if v == east && g.field[p.right(g)] == empty {
			buf = append(buf, p)
		}
	}
	for _, p := range buf {
		g.field[p], g.field[p.right(g)] = g.field[p.right(g)], g.field[p]
	}
	for p, v := range g.field {
		if v == south && g.field[p.down(g)] == empty {
			buf2 = append(buf2, p)
		}
	}

	for _, p := range buf2 {
		g.field[p], g.field[p.down(g)] = g.field[p.down(g)], g.field[p]
	}
	return len(buf) + len(buf2)
}

func (g grid) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for i := 0; i < g.h; i++ {
		for j := 0; j < g.w; j++ {
			v, ok := g.field[point{j, i}]
			if !ok {
				sb.WriteString(".")
				continue
			}
			switch v {
			case east:
				sb.WriteString(">")
			case south:
				sb.WriteString("v")
			case empty:
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p point) right(g grid) point {
	if p.x >= g.w-1 {
		return point{0, p.y}
	}
	return point{p.x + 1, p.y}
}

func (p point) down(g grid) point {
	if p.y >= g.h-1 {
		return point{p.x, 0}
	}
	return point{p.x, p.y + 1}
}

func part1(input []string) (answer1 int) {
	g := parse(input)

	for {
		moves := g.move()
		answer1++
		if moves == 0 {
			break
		}
	}

	return answer1
}

func parse(input []string) grid {
	g := grid{field: make(map[point]int, len(input)*len(input[0]))}
	g.h, g.w = len(input), len(input[0])
	for i, str := range input {
		for j, char := range str {
			p := point{j, i}
			switch char {
			case '.':
				g.field[p] = empty
			case '>':
				g.field[p] = east
			case 'v':
				g.field[p] = south
			}
		}
	}

	return g
}

func main() {
	input := "input.txt"
	a1 := part1(helpers.LoadString(input))
	fmt.Println("First answer: ", a1)
}
