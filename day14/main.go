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

func solve(file io.Reader) (answer1, answer2 int) {
	const (
		part1Steps = 10
		part2Steps = 40
	)
	uniqPairs := make(map[string]int)

	template, rules := parse(file)

	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		in := rules[pair]
		uniqPairs[string([]uint8{pair[0], in[0]})]++
		uniqPairs[string([]uint8{in[0], pair[1]})]++
	}
	for j := 0; j < part2Steps-1; j++ {
		if j == part1Steps-1 {
			answer1 = answer(uniqPairs)
		}
		stepPairs := make(map[string]int)
		for pair := range uniqPairs {
			in := rules[pair]
			stepPairs[string([]uint8{pair[0], in[0]})] += uniqPairs[pair]
			stepPairs[string([]uint8{in[0], pair[1]})] += uniqPairs[pair]
		}
		uniqPairs = stepPairs
	}
	answer2 = answer(uniqPairs)
	return answer1, answer2
}

func answer(m map[string]int) int {
	uniqChars := make(map[uint8]int)
	for k, v := range m {
		uniqChars[k[1]] += v
	}
	min, max := 1<<(strconv.IntSize-1)-1, 0
	for _, v := range uniqChars {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func parse(file io.Reader) (string, map[string]string) {
	rules := make(map[string]string)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		r := strings.Split(scanner.Text(), " -> ")
		rules[r[0]] = r[1]
	}
	return template, rules
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
