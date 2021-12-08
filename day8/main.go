package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// num of segments
	zero  = 6
	one   = 2
	two   = 5
	three = 5
	four  = 4
	five  = 5
	six   = 6
	seven = 3
	eight = 7
	nine  = 6
)

func parse(file io.Reader) ([][]string, [][]string) {
	ins := make([][]string, 0)
	outs := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		init := strings.Split(scanner.Text(), " | ")
		in := strings.Split(init[0], " ")
		out := strings.Split(init[1], " ")
		ins = append(ins, in)
		outs = append(outs, out)
	}

	return ins, outs
}

func uniqSegments(v int) bool {
	if v == one || v == four || v == seven || v == eight {
		return true
	}
	return false
}

func part1(outs [][]string) int {
	cnt := make(map[int]int)
	for _, out := range outs {
		for _, v := range out {
			if uniqSegments(len(v)) {
				cnt[len(v)]++
			}
		}
	}
	return answer(cnt)
}

func answer(cnt map[int]int) int {
	var result int
	for _, v := range cnt {
		result += v
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, outs := parse(file)
	a1 := part1(outs)
	fmt.Printf("First answer: %d\n", a1)
}
