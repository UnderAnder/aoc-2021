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

const (
	asciiZero = 48
	asciiOne  = 49
	noCommon  = 99
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
	answer := int(gamma) * int(epsilon)
	return answer
}

func part2(file io.Reader) int {
	var generator, scrubber string
	var mComBuf []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mComBuf = append(mComBuf, scanner.Text())
	}
	lComBuf := make([]string, len(mComBuf))
	copy(lComBuf, mComBuf)

	for i := 0; i < len(mComBuf[0]); i++ {
		var tmp []string
		mCom := mostCommon(mComBuf, i)
		for _, v := range mComBuf {
			if v[i] == mCom || (v[i] == asciiOne && mCom == noCommon) {
				tmp = append(tmp, v)
			}
		}
		if len(tmp) == 1 {
			generator = tmp[0]
		}
		mComBuf = tmp
	}

	for i := 0; i < len(mComBuf[0]); i++ {
		var tmp []string
		mCom := mostCommon(lComBuf, i)
		for _, v := range lComBuf {
			if v[i] != mCom && mCom != noCommon || (v[i] == asciiZero && mCom == noCommon) {
				tmp = append(tmp, v)
			}
		}
		if len(tmp) == 1 {
			scrubber = tmp[0]
		}
		lComBuf = tmp
	}
	oxy, _ := strconv.ParseInt(generator, 2, 64)
	co2, _ := strconv.ParseInt(scrubber, 2, 64)
	answer := oxy * co2
	return int(answer)
}

func mostCommon(s []string, pos int) uint8 {
	var zeros, ones int
	for _, v := range s {
		switch v[pos] {
		case asciiZero:
			zeros++
		case asciiOne:
			ones++
		}
	}
	switch {
	case zeros > ones:
		return asciiZero
	case ones > zeros:
		return asciiOne
	default:
		return noCommon
	}
}
