package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(first(file))
}

func first(file io.Reader) int {
	var measurement, temp, counter int
	var err error
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		measurement, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if measurement > temp {
			counter++
		}
		temp = measurement
	}
	counter-- // no previous measurement
	return counter
}
