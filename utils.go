package main

import (
	"strconv"
)

// Matches checks if a byte equals an array
func Matches(item string, expected string) bool {
	return item == expected
}

// IsNumber checks if byte is a number
func IsNumber(b string) bool {
	if _, err := strconv.Atoi(b); err == nil {
		return true
	}

	return false
}

// GetToWhiteSpace grabs the current string until whitespace
func GetToWhiteSpace(line []string, index int) (response string, count int) {
	start := line[index:]

	for _, item := range start {
		val := item

		if val == " " || val == "\n" {
			break
		}
		response += val
		count++
	}

	return response, count
}
