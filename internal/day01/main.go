package main

import (
	"github.com/leodiazsanchez/aoc-2025/internal"
	"github.com/leodiazsanchez/aoc-2025/pkg/utils"
	"log"
	"os"
	"strconv"
)

type solve struct{}

func main() {
	r := internal.NewRunner(&solve{})
	r.Run()
}

func (s *solve) Part1() string {
	lines, err := utils.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	pos := 50

	for _, line := range lines {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		if dir == 'L' {
			steps = -steps
		}

		pos = ((steps + pos) + 100) % 100

		if pos == 0 {
			count++
		}
	}

	return strconv.Itoa(count)
}

func (s *solve) Part2() string {
	lines, err := utils.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	pos := 50
	count := 0

	for _, line := range lines {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		switch dir {
		case 'R':
			for range steps {
				pos = (pos + 1) % 100
				if pos == 0 {
					count++
				}
			}
		case 'L':
			for range steps {
				pos = (pos - 1 + 100) % 100
				if pos == 0 {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count)
}
