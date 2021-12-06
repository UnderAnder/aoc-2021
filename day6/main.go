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

type fishes map[int]int

func parse(file io.Reader) fishes {
	fishesInit := make(fishes)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	init := strings.Split(scanner.Text(), ",")
	for _, s := range init {
		fishesInit[mustAtoi(s)]++
	}
	return fishesInit
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int from input")
	}
	return i
}

func calcFishes(f fishes) fishes {
	const (
		firstCycle = 8
		cycle      = 6
	)
	f2 := make(fishes)

	for i := 1; i <= firstCycle; i++ {
		if i == cycle || i == firstCycle {
			f2[i] = f[0]
		}
		f2[i-1] += f[i]
	}
	return f2
}

func calcAnswer(f fishes) int {
	var result int
	for _, v := range f {
		result += v
	}
	return result
}

func solve(f fishes) (int, int) {
	const (
		part1Days = 80
		part2Days = 256
	)
	var answer1, answer2 int

	for i := 0; i < part2Days; i++ {
		f = calcFishes(f)
		if i == part1Days-1 {
			answer1 = calcAnswer(f)
		}
	}
	answer2 = calcAnswer(f)

	return answer1, answer2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f := parse(file)
	a1, a2 := solve(f)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
