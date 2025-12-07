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
	splits := 0
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == 'S' {
				splits = bfs(lines, r, c)
			}
		}
	}
	return strconv.Itoa(splits)
}

func (s *solve) Part2() string {
	endings := 0
	lines, _ := utils.ReadLines(os.Args[1])
	m := make(map[node]int)
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == 'S' {
				endings = endingsRec(lines, r, c, m)
			}
		}
	}
	return strconv.Itoa(endings)
}

func bfs(lines []string, r, c int) int {
	sum := 0
	q := make([]node, 0)
	visited := make(map[node]bool)
	down, err := utils.GetMatrixByteAt(lines, r+1, c)
	if err != nil {
		return sum
	}

	q = append(q, node{
		r:     r + 1,
		c:     c,
		value: down,
	})

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if visited[n] {
			continue
		}
		visited[n] = true

		switch n.value {
		case '.':
			if down, err = utils.GetMatrixByteAt(lines, n.r+1, n.c); err == nil {
				q = append(q, node{
					r:     n.r + 1,
					c:     n.c,
					value: down,
				})
			}
		case '^':
			if left, err := utils.GetMatrixByteAt(lines, n.r, n.c-1); err == nil {
				q = append(q, node{
					r:     n.r,
					c:     n.c - 1,
					value: left,
				})
			}

			if right, err := utils.GetMatrixByteAt(lines, n.r, n.c+1); err == nil {
				q = append(q, node{
					r:     n.r,
					c:     n.c + 1,
					value: right,
				})
			}

			sum++
		}
	}
	return sum
}

func endingsRec(lines []string, r, c int, m map[node]int) int {
	b, err := utils.GetMatrixByteAt(lines, r+1, c)
	if err != nil {
		return 1

	}

	n := node{
		r:     r,
		c:     c,
		value: b,
	}

	if endings, ok := m[n]; ok {
		return endings
	}

	sum := 0
	switch n.value {
	case '.':
		sum = endingsRec(lines, r+1, c, m)
	case '^':
		sum = endingsRec(lines, r, c-1, m) + endingsRec(lines, r, c+1, m)
	}

	m[n] = sum
	return sum
}

type node struct {
	r, c  int
	value byte
}
