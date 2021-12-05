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

const size = 5

type point struct {
	x, y int
}

type entry struct {
	p1, p2 point
}

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

func printField(field map[point]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if v, ok := field[point{x: j, y: i}]; ok {
				fmt.Print(v)
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func part1(entries []entry) int {
	var answer int
	field := make(map[point]int)

	for _, e := range entries {
		var start, end int
		if e.p1.x == e.p2.x {
			start, end = e.p1.y, e.p2.y
			if e.p1.y > e.p2.y {
				start, end = end, start
			}
			for i := start; i <= end; i++ {
				field[point{e.p2.x, i}]++
			}
		}
		if e.p1.y == e.p2.y {
			start, end = e.p1.x, e.p2.x
			if e.p1.x > e.p2.x {
				start, end = end, start
			}
			for i := start; i <= end; i++ {
				field[point{i, e.p2.y}]++
			}
		}
	}
	for _, i := range field {
		if i > 1 {
			answer++
		}
	}
	return answer
}

func part2(entries []entry) int {
	var answer int
	field := make(map[point]int)

	for _, e := range entries {
		var start, end int
		if e.p1.x == e.p2.x {
			start = e.p1.y
			end = e.p2.y
			if e.p1.y > e.p2.y {
				start = e.p2.y
				end = e.p1.y
			}
			for i := start; i <= end; i++ {
				field[point{e.p2.x, i}]++
			}
			continue
		}
		if e.p1.y == e.p2.y {
			start = e.p1.x
			end = e.p2.x
			if e.p1.x > e.p2.x {
				start = e.p2.x
				end = e.p1.x
			}
			for i := start; i <= end; i++ {
				field[point{i, e.p2.y}]++
			}
			continue
		}
		if e.p2.y == e.p2.x && e.p1.y == e.p1.x {
			start, end = e.p1.x, e.p2.x
			if e.p1.x > e.p2.x {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				field[point{i, i}]++
			}
			continue
		}
		if e.p1.x == e.p2.y && e.p1.y == e.p2.x {
			start, end = e.p1.x, e.p2.x
			if e.p1.x > e.p2.x {
				start, end = end, start
			}
			j := end
			for i := start; i <= end; i++ {
				field[point{j, i}]++
				j--
			}
		}
		if e.p1.x+e.p2.y == e.p2.x+e.p1.y {
			start, end = e.p1.x, e.p2.x
			start2, end2 := e.p1.y, e.p2.y
			if e.p1.x > e.p2.x {
				start, end = end, start
			}
			if e.p1.y > e.p2.y {
				start2, end2 = end2, start2
			}
			j := start2
			for i := start; i <= end; i++ {
				log.Printf("x:%d y:%d", i, j)
				field[point{i, j}]++
				j++
			}
		}
		if e.p1.x+e.p1.y == e.p2.x+e.p2.y {
			start, end = e.p1.x, e.p2.x
			start2, end2 := e.p1.y, e.p2.y
			if e.p1.x > e.p2.x {
				start, end = end, start
			}
			if e.p1.y > e.p2.y {
				start2, end2 = end2, start2
			}
			j := end2
			for i := start; i <= end; i++ {
				field[point{i, j}]++
				j--
			}
		}
	}
	for _, i := range field {
		if i > 1 {
			answer++
		}
	}
	return answer
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
