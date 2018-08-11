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
	// LeftParen token
	LeftParen tokenType = 6
	// RightParen token
	RightParen tokenType = 7
	// LessThan token
	LessThan tokenType = 8
	// GreaterThan token
	GreaterThan tokenType = 9
	// OpenBrace token
	OpenBrace tokenType = 10
	// CloseBrace token
	CloseBrace tokenType = 11
	// Function token
	Function tokenType = 12
)

// Token token
type Token struct {
	value     string
	tokenType tokenType
}

// TokenList array of tokens
type TokenList []Token

// Add add token to the list
func (t TokenList) Add(value Token) TokenList {
	return append(t, value)
}

// Get get an item at a position
func (t TokenList) Get(index int) Token {
	return t[index]
}

// Find search for a token in the list
func (t TokenList) Find(item string) (match Token, position int) {
	for i := 0; i < len(t); i++ {
		if Matches(t[i].value, item) {
			return t[i], i
		}
	}

	return Token{}, -1
}
