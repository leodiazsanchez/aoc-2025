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

		if pos = (pos + steps + 100) % 100; pos == 0 {
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

		step := 1
		if dir == 'L' {
			step = -1
		}

		for range steps {
			if pos = (pos + step + 100) % 100; pos == 0 {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}
