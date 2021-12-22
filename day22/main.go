package main

import (
	"errors"
	"fmt"
	"strings"

	"aoc/pkg/helpers"
)

func inOrder(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type cube struct {
	x, y, z int
}

type cuboid struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
	on     bool
}

func (c *cuboid) volume() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func (c *cuboid) intersection(c2 cuboid) (*cuboid, error) {
	x1, x2 := max(c.x1, c2.x1), min(c.x2, c2.x2)
	y1, y2 := max(c.y1, c2.y1), min(c.y2, c2.y2)
	z1, z2 := max(c.z1, c2.z1), min(c.z2, c2.z2)
	if x2 < x1 || y2 < y1 || z2 < z1 {
		return nil, errors.New("intersection not found")
	}
	intersection := cuboid{x1, x2, y1, y2, z1, z2, false}
	return &intersection, nil
}

func getIntersections(c cuboid, cuboids []cuboid) []cuboid {
	intersections := make([]cuboid, 0, len(cuboids)+1)
	if c.on {
		intersections = append(intersections, c)
	}

	for _, cuboid := range cuboids {
		if intersection, err := c.intersection(cuboid); err == nil {
			intersection.on = !cuboid.on
			intersections = append(intersections, *intersection)
		}
	}

	return intersections
}

func count(cuboids []cuboid) (cnt int) {
	for _, c := range cuboids {
		if c.on {
			cnt += c.volume()
		} else {
			cnt -= c.volume()
		}
	}
	return cnt
}

func part2(input []string) int {
	cuboids := parse(input)
	all := make([]cuboid, 0)

	for _, c := range cuboids {
		intersections := getIntersections(c, all)
		all = append(all, intersections...)
	}

	cubesOn := count(all)

	return cubesOn
}

func part1(input []string) (answer1 int) {
	steps := parse(input)
	cubesState := make(map[cube]bool)

	for _, step := range steps {
		for x := max(-50, step.x1); x <= min(50, step.x2); x++ {
			for y := max(-50, step.y1); y <= min(50, step.y2); y++ {
				for z := max(-50, step.z1); z <= min(50, step.z2); z++ {
					cubesState[cube{x, y, z}] = step.on
				}
			}
		}
	}

	for cube, isOn := range cubesState {
		if !isOn {
			delete(cubesState, cube)
		}
	}

	return len(cubesState)
}

func parse(input []string) (cuboids []cuboid) {
	for _, step := range input {
		s := strings.Split(step, " ")
		str := strings.Split(s[1], ",")
		x := strings.Split(str[0][2:], "..")
		y := strings.Split(str[1][2:], "..")
		z := strings.Split(str[2][2:], "..")
		x1, x2 := inOrder(helpers.MustAtoi(x[0]), helpers.MustAtoi(x[1]))
		y1, y2 := inOrder(helpers.MustAtoi(y[0]), helpers.MustAtoi(y[1]))
		z1, z2 := inOrder(helpers.MustAtoi(z[0]), helpers.MustAtoi(z[1]))

		cuboids = append(cuboids, cuboid{x1, x2, y1, y2, z1, z2, s[0] == "on"})
	}

	return cuboids
}

func main() {
	input := "input.txt"
	a1 := part1(helpers.LoadString(input))
	a2 := part2(helpers.LoadString(input))
	fmt.Println("First answer: ", a1)
	fmt.Println("Second answer: ", a2)
}
