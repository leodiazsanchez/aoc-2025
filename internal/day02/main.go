package main

import (
	"github.com/leodiazsanchez/aoc-2025/internal"
	"os"
	"strconv"
	"strings"
)

type solve struct{}

func main() {
	r := internal.NewRunner(&solve{})
	r.Run()
}

func (s *solve) Part1() string {
	b, _ := os.ReadFile(os.Args[1])
	str := string(b)
	intervals := strings.Split(str, ",")
	sum := 0
	for _, interval := range intervals {
		ranges := strings.Split(interval, "-")
		l, _ := strconv.Atoi(ranges[0])
		r, _ := strconv.Atoi(ranges[1])
		for i := l; i <= r; i++ {
			num := strconv.Itoa(i)
			if len(num)%2 != 0 {
				continue
			}
			half := len(num) / 2
			if strings.EqualFold(num[:half], num[half:]) {
				sum += i
			}
		}
	}
	return strconv.Itoa(sum)
}

func (s *solve) Part2() string {
	b, _ := os.ReadFile(os.Args[1])
	str := string(b)
	intervals := strings.Split(str, ",")
	sum := 0
	for _, interval := range intervals {
		ranges := strings.Split(interval, "-")
		l, _ := strconv.Atoi(ranges[0])
		r, _ := strconv.Atoi(ranges[1])
		for i := l; i <= r; i++ {
			num := strconv.Itoa(i)
		loop:
			for j := 1; j < len(num); j++ {
				curr := num[:j]
				if len(num)%j != 0 {
					continue
				}
				for _, v := range splitEveryN(num, j) {
					if !strings.EqualFold(curr, v) {
						continue loop
					}
				}
				sum += i
				break
			}
		}
	}
	return strconv.Itoa(sum)
}

func splitEveryN(s string, n int) []string {
	out := make([]string, 0)
	for i := 0; i < len(s); i += n {
		out = append(out, s[i:i+n])
	}
	return out
}
