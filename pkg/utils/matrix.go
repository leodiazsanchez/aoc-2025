package utils

import "fmt"

func GetMatrixByteAt(rows []string, r, c int) (byte, error) {
	rowCount := len(rows)
	if r >= 0 && r < rowCount && c >= 0 && c < len(rows[r]) {
		return rows[r][c], nil
	}
	return 0, fmt.Errorf("out of bounds")
}
