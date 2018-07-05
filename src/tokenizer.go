package main

// Tokenize parses the line
func Tokenize(line []string) TokenList {
	tokens := TokenList{}

	for index := 0; index < len(line); index++ {
		currentByte := line[index]

		// Ignore whitespace
		if Matches(currentByte, " ") {
			continue
		}

		word, wordCount := GetToWhiteSpace(line, index)

		if word == "let" {
			tokens = tokens.Add(Token{value: "let", tokenType: Variable})
			index += wordCount
		} else if Matches(currentByte, "=") {
			tokens = tokens.Add(Token{value: "=", tokenType: Assignment})
		} else if IsNumber(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.Add(Token{value: token, tokenType: Number})
		} else if Matches(currentByte, "(") {
			tokens = tokens.Add(Token{value: currentByte, tokenType: LeftParen})
		} else if Matches(currentByte, ")") {
			tokens = tokens.Add(Token{value: currentByte, tokenType: RightParen})
		} else if Matches(currentByte, ">") {
			tokens = tokens.Add(Token{value: currentByte, tokenType: LessThan})
		} else if Matches(currentByte, "{") {
			tokens = tokens.Add(Token{value: currentByte, tokenType: OpenBrace})
		} else if Matches(currentByte, "}") {
			tokens = tokens.Add(Token{value: currentByte, tokenType: CloseBrace})
		} else {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.Add(Token{value: token, tokenType: Identifier})
		}
	}

	tokens = tokens.Add(Token{value: "", tokenType: End})

	return tokens
}
