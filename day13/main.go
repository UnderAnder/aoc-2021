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
	field map[point]struct{}
	h, w  int
}

type point struct {
	x, y int
}

func (p point) plus(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

func (g *grid) inBound(p point) bool {
	if p.x >= g.w || p.x < 0 || p.y >= g.h || p.y < 0 {
		return false
	}
	return true
}

func (g *grid) String() {
	for i := 0; i <= g.h; i++ {
		for j := 0; j <= g.w; j++ {
			if _, ok := g.field[point{j, i}]; ok {
				fmt.Printf("#")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) foldY(y int) {
	for k := range g.field {
		if k.y > y {
			py := (y - (k.y % y)) % y
			g.field[point{k.x, py}] = struct{}{}
			delete(g.field, k)
		}
	}
	g.h -= y + 1
}

func (g *grid) foldX(x int) {
	for k := range g.field {
		if k.x > x {
			px := (x - (k.x % x)) % x
			g.field[point{px, k.y}] = struct{}{}
			delete(g.field, k)
		}
		g.w -= x + 1
	}
}

func solve(file io.Reader) (answer int) {
	xs, ys := make([]int, 0, 10), make([]int, 0, 10)
	mx, my := 0, 0
	g := grid{field: make(map[point]struct{})}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if strings.HasPrefix(scanner.Text(), "fold") {
			s := strings.Split(scanner.Text(), "=")
			if strings.HasSuffix(s[0], "x") {
				xs = append(xs, helpers.MustAtoi(s[1]))
				continue
			}
			ys = append(ys, helpers.MustAtoi(s[1]))
			continue
		}
		s := strings.Split(scanner.Text(), ",")
		x := helpers.MustAtoi(s[0])
		if x > mx {
			mx = x
		}
		y := helpers.MustAtoi(s[1])
		if y > my {
			my = y
		}
		p := point{x, y}
		g.field[p] = struct{}{}
	}
	g.h = my
	g.w = mx
	log.Println(g)
	log.Println(xs)
	log.Println(ys)
	g.foldX(655)
	answer = len(g.field)
	return answer
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
	fmt.Printf("First answer: %d\n", a1)
}
