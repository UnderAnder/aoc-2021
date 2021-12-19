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

type counter map[point]int

func (c counter) add(p []point) {
	for _, v := range p {
		c[v]++
	}
}

type point struct {
	x, y, z int
}

func (p point) add(p2 point) point {
	return point{
		x: p.x + p2.x,
		y: p.y + p2.y,
		z: p.z + p2.z,
	}
}

func (p point) sub(p2 point) point {
	return point{
		x: p.x - p2.x,
		y: p.y - p2.y,
		z: p.z - p2.z,
	}
}

type scanner struct {
	points []point
	offset point
}

func (s scanner) matches(points []point) (offset point, ok bool) {
	offsetFreqs := make(map[point]int)
	for i := range s.points {
		for j := range points {
			offsetFreqs[s.points[i].sub(points[j])]++
		}
	}
	for o, freq := range offsetFreqs {
		if freq < 12 {
			continue
		}
		ok = true
		offset = o
		break
	}
	return offset, ok
}

func (s scanner) orientation(face, rotation int) []point {
	oriented := make([]point, 0, len(s.points))
	for i := range s.points {
		oriented = append(oriented, s.points[i].orientation(face, rotation))
	}
	return oriented
}

func (p point) orientation(face, rotation int) point {
	switch face {
	case 1:
	case 2:
		p.x, p.y = p.y, -p.x
	case 3:
		p.x, p.z = p.z, -p.x
	case 4:
		p.x, p.y = -p.x, -p.y
	case 5:
		p.x, p.y = -p.y, p.x
	case 6:
		p.x, p.z = -p.z, p.x
	default:
		log.Fatal("there is only 6 faces")
	}
	switch rotation {
	case 1:
	case 2:
		p.y, p.z = -p.z, p.y
	case 3:
		p.y, p.z = p.z, -p.y
	case 4:
		p.y, p.z = -p.y, -p.z
	default:
		log.Fatal("there is only 4 rotations")
	}
	return p
}

func move(s []point, offset point) []point {
	result := make([]point, 0, len(s))
	for i := range s {
		result = append(result, s[i].add(offset))
	}
	return result
}

func solve(file io.Reader) (answer1, answer2 int) {
	scanners := parse(file)
	cnt := counter{}
	solved := make([]scanner, 0, 30)
	solved = append(solved, scanners[0])
	cnt.add(solved[0].points)
	unsolved := scanners[1:]
	for len(unsolved) > 0 {
		for i := 0; i < len(unsolved); i++ {
			s, ok := brute(unsolved[i], solved)
			if !ok {
				continue
			}
			solved = append(solved, s)
			cnt.add(s.points)
			unsolved[i] = unsolved[len(unsolved)-1]
			unsolved = unsolved[:len(unsolved)-1]
			break
		}
		fmt.Println("Unsolved remain: ", len(unsolved))
	}
	max := part2(solved)

	return len(cnt), max
}

func part2(solved []scanner) int {
	max := 0
	for i := 0; i < len(solved); i++ {
		for j := 0; j < len(solved); j++ {
			if i == j {
				continue
			}
			d := distance(solved[i].offset, solved[j].offset)
			if d > max {
				max = d
			}
		}
	}
	return max
}

func brute(s scanner, solved []scanner) (scanner, bool) {
	for f := 1; f <= 6; f++ {
		for r := 1; r <= 4; r++ {
			reoriented := s.orientation(f, r)
			for j := range solved {
				offset, ok := solved[j].matches(reoriented)
				if !ok {
					continue
				}
				match := scanner{
					points: move(reoriented, offset),
					offset: offset,
				}
				return match, true
			}
		}
	}
	return scanner{}, false
}

func distance(s1, s2 point) int {
	return helpers.Abs(s1.x-s2.x) + helpers.Abs(s1.y-s2.y) + helpers.Abs(s1.z-s2.z)
}

func parse(file io.Reader) []scanner {
	scanners := make([]scanner, 0, 30)
	s := scanner{points: make([]point, 0, 26)}
	scan := bufio.NewScanner(file)
	scan.Scan()
	for scan.Scan() {
		if scan.Text() == "" {
			continue
		}
		if strings.HasPrefix(scan.Text(), "---") {
			scanners = append(scanners, s)
			s = scanner{points: make([]point, 0, 26)}
			continue
		}
		str := strings.Split(scan.Text(), ",")
		s.points = append(
			s.points, point{
				x: helpers.MustAtoi(str[0]),
				y: helpers.MustAtoi(str[1]),
				z: helpers.MustAtoi(str[2]),
			},
		)
	}
	scanners = append(scanners, s)
	return scanners
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
