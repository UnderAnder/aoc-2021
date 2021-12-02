package main

import (
	"bufio"
	"bytes"
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
	buf1 := &bytes.Buffer{}
	buf2 := io.TeeReader(file, buf1)
	fmt.Printf("First answer: %d\n", part1(buf2))
	fmt.Printf("Second answer: %d\n", part2(buf1))
}

func part1(file io.Reader) int {
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

func part2(file io.Reader) int {
	const threeMeasurement = 3
	var measurement, counter, sum1 int
	var err error
	temp := make([]int, threeMeasurement, threeMeasurement)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		measurement, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		temp[i%threeMeasurement] = measurement
		sum2 := sum(temp)
		if sum2 > sum1 {
			counter++
		}
		sum1 = sum2
	}
	counter -= threeMeasurement // no measurement
	return counter
}

func sum(s []int) int {
	var result int
	for _, v := range s {
		result += v
	}
	return result
}
