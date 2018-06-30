package main

type tokenType int

const (
	// Empty nothing
	Empty tokenType = -1
	// Variable token
	Variable tokenType = 0
	// Identifier token
	Identifier tokenType = 1
	// Assignment token
	Assignment tokenType = 2
	// Number token
	Number tokenType = 3
	// End token
	End tokenType = 4
	// Program token
	Program tokenType = 5
)

// Token token
type Token struct {
	value     string
	tokenType tokenType
}
