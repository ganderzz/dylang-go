package main

type tokenList []Token

func (t tokenList) add(value Token) tokenList {
	return append(t, value)
}

func (t tokenList) lookAhead(loc int) Token {
	return t[loc]
}

// Tokenize parses the line
func Tokenize(line []string) []Token {
	tokens := tokenList{}

	for index := 0; index < len(line); index++ {
		currentByte := line[index]

		// Ignore whitespace
		if Matches(currentByte, " ") {
			continue
		}

		word, wordCount := GetToWhiteSpace(line, index)

		if word == "let" {
			tokens = tokens.add(Token{value: "let", tokenType: Variable})
			index += wordCount
		} else if Matches(currentByte, "=") {
			tokens = tokens.add(Token{value: "=", tokenType: Assignment})
		} else if IsNumber(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.add(Token{value: token, tokenType: Number})
		} else {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.add(Token{value: token, tokenType: Identifier})
		}
	}

	tokens = tokens.add(Token{value: "", tokenType: End})

	return tokens
}
