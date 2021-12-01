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
	fmt.Println(second(file))
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

func second(file io.Reader) int {
	const threeMeasurement = 3
	var measurement, counter int
	var err error
	temp := make([]int, threeMeasurement, threeMeasurement)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		measurement, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum1 := sum(temp)
		temp[i%threeMeasurement] = measurement
		sum2 := sum(temp)
		if sum2 > sum1 {
			counter++
		}
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
