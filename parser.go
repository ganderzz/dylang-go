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
func Parse(tokens []Token) Root {
	root := Root{name: "Program", value: parse(tokens)}

	return root
}

func parse(tokens []Token) *Tree {
	for i := 0; i < len(tokens); i++ {
		value := tokens[i]

		if value.tokenType == Variable {
			left := &Tree{
				left: &Tree{
					value: tokens[i]},
				right: &Tree{
					value: tokens[i+1]}}

			r := &Tree{value: tokens[i+2], left: left, right: parse(tokens[i+2:])}
			i = i + 2

			return r
		} else if value.tokenType == Number {
			return &Tree{value: tokens[i]}
		}
	}

	return &Tree{}
}
