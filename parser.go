package main

// Root t
type Root struct {
	name  string
	value *Tree
}

// Tree t
type Tree struct {
	value Token
	left  *Tree
	right *Tree
}

// Parse parses a list of tokens into an ast
func Parse(tokens []Token) *Root {
	root := &Root{name: "Program", value: parse(tokens)}

	return root
}

func parse(tokens []Token) *Tree {
	for i := 0; i < len(tokens); i++ {
		value := tokens[i]

		if value.tokenType == Variable {
			name := &Tree{value: tokens[i]}
			value := &Tree{value: tokens[i+1]}

			variable := &Tree{
				value: Token{tokenType: Empty, value: ""},
				left:  name,
				right: value}

			r := &Tree{value: tokens[i+2], left: variable, right: parse(tokens[i+3:])}
			i = i + 2

			return r
		} else if value.tokenType == Number {
			return &Tree{value: tokens[i]}
		}
	}

	return &Tree{}
}
