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

type positions []int

func parse(file io.Reader) positions {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	init := strings.Split(scanner.Text(), ",")
	p := sliceAtoi(init)

	return p
}

func max(s []int) int {
	var result int
	for _, v := range s {
		if v > result {
			result = v
		}
	}
	return result
}

func sliceAtoi(s []string) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		result = append(result, mustAtoi(v))
	}
	return result
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int from input")
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solve(p positions) (int, int) {
	maxPos := max(p)
	return part1(p, maxPos), part2(p, maxPos)
}

func part1(p positions, maxPos int) int {
	cnt := make(map[int]int)
	for i := 0; i <= maxPos; i++ {
		for _, v := range p {
			cnt[i] += abs(i - v)
		}
	}
	return answer(cnt)
}

func part2(p positions, maxPos int) int {
	cnt := make(map[int]int)
	for i := 0; i <= maxPos; i++ {
		for _, v := range p {
			cnt[i] += (abs(i-v) + 1) * abs(i-v) / 2
		}
	}
	return answer(cnt)
}

func answer(cnt map[int]int) int {
	result := 1<<(strconv.IntSize-1) - 1
	for _, v := range cnt {
		if v < result {
			result = v
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	p := parse(file)
	a1, a2 := solve(p)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
