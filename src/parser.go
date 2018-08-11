package main

// Root t
type Root struct {
	name  string
	value []*Tree
}

type tokenList = []Token

// Tree tree type
type Tree struct {
	value interface{}
	left  *Tree
	right *Tree
}

func skipTo(tokens []Token, i int, item string) int {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].value == item {
			return i
		}
	}

	return -1
}

// Parse parses a list of tokens into an ast
func Parse(tokens []Token) *Root {
	tree := []*Tree{}

	for i := 0; i < len(tokens); i++ {
		val, index := parse(tokens[i:])
		tree = append(tree, val)
		i += index
	}

	root := &Root{name: "Program", value: tree}

	return root
}

func parse(tokens TokenList) (tree *Tree, index int) {
	for i := 0; i < len(tokens); i++ {
		value := tokens[i]

		if value.tokenType == Variable {
			name := &Tree{value: tokens[i]}
			value := &Tree{value: tokens[i+1]}

			variable := &Tree{
				value: Token{tokenType: Empty, value: ""},
				left:  name,
				right: value}
			right, _ := parse(tokens[i+3:])

			r := &Tree{value: tokens[i+2], left: variable, right: right}

			return r, i + 3
		} else if value.tokenType == Number {
			return &Tree{value: tokens[i]}, i
		} else if value.tokenType == Identifier {
			_, posStart := tokens.Find("(")
			_, posEnd := tokens.Find(")")
			_, braceStart := tokens.Find("{")

			// Found a match
			if posStart > -1 && posEnd > -1 && braceStart > -1 {
				fncToken := Token{tokenType: Function, value: tokens[i].value}
				end := skipTo(tokens, i, "}")
				start := skipTo(tokens, i, "{") + 2

				return &Tree{value: fncToken, right: &Tree{value: tokens[start:end]}}, i
			}
		} else if value.tokenType == End {
			return &Tree{value: tokens[i]}, i
		}
	}

	return &Tree{}, len(tokens)
}
