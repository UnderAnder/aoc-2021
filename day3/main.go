package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
}

func part1(file io.Reader) int {
	const asciiZero = 48
	var total int
	var gamma, epsilon uint16
	zeroCounter := make([]int, 12, 12)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total++
		for i, v := range scanner.Text() {
			if v == asciiZero {
				zeroCounter[i%len(scanner.Text())]++
			}
		}
	}
	for i, v := range zeroCounter {
		var mask uint16 = 1 << uint16(len(zeroCounter)-i-1)
		switch {
		case v > total/2:
			epsilon = (epsilon & ^mask) | (1 << uint16(len(zeroCounter)-i-1))
		case v < total/2:
			gamma = (gamma & ^mask) | (1 << uint16(len(zeroCounter)-i-1))
		default:
			log.Fatal("no most common bit in the corresponding position")
		}
	}
	log.Printf("%.16b", gamma)
	log.Printf("%.16b", epsilon)
	answer := int(gamma) * int(epsilon)
	return answer
}
