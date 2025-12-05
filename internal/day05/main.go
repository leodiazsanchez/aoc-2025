package main

import (
	"github.com/leodiazsanchez/aoc-2025/internal"
	"maps"
	"os"
	"slices"
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
	half := strings.Split(string(b), "\n\n")
	idRanges := strings.Split(half[0], "\n")
	ids := strings.Split(half[1], "\n")
	m := make(map[int]bool)

	for _, id := range ids {
		for _, idRange := range idRanges {
			parts := strings.Split(idRange, "-")
			l, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			x, _ := strconv.Atoi(id)
			if x >= l && x <= h {
				m[x] = true
			}
		}
	}

	return strconv.Itoa(len(m))
}

func (s *solve) Part2() string {
	b, _ := os.ReadFile(os.Args[1])
	half := strings.Split(string(b), "\n\n")
	idRanges := strings.Split(half[0], "\n")
	m := make(map[r]bool)
	sum := 0

	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		l, _ := strconv.Atoi(parts[0])
		h, _ := strconv.Atoi(parts[1])
		m[r{l, h}] = true
	}

	merge(m)

	for v := range maps.Keys(m) {
		sum += v.h + 1 - v.l
	}

	return strconv.Itoa(sum)
}

func merge(m map[r]bool) map[r]bool {
	changed := true

	for changed {
		changed = false
		keys := make([]r, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}

		for _, v := range keys {
			idx := slices.IndexFunc(keys, func(v1 r) bool {
				return v != v1 && v.l <= v1.h && v1.l <= v.h
			})

			if idx != -1 {
				v1 := keys[idx]
				delete(m, v)
				delete(m, v1)
				m[r{l: min(v.l, v1.l), h: max(v.h, v1.h)}] = true
				changed = true
				break
			}
		}
	}

	return m
}

type r struct {
	l, h int
}
