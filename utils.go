package main

import (
	"strconv"
	"unicode"
)

// Matches checks if a byte equals an array
func Matches(item byte, expected string) bool {
	return string(item) == expected
}

// IsLetter checks if byte is a valid unicode character
func IsLetter(b byte) bool {
	return unicode.IsLetter(rune(b))
}

// IsNumber checks if byte is a number
func IsNumber(b byte) bool {
	if _, err := strconv.ParseInt(string(b), 10, 64); err == nil {
		return true
	}

	return false
}

// GetToWhiteSpace grabs the current string until whitespace
func GetToWhiteSpace(line []byte, index int) (response string, count int) {
	start := line[index:]

	for _, item := range start {
		val := string(item)

		if val == " " {
			break
		}
		response += val
		count++
	}

	return response, count
}
