package main

import (
	"github.com/leodiazsanchez/aoc-2025/internal"
	"github.com/leodiazsanchez/aoc-2025/pkg/utils"
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
	lines, _ := utils.ReadLines(os.Args[1])
	return d4(lines, false)
}

func (s *solve) Part2() string {
	lines, _ := utils.ReadLines(os.Args[1])
	return d4(lines, true)
}

func d4(lines []string, isP2 bool) string {
	sum := 0
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == '.' {
				continue
			}
			adj := strings.Builder{}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if b, err := utils.GetMatrixByteAt(lines, r+1*i, c+1*j); err == nil && i == 0 && j == 0 {
						adj.WriteByte(b)
					}
				}
			}
			if strings.Count(adj.String(), "@") < 4 {
				sum++
				if isP2 {
					b := []byte(lines[r])
					b[c] = '.'
					lines[r] = string(b)
					r = 0
				}
			}
		}
	}
	return strconv.Itoa(sum)
}
