package main

// Generate g
func Generate(root Root) string {
	code := "(function() {"
	base := root.value
	PrintTree(root)

	code = code + parseTreeToCode(base)

	return code + "})()"
}

func parseTreeToCode(tree *Tree) string {
	stringBuilder := ""

	if tree == nil {
		return stringBuilder
	}

	if tree.left != nil {
		stringBuilder += parseTreeToCode(tree.left)
	}
	if tree.right != nil {
		stringBuilder += parseTreeToCode(tree.left)
	}

	return stringBuilder + getCode(tree.value)
}

func getCode(token Token) string {
	switch token.tokenType {
	case Variable:
		return "var "
	case Identifier:
		return token.value
	case Assignment:
		return "="
	case Number:
		return token.value
	}

	return ""
}
