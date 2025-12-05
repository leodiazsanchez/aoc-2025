# Advent of Code 2025

Solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Setup
- **Clone the repository**
   ```bash
   git clone <repository-url>
   cd aoc-2025
   ```
## How to Run

Run a specific day's solution:
```bash
make run day=<day-number>
```

Get a specific day's input:
```bash
make get day=<day-number>
```

## Repository Structure

The repo follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

```
aoc-2025/
├── cmd/
│   └── aoc/
│       └── main.go          # Runs a given solution
├── internal/
│   ├── day.go               # Day interface and runner
│   └── dayXX/
│       └── main.go          # Day XX solution
├── inputs/
│   └── dayXX.txt            # Day XX input file
├── pkg/
│   └── utils/               # Utilities for solutions
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
