package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type entry struct {
	p1, p2 point
}

type field map[point]int

func parse(file io.Reader) []entry {
	entries := make([]entry, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		point1 := strings.Split(points[0], ",")
		p1 := point{
			x: mustAtoi(point1[0]),
			y: mustAtoi(point1[1]),
		}
		point2 := strings.Split(points[1], ",")
		p2 := point{
			x: mustAtoi(point2[0]),
			y: mustAtoi(point2[1]),
		}
		entries = append(entries, entry{p1: p1, p2: p2})
	}
	return entries
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int from input")
	}
	return i
}

func printTestField(f field, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if v, ok := f[point{x: j, y: i}]; ok {
				fmt.Print(v)
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func move(a, b int) int {
	switch {
	case a < b:
		return a + 1
	case a > b:
		return a - 1
	default:
		return a
	}
}

func answer(f field) (answer int) {
	for _, i := range f {
		if i > 1 {
			answer++
		}
	}
	return answer
}

func part1(entries []entry) int {
	f := make(field)

	for _, e := range entries {
		start, end := e.p1, e.p2
		i := start
		for i != end {
			if i.x == end.x || i.y == end.y {
				f[i]++
			}
			i = point{move(i.x, end.x), move(i.y, end.y)}
		}
		if start.x == end.x || start.y == end.y {
			f[i]++
		}
	}

	return answer(f)
}

func part2(entries []entry) int {
	f := make(field)

	for _, e := range entries {
		start, end := e.p1, e.p2
		for start != end {
			f[start]++
			start = point{move(start.x, end.x), move(start.y, end.y)}
		}
		f[start]++
	}
	return answer(f)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	entries := parse(file)

	fmt.Printf("First answer: %d\n", part1(entries))
	fmt.Printf("Second answer: %d\n", part2(entries))
}
