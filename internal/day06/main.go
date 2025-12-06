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
	cols := make([][]string, len(strings.Fields(lines[0])))
	for _, line := range lines {
		fields := strings.Fields(line)
		for i, field := range fields {
			cols[i] = append(cols[i], field)
		}
	}

	sum := 0
	for _, col := range cols {
		operator := col[len(col)-1]
		tmp := 0
		for _, num := range col[:len(col)-1] {
			n, _ := strconv.Atoi(num)
			switch operator {
			case "+":
				tmp += n
			case "*":
				if tmp == 0 {
					tmp = n
					continue
				}
				tmp *= n
			}
		}
		sum += tmp
	}
	return strconv.Itoa(sum)
}

func (s *solve) Part2() string {
	lines, _ := utils.ReadLines(os.Args[1])
	cols := make([]string, len(lines[0]))
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if c >= len(cols) {
				cols = append(cols, make([]string, c-len(cols)+1)...)
			}
			cols[c] += string(lines[r][c])
		}
	}

	sum := 0
	tmp := 0
	operators := "+*"
	var operator string
	for _, col := range cols {
		last := col[len(col)-1:]
		if strings.Contains(operators, last) {
			operator = last
			col = col[:len(col)-1]
		}

		if col = strings.TrimSpace(col); col == "" {
			sum += tmp
			tmp = 0
			continue
		}

		n, _ := strconv.Atoi(col)
		switch operator {
		case "+":
			tmp += n
		case "*":
			if tmp == 0 {
				tmp = n
				continue
			}
			tmp *= n
		}
	}

	sum += tmp

	return strconv.Itoa(sum)
}
