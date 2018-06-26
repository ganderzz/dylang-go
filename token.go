package main

type tokenType int

const (
	// Variable token
	Variable tokenType = 0
	// Identifier token
	Identifier tokenType = 1
	// Assignment token
	Assignment tokenType = 2
	// Number token
	Number tokenType = 3
)

// Token token
type Token struct {
	token     string
	tokenType tokenType
}
