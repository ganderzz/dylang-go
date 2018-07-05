package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Invalid arguments, requires path to dylang file")
		return
	}

	filePath := args[1]

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		tokens := []Token{}

		for scanner.Scan() {
			line := scanner.Text()
			split := strings.Split(line, "")

			if len(split) > 0 {
				tokens = append(tokens, Tokenize(split)...)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		ast := Parse(tokens)
		generatedCode := Generate(ast)

		fmt.Println(generatedCode)
	}
}
