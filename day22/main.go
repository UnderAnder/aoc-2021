package main

import (
	"fmt"
	"strings"

	"aoc/pkg/helpers"
)

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

func inOrder(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

type cube struct {
	x, y, z int
}

type cuboid struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

type Instruction struct {
	cuboid cuboid
	on     bool
}

func (c *cuboid) volume() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func part1(input []string) (answer1 int) {
	steps := parse(input)
	cubesState := make(map[cube]bool)

	for _, step := range steps {
		for x := max(-50, step.cuboid.x1); x <= min(50, step.cuboid.x2); x++ {
			for y := max(-50, step.cuboid.y1); y <= min(50, step.cuboid.y2); y++ {
				for z := max(-50, step.cuboid.z1); z <= min(50, step.cuboid.z2); z++ {
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

func parse(input []string) (instructions []Instruction) {
	for _, step := range input {
		s := strings.Split(step, " ")
		str := strings.Split(s[1], ",")
		x := strings.Split(str[0][2:], "..")
		y := strings.Split(str[1][2:], "..")
		z := strings.Split(str[2][2:], "..")
		x1, x2 := inOrder(helpers.MustAtoi(x[0]), helpers.MustAtoi(x[1]))
		y1, y2 := inOrder(helpers.MustAtoi(y[0]), helpers.MustAtoi(y[1]))
		z1, z2 := inOrder(helpers.MustAtoi(z[0]), helpers.MustAtoi(z[1]))

		cuboid := cuboid{x1, x2, y1, y2, z1, z2}
		instruction := Instruction{cuboid, s[0] == "on"}

		instructions = append(instructions, instruction)
	}
	return instructions
}

func main() {
	input := "input.txt"
	a1 := part1(helpers.LoadString(input))
	fmt.Println("First answer: ", a1)
}
