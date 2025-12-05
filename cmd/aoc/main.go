package main

import (
	"fmt"
	"github.com/leodiazsanchez/aoc-2025/internal"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

var dayRegex = regexp.MustCompile(`^\d{1,2}$`)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected arguments: <day>")
	}

	day := os.Args[1]

	if !dayRegex.MatchString(day) {
		log.Fatalf("invalid day '%s'. Day must be a 1 or 2 digit number.\n", day)
	}

	if len(os.Args) > 2 && os.Args[2] == "-get" {
		d, err := strconv.Atoi(day)
		if err != nil {
			log.Fatalf("invalid day number: %v", err)
		}

		if err = internal.GetPuzzleInput(d); err != nil {
			log.Fatalf("failed to get puzzle input for day %s: %v", day, err)
		}

		fmt.Printf("Successfully downloaded input for day %s\n", day)
		return
	}

	programPath := filepath.Join("internal", fmt.Sprintf("day%02s", day), "main.go")
	inputPath := filepath.Join("inputs", fmt.Sprintf("day%02s.txt", day))

	if _, err := os.Stat(programPath); os.IsNotExist(err) {
		log.Fatalf("solution not found day %s.\nExpected source file: %s\n", day, programPath)
	}

	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		log.Fatalf("input file not found for day %s.\nExpected input file: %s", day, inputPath)
	}

	cmd := exec.Command("go", "run", programPath, inputPath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("failed to run Day %s!!!\n--- Output ---\n%s", day, output)
	}

	fmt.Printf("--- Day %s ---\n", day)
	fmt.Println(string(output))
}
