package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength int    // length in bits.
}

func (b BitString) At(i int) byte {
	if i < 0 || i >= b.BitLength {
		return 0
	}
	x := (i - 1) / 8
	y := 8 - uint(i%8)

	if y == 8 {
		return b.Bytes[x] & 1
	}
	return b.Bytes[x] >> y & 1
}

func parseBitString(bytes []byte) (ret BitString, err error) {
	if len(bytes) == 0 {
		log.Println("zero length BIT STRING")
		return
	}
	ret.BitLength = (len(bytes)) * 8
	ret.Bytes = bytes
	return
}

func solve(file io.Reader) (answer1 int) {
	p := parse(file)
	bitStr, err := parseBitString(p)
	if err != nil {
		log.Println(err)
	}
	log.Println(bitStr)

	log.Printf("%T, %08b\n", bitStr.Bytes, bitStr.Bytes)
	answer1 = process(bitStr, 1)

	return answer1
}

func process(bitStr BitString, i int) int {
	var versionSum int
packet:
	for i <= bitStr.BitLength {
		log.Println("i: ", i)
		var version byte
		end := i + 2
		for ; i <= end; i++ {
			bit := bitStr.At(i)
			// log.Println(bit)
			version |= (bit << byte(end-i))
		}
		log.Println("ver: ", version)
		versionSum += int(version)
		var typeID byte
		end = i + 2
		for ; i <= end; i++ {
			bit := bitStr.At(i)
			typeID |= (bit << byte(end-i))
		}
		log.Println("typeID: ", typeID)

		if typeID == 4 {
			var num uint64
			for ; i <= bitStr.BitLength; i += 5 {
				bit := bitStr.At(i)
				for j := 1; j <= 4; j++ {
					b := bitStr.At(i + j)
					log.Println("b: ", b, i+j)
					num |= uint64(b) << uint64(4-j)
				}
				if bit == 1 {
					num <<= 4
				}
				if bit == 0 {
					i += 5
					log.Println("num: ", num)
					log.Printf("%T, %64b\n", num, num)
					continue packet
				}
			}
		}
		length := bitStr.At(i)
		i++
		log.Println("length: ", length)
		var subLen byte
		var subNum byte
		switch length {
		case 0:
			end = i + 14
			for ; i <= end; i++ {
				bit := bitStr.At(i)
				subLen |= (bit << byte(end-i))
			}
			continue packet
		case 1:
			end = i + 10
			for ; i <= end; i++ {
				bit := bitStr.At(i)
				subNum |= (bit << byte(end-i))
			}
			continue packet
		}
	}
	return versionSum
}

func parse(file io.Reader) []byte {
	p := make([]byte, 0, 100)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()
	for i := 0; i < len(s)-1; i += 2 {
		v, err := strconv.ParseUint(s[i:i+2], 16, 8)
		if err != nil {
			log.Fatal("can't parse input ", err)
		}
		log.Println(s[i : i+2])
		p = append(p, byte(v))
	}

	return p
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	a1 := solve(file)
	fmt.Println("First answer: ", a1)
}
