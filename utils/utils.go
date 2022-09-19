package utils

import (
	"bufio"
	"os"
	"strconv"
)

// StringToInt
//
// # Convert string to int
//
// Parameters
//
//	s: string
//
// Returns
//
//	int
//	error
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToFloat32
//
// # Convert string to float32
//
// Parameters
//
//	s: string
//
// Returns
//
//	float32
//	error
func StringToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

// IsFileExist
//
// # Check file exist
//
// Parameters
//
//	path: file path
//
// Returns
//
//	bool
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// File text read and return string list (line by line)
//
// Parameters
//
//	path: file path
//
// Returns
//
//	[]string
//	error
func ReadFileText(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
