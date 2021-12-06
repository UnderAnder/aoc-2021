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

type fishes []int

func parse(file io.Reader) fishes {
	fishesInit := make(fishes, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	init := strings.Split(scanner.Text(), ",")
	for _, s := range init {
		fishesInit = append(fishesInit, mustAtoi(s))
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

func part1(f fishes) int {
	const (
		days       = 80
		firstCycle = 8
	)

	for i := 0; i < days; i++ {
		for j := range f {
			if f[j] == 0 {
				f[j] = 6
				f = append(f, firstCycle)
				continue
			}
			f[j]--
		}
	}
	return len(f)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f := parse(file)

	fmt.Printf("First answer: %d\n", part1(f))
	// fmt.Printf("Second answer: %d\n", part2(fishes))
}
