package main

// ParseLine parses the line
func ParseLine(line []byte) []Token {
	tokens := []Token{}

	for index := 0; index < len(line); index++ {
		currentByte := line[index]

		// Ignore whitespace
		if currentByte == 0 || Matches(currentByte, " ") || Matches(currentByte, "\n") {
			continue
		}

		if Matches(currentByte, "l") && Matches(line[index+1], "e") && Matches(line[index+2], "t") {
			tokens = append(tokens, Token{token: "let", tokenType: Variable})
			index += 2
		} else if IsLetter(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = append(tokens, Token{token: token, tokenType: Identifier})
		} else if Matches(currentByte, "=") {
			tokens = append(tokens, Token{token: "=", tokenType: Assignment})
		} else if IsNumber(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = append(tokens, Token{token: token, tokenType: Number})
		}
	}

	return tokens
}
