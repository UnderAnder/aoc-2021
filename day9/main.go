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

type points [][]int

func parse(file io.Reader) points {
	p := make(points, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := sliceAtoi(strings.Split(scanner.Text(), ""))
		p = append(p, s)
	}
	return p
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

func solve(ps points) int {
	var answer int
	lps := make([]int, 0)
	for i := range ps {
		for j := range ps[i] {
			if lowPoint(i, j, ps) {
				lps = append(lps, ps[i][j])
			}
		}
	}
	for _, v := range lps {
		answer += v + 1
	}
	return answer
}

func lowPoint(i, j int, ps points) bool {
	result := true
	pnt := ps[i][j]
	if j+1 < len(ps[i]) {
		result = result && ps[i][j+1] > pnt
	}
	if j > 0 {
		result = result && ps[i][j-1] > pnt
	}
	if i+1 < len(ps) {
		result = result && ps[i+1][j] > pnt
	}
	if i > 0 {
		result = result && ps[i-1][j] > pnt
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
	a1 := solve(p)
	fmt.Printf("First answer: %d\n", a1)
}
