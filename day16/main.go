package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength int    // length in bits.
	Pos       int
}

func (b *BitString) At(i int) byte {
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

func (b *BitString) Read(num int) uint64 {
	var bits uint64
	for i := 1; i <= num; i++ {
		bits |= uint64(b.At(b.Pos+i)) << (num - i)
	}
	b.Pos += num
	return bits
}

func parseBitString(bytes []byte) (ret BitString) {
	if len(bytes) == 0 {
		log.Println("zero length BIT STRING")
		return
	}
	ret.BitLength = (len(bytes)) * 8
	ret.Bytes = bytes
	return ret
}

type packet struct {
	version uint64
	typeID  uint64
	value   uint64
	packets []packet
}

func verSum(packet packet) uint64 {
	sum := packet.version
	for _, p := range packet.packets {
		sum += verSum(p)
	}
	return sum
}

func eval(p packet) int {
	var result int
	switch p.typeID {
	case 0:
		for i := range p.packets {
			result += eval(p.packets[i])
		}
	case 1:
		result = 1
		for i := range p.packets {
			result *= eval(p.packets[i])
		}
	case 2:
		result = math.MaxInt32
		for i := range p.packets {
			tmp := eval(p.packets[i])
			if tmp < result {
				result = tmp
			}
		}
	case 3:
		result = 0
		for i := range p.packets {
			tmp := eval(p.packets[i])
			if tmp > result {
				result = tmp
			}
		}
	case 4:
		return int(p.value)
	case 5:
		result = 0
		if len(p.packets) == 2 && p.packets[0].value > p.packets[1].value {
			result = 1
		}
	case 6:
		result = 0
		if len(p.packets) == 2 && p.packets[0].value < p.packets[1].value {
			result = 1
		}
	case 7:
		result = 0
		if len(p.packets) == 2 && p.packets[0].value == p.packets[1].value {
			result = 1
		}
	}
	return result
}

func (b *BitString) process() packet {
	packets := make([]packet, 0, 10)
	version := b.Read(3)
	typeID := b.Read(3)
	if b.Pos > b.BitLength {
		return packet{}
	}
	if typeID == 4 {
		var num uint64
		for {
			bit := b.Read(1)
			num += b.Read(4)
			if bit == 1 {
				num <<= 4
			}
			if bit == 0 {
				break
			}
		}
		return packet{
			version: version,
			typeID:  typeID,
			value:   num,
		}
	}
	length := b.Read(1)
	switch length {
	case 0:
		subLen := b.Read(15)
		exitPoint := b.Pos + int(subLen)
		for {
			pack := b.process()
			packets = append(packets, pack)
			if b.Pos == exitPoint {
				break
			}
		}
	case 1:
		subNum := b.Read(11)
		for j := 0; j < int(subNum); j++ {
			pack := b.process()
			packets = append(packets, pack)
		}
	}
	return packet{
		version: version,
		typeID:  typeID,
		packets: packets,
	}
}

func solve(file io.Reader) (answer1, answer2 int) {
	parsed := parse(file)
	bitStr := parseBitString(parsed)
	p := bitStr.process()
	answer1 = int(verSum(p))
	answer2 = eval(p)
	return answer1, answer2
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
	a1, a2 := solve(file)
	fmt.Println("First answer: ", a1)
	fmt.Println("Second answer: ", a2)
}
