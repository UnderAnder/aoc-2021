package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type grid struct {
	field   map[point]bool
	h, w    int
	iter    int
	zeroLit bool
}

type point struct {
	x, y int
}

func (g grid) neighbors(p point) []bool {
	bs := make([]bool, 0, 9)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			v, ok := g.field[point{p.x + j, p.y + i}]
			// pixels are lit out of the field
			if !ok && g.zeroLit && g.iter%2 == 0 {
				bs = append(bs, true)
				continue
			}
			bs = append(bs, v)
		}
	}
	return bs
}

func (g grid) getNum(p point) int {
	var num uint16
	bs := g.neighbors(p)
	for i := 0; i < len(bs); i++ {
		if bs[i] {
			num |= 1 << (len(bs) - 1 - i)
		}
	}
	return int(num)
}

func (g grid) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for i := -g.iter; i < g.h+g.iter; i++ {
		for j := -g.iter; j < g.w+g.iter; j++ {
			v := g.field[point{j, i}]
			if !v {
				sb.WriteString(".")
				continue
			}
			sb.WriteString("#")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func newGrid(ps [][]bool) grid {
	g := grid{field: make(map[point]bool)}
	for i := range ps {
		for j := range ps[i] {
			g.field[point{j, i}] = ps[i][j]
		}
	}
	g.h, g.w = len(ps), len(ps[0])
	return g
}

func (g *grid) enhance(algo map[int]bool) {
	out := make(map[point]bool)
	g.iter++
	for i := -g.iter; i < g.h+g.iter; i++ {
		for j := -g.iter; j < g.w+g.iter; j++ {
			k := point{j, i}
			num := g.getNum(k)
			out[k] = algo[num]
		}
	}
	g.field = out
}

func solve(file io.Reader) (answer1, answer2 int) {
	algo, in := parse(file)
	inp := newGrid(in)
	inp.zeroLit = algo[0]

	for i := 0; i < 50; i++ {
		inp.enhance(algo)
		if i == 1 {
			for _, b := range inp.field {
				if b {
					answer1++
				}
			}
		}
	}
	for _, b := range inp.field {
		if b {
			answer2++
		}
	}

	return answer1, answer2
}

func parse(file io.Reader) (map[int]bool, [][]bool) {
	p := make([][]bool, 0, 100)
	algo := make(map[int]bool)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for i, v := range scanner.Text() {
		algo[i] = v == '#'
	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		s := make([]bool, 0, 100)
		for _, v := range scanner.Text() {
			s = append(s, v == '#')
		}
		p = append(p, s)
	}
	return algo, p
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	a1, a2 := solve(file)
	fmt.Println("First answer: ", a1)
	fmt.Println("Second answer: ", a2)
}
