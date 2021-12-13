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

type fold struct {
	x, y int
}

func (g *grid) String() {
	for i := 0; i <= g.h; i++ {
		for j := 0; j <= g.w; j++ {
			if _, ok := g.field[point{j, i}]; !ok {
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func (g *grid) fold(fold fold) {
	for k := range g.field {
		if fold.x != -1 {
			if k.x > fold.x {
				px := (fold.x - (k.x % fold.x)) % fold.x
				g.field[point{px, k.y}] = struct{}{}
				delete(g.field, k)
			}
		} else {
			if k.y > fold.y {
				py := (fold.y - (k.y % fold.y)) % fold.y
				g.field[point{k.x, py}] = struct{}{}
				delete(g.field, k)
			}
		}
	}
	if fold.x != -1 {
		g.w -= fold.x
	} else {
		g.h -= fold.y
	}
}

func solve(file io.Reader) {
	folds, g := parse(file)
	for i, f := range folds {
		if i == 1 {
			fmt.Printf("First answer: %d\n", len(g.field))
		}
		g.fold(f)
	}
	fmt.Println("Second answer:")
	g.String()
}

func parse(file io.Reader) ([]fold, grid) {
	folds := make([]fold, 0, 20)
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
				folds = append(folds, fold{helpers.MustAtoi(s[1]), -1})

				continue
			}
			folds = append(folds, fold{-1, helpers.MustAtoi(s[1])})
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
	return folds, g
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	solve(file)
}
