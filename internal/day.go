package internal

import (
	"fmt"
)

type Day interface {
	Part1() string
	Part2() string
}

type Runner struct {
	Day Day
}

func NewRunner(day Day) *Runner {
	return &Runner{Day: day}
}

func (r *Runner) Run() {
	fmt.Printf("Part 1: %s\n", r.Day.Part1())
	fmt.Printf("Part 2: %s\n", r.Day.Part2())
}
