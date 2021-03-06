package helpers

import (
	"log"
	"math"
	"strconv"
	"unicode"
)

// SliceAtoi convert slice of strings to slice of integers
func SliceAtoi(s []string) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		result = append(result, MustAtoi(v))
	}
	return result
}

// MustAtoi convert string to integer
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int")
	}
	return i
}

// PowInt returns x**y
func PowInt(x, y int) int {
	if y == 0 {
		return 1
	}
	result := x
	for i := 2; i <= y; i++ {
		result *= x
	}
	return result
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Sum(s []int) int {
	res := 0
	for i := range s {
		res += s[i]
	}
	return res
}

func Mul(s []int) int {
	res := 1
	for i := range s {
		res *= s[i]
	}
	return res
}

func Min(s []int) int {
	res := math.MaxInt64
	for i := range s {
		if s[i] < res {
			res = s[i]
		}
	}
	return res
}

func Max(s []int) int {
	res := 0
	for i := range s {
		if s[i] > res {
			res = s[i]
		}
	}
	return res
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Intersection(i1, i2 []int) []int {
	result := make([]int, 0, len(i1))
	hash := make(map[int]bool)
	for _, e := range i1 {
		hash[e] = false
	}
	for _, e := range i2 {
		if v, ok := hash[e]; ok && !v {
			hash[e] = true
			result = append(result, e)

		}
	}

	return result
}

func Counter(s []int) map[int]int {
	result := make(map[int]int)
	for i := range s {
		result[s[i]]++
	}
	return result
}
