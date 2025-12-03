package main

import (
	"github.com/leodiazsanchez/aoc-2025/internal"
	"github.com/leodiazsanchez/aoc-2025/pkg/utils"
	"os"
	"strconv"
)

type solve struct{}

func main() {
	r := internal.NewRunner(&solve{})
	r.Run()
}

func (s *solve) Part1() string {
	lines, _ := utils.ReadLines(os.Args[1])
	return strconv.Itoa(d3(lines, 2))
}

func (s *solve) Part2() string {
	lines, _ := utils.ReadLines(os.Args[1])
	return strconv.Itoa(d3(lines, 12))
}

func d3(lines []string, size int) int {
	sum := 0
	for _, line := range lines {
		out := ""
		startIdx := 0
		minTrailingDigits := size - 1
		for range size {
			n, idx := findFirstMaxNum(line, startIdx, minTrailingDigits)
			out += strconv.Itoa(n)
			startIdx = idx + 1
			minTrailingDigits--
		}
		s, _ := strconv.Atoi(out)
		sum += s
	}

	return sum
}

func findFirstMaxNum(line string, startIdx, minTrailingDigits int) (int, int) {
	m := 0
	idx := 0
	for i := startIdx; i < len(line); i++ {
		c, _ := strconv.Atoi(string(line[i]))
		valid := len(line)-i > minTrailingDigits
		if c > m && valid {
			m = c
			idx = i
		}
	}

	return m, idx
}
