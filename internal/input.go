package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func GetPuzzleInput(day int) error {
	cookie := os.Getenv("SESSION")
	if cookie == "" {
		return fmt.Errorf("SESSION environment variable not set")
	}

	client := http.Client{Timeout: time.Second * 5}
	u := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	input, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("inputs/day%02d.txt", day), input, 0644)
	if err != nil {
		return err
	}

	return nil
}
