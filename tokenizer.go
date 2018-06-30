package main

type tokenList []Token

func (t tokenList) add(value Token) tokenList {
	return append(t, value)
}

// Tokenize parses the line
func Tokenize(line []byte) []Token {
	tokens := tokenList{}

	for index := 0; index < len(line); index++ {
		currentByte := line[index]

		// Ignore whitespace
		if currentByte == 0 || Matches(currentByte, " ") || Matches(currentByte, "\n") {
			continue
		}

		if Matches(currentByte, "l") && Matches(line[index+1], "e") && Matches(line[index+2], "t") {
			tokens = tokens.add(Token{value: "let", tokenType: Variable})
			index += 2
		} else if IsLetter(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.add(Token{value: token, tokenType: Identifier})
		} else if Matches(currentByte, "=") {
			tokens = tokens.add(Token{value: "=", tokenType: Assignment})
		} else if IsNumber(currentByte) {
			token, total := GetToWhiteSpace(line, index)
			index += total

			tokens = tokens.add(Token{value: token, tokenType: Number})
		}
	}

	return tokens
}
