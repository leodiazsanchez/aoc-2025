package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
