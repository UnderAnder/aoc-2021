package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf1 := &bytes.Buffer{}
	buf2 := io.TeeReader(file, buf1)
	fmt.Printf("First answer: %d\n", part1(buf2))
	fmt.Printf("Second answer: %d\n", part2(buf1))
}

func part1(file io.Reader) int {
	var position, depth int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		command := str[0]
		units, err := strconv.Atoi(str[1])
		if err != nil {
			log.Fatal(err)
		}
		switch command {
		case "forward":
			position += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	answer := position * depth
	return answer
}

func part2(file io.Reader) int {
	var position, depth, aim int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		command := str[0]
		units, err := strconv.Atoi(str[1])
		if err != nil {
			log.Fatal(err)
		}
		switch command {
		case "forward":
			position += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	answer := position * depth
	return answer
}
