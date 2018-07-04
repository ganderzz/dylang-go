package main

// Root t
type Root struct {
	name  string
	value []*Tree
}

// Tree t
type Tree struct {
	value Token
	left  *Tree
	right *Tree
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

func parse(tokens []Token) (tree *Tree, index int) {
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
		} else if value.tokenType == End {
			return &Tree{value: tokens[i]}, i
		}
	}

	return &Tree{}, len(tokens)
}
