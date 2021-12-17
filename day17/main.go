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

func search(dx, dy int, target map[string]int) (int, bool) {
	var x, y, maxHeight int
	for {
		x += dx
		y += dy
		if y > maxHeight {
			maxHeight = y
		}
		if dx < 0 {
			dx++
		} else if dx > 0 {
			dx--
		}
		dy--
		if x >= target["xMin"] && x <= target["xMax"] && y >= target["yMin"] && y <= target["yMax"] {
			return maxHeight, true
		}

		if y < target["yMin"] || (dx == 0 && (x > target["xMax"] || x < target["xMin"])) {
			return maxHeight, false
		}
	}
}

func solve(file io.Reader) (answer1 int) {
	target := parse(file)
	log.Println(target)
	maxHeight := 0
	for i := 0; i < 300; i++ {
		for j := target["yMin"]; j < 300; j++ {
			h, ok := search(i, j, target)
			if ok && h > maxHeight {
				maxHeight = h
			}
		}
	}
	answer1 = maxHeight
	return answer1
}

func parse(file io.Reader) map[string]int {
	target := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ": ")
	s = strings.Split(s[1], ", ")
	x := strings.Split(s[0], "..")
	y := strings.Split(s[1], "..")
	target["xMin"] = helpers.MustAtoi(x[0][2:])
	target["xMax"] = helpers.MustAtoi(x[1])
	target["yMin"] = helpers.MustAtoi(y[0][2:])
	target["yMax"] = helpers.MustAtoi(y[1])
	return target
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
